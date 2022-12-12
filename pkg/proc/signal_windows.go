//go:build windows
// +build windows

package proc

import (
	"os"
	"os/signal"
	"syscall"
)

func DealSignal(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	select {
	case sig := <-c:
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fn()
		}
	}
}
