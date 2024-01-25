package shutdown

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	ch  = make(chan os.Signal, 1)
	arr = []func(){}
)

func Bind(c ...func()) {
	arr = append(arr, c...)
}

func Shutdown() {
	for _, c := range arr {
		c()
	}
	arr = []func(){}
}

func init() {
	go func() {
		defer Shutdown()
		signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		if s, ok := <-ch; ok {
			logrus.Warnf("detect system signal: %s", s.String())
		}
	}()
}
