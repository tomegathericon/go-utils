package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/log"
	"go.uber.org/zap"
)

func Middleware(ctx context.Context) gin.HandlerFunc {
	dc := ctx
	return func(c *gin.Context) {
		if ctx == nil {
			cc, ok := c.Get("oc")
			if !ok {
				panic("no oc context")
			}
			dc = cc.(context.Context)
		}
		fields := []zap.Field{
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("clientIP", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("status", c.Writer.Status()),
			zap.String("requestURI", c.Request.RequestURI),
			zap.String("referer", c.Request.Referer()),
			zap.String("remoteAddr", c.Request.RemoteAddr),
		}
		l := log.FromContext(dc)
		l.WithOpenTelemetryTraces(c.Request.Context()).Info(c.Request.URL.Path, fields...)
		c.Next()
	}
}
