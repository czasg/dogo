package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"proj/lifecycle"
	"time"
)

func Run(ctx context.Context, handler http.Handler, cfg lifecycle.HttpConfig) error {
	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
	lifecycle.AddShutdown(func() {
		if err := ctx.Err(); err != nil {
			logrus.Errorf("server err: %v", err)
		}
		newCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GraceTimeout)*time.Second)
		defer cancel()
		if err := s.Shutdown(newCtx); err != nil {
			logrus.Errorf("server shutdown error: %v", err)
		}
	})
	logrus.WithFields(logrus.Fields{
		"Addr": s.Addr,
	}).Info("server start")
	return s.ListenAndServe()
}
