package context

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func NewSignalContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		defer cancel()
		if s, ok := <-ch; ok {
			logrus.Warnf("detect system signal: %s", s.String())
		}
	}()
	return ctx, func() {
		cancel()
		close(ch)
	}
}
