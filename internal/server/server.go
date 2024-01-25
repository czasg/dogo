package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"proj/public/config"
	"time"
)

func Run(ctx context.Context, handler http.Handler, cfg config.HttpConfig) error {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
	go func() {
		<-ctx.Done()
		// grace time before close
		newCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GraceTimeout)*time.Second)
		defer cancel()
		if err := s.Shutdown(newCtx); err != nil {
			logrus.Errorf("server shutdown error: %v", err)
		}
	}()
	return s.ListenAndServe()
}
