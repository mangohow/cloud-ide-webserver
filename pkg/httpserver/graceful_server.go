package httpserver

import (
	"context"
	"fmt"
	"github.com/mangohow/cloud-ide-webserver/pkg/proc"
	mytime "github.com/mangohow/cloud-ide-webserver/pkg/utils"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils/waitgroup"
	"net/http"
	"time"
)

var (
	wg waitgroup.WaitGroupWapper
)

func NewServer(host string, port int, handler http.Handler) *http.Server {
	var addr string
	if host == "" || port == 0 {
		addr = ":8080"
	} else {
		addr = fmt.Sprintf("%s:%d", host, port)
	}

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func ListenAndServe(server *http.Server) {
	wg.Go(func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Println("Server shutdown at", time.Now().Format(mytime.FormatDateTime))
	})
}

type AfterCloseHandler func()

// WaitForShutdown 监听linux信号，收到信号，停止服务
func WaitForShutdown(server *http.Server, handlers ...AfterCloseHandler) {
	proc.DealSignal(func() {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
		defer cancelFunc()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("An error occurs when Server shut:%v", err)
		}

		for _, handler := range handlers {
			handler()
		}
	})

	wg.Wait()
}
