package trace

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"proj/public/utils"
	"sync/atomic"
	"time"
)

func NewTraceHandler() gin.HandlerFunc {
	var n int64
	s := utils.Snow{}
	return func(c *gin.Context) {
		key := c.GetHeader(string(TraceKey))
		if key == "" {
			key = fmt.Sprintf(
				traceFormat,
				s.Next(),
				time.Now().UnixNano(),
				atomic.AddInt64(&n, 1),
			)
		}
		c.Set(string(TraceKey), key)
		c.Header(string(TraceKey), key)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), TraceKey, key))
		c.Next()
	}
}

type traceKey string

const (
	TraceKey    traceKey = "X-Trace-Key"
	traceFormat string   = "trace-%d-%d-%d"
)

func RequestWithContext(ctx context.Context, req *http.Request) *http.Request {
	key, ok := ctx.Value(TraceKey).(string)
	if ok {
		req.Header.Set(string(TraceKey), key)
	}
	return req
}

func LogWithContext(ctx context.Context, entry *logrus.Entry) *logrus.Entry {
	key, ok := ctx.Value(TraceKey).(string)
	if ok {
		return entry.WithField(string(TraceKey), key)
	}
	return entry
}
