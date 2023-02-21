package service

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/mangohow/cloud-ide-webserver/conf"
	"github.com/mangohow/cloud-ide-webserver/internal/dao/rdis"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/smtp"
	"time"
)

type EmailConfig struct {
	host   string
	port   uint32
	sender string
	auth   string
}

type EmailService struct {
	logger *logrus.Logger
	ch     chan *email.Email
	pool   *email.Pool
	config *EmailConfig
}

func NewEmailService() *EmailService {
	return &EmailService{
		logger: logger.Logger(),
		ch:     make(chan *email.Email, 1024),
		config: &EmailConfig{
			host:   conf.EmailConfig.Host,
			port:   conf.EmailConfig.Port,
			sender: conf.EmailConfig.SenderEmail,
			auth:   conf.EmailConfig.AuthCode,
		},
	}
}

var numbers = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func (e *EmailService) Send(addr string) error {
	// 生成6位数验证码
	validateCode := make([]byte, 6)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		validateCode[i] = numbers[rand.Intn(10)]
	}

	e.logger.Debugf("adrr:%s, code:%s", addr, string(validateCode))
	m := &email.Email{
		From:    e.config.sender,
		To:      []string{addr},
		Subject: "Cloud Code验证码",
		Text:    append([]byte("您好,您的验证码为:"), validateCode...),
		Sender:  "Cloud Code",
	}

	// 存入redis
	err := rdis.AddEmailValidateCode(addr, string(validateCode))
	if err != nil {
		return err
	}

	// 发送邮件
	e.ch <- m

	return nil
}

func (e *EmailService) Start() error {
	pool, err := email.NewPool(fmt.Sprintf("%s:%d", e.config.host, e.config.port),
		4, smtp.PlainAuth("", e.config.sender, e.config.auth, e.config.host))
	if err != nil {
		e.logger.Errorf("connect to mail failed, err:%v", err)
		return err
	}
	e.pool = pool

	for i := 0; i < 4; i++ {
		go func() {
			for m := range e.ch {
				err := pool.Send(m, 10*time.Second)
				if err != nil {
					e.logger.Errorf("send email error:%v", err)
				}
			}

		}()
	}

	return nil
}
