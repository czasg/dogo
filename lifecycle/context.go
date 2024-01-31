package lifecycle

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	RootContext, cancel = context.WithCancel(context.Background())
	RootShutdown        = []func(){cancel}
)

func AddShutdown(shutdown ...func()) {
	RootShutdown = append(RootShutdown, shutdown...)
}

func init() {
	go func() {
		defer func() {
			for _, shutdown := range RootShutdown {
				shutdown()
			}
		}()
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		if s, ok := <-ch; ok {
			logrus.Warnf("detect system signal: %s", s.String())
		}
	}()
}
