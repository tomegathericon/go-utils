package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/tracing/models"
)

func OpenTelemetryTracing(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("oc", ctx)
		h := models.NewHttpTracer()
		h.SetStatus(c.Writer.Status())
		h.SetRoute(c.FullPath())
		h.SetContext(c.Request.Context())
		h.Start()
		c.Request = c.Request.WithContext(h.Context())
		c.Next()
	}
}
