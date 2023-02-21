package rdis

import (
	"context"
	"errors"
	"time"
)

func AddEmailValidateCode(addr string, code string) error {
	return client.Set(context.Background(), addr, code, time.Minute*5).Err()
}

var ErrVerifyFailed = errors.New("验证失败")

func VerifyEmailValidateCode(addr string, code string) error {
	cmd := client.Get(context.Background(), addr)
	if err := cmd.Err(); err != nil {
		return err
	}
	if code != cmd.Val() {
		return ErrVerifyFailed
	}

	return nil
}
