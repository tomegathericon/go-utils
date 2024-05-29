package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func Log4Gin(log *zap.Logger, timeFormat string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
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
			zap.String("traceID", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()),
			zap.String("spanID", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()),
		}
		log.Info(c.Request.URL.Path, fields...)
	}
}
