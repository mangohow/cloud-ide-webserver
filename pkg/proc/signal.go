//go:build linux || darwin
// +build linux darwin

package proc

import (
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func DealSignal(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)

	setLogger(logger.Logger())
	var stopper Stopper
	for {
		select {
		case sig := <-c:
			switch sig {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				fn()
				return
			case syscall.SIGUSR1:
				if stopper == nil {
					stopper = StartProfile()
				} else {
					stopper.Stop()
					stopper = nil
				}
			case syscall.SIGUSR2:
				dumpGoroutines()
			}
		}
	}

}
