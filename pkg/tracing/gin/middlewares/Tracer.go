package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/tracing/models"
	"go.opentelemetry.io/otel/propagation"
)

func OpenTelemetryTracing(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("oc", ctx)
		prop := propagation.TraceContext{}
		ctx := prop.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		h := models.NewHttpTracer()
		h.SetStatus(c.Writer.Status())
		h.SetRoute(c.FullPath())
		h.SetContext(ctx)
		h.Start()
		c.Request = c.Request.WithContext(h.Context())
		c.Next()
	}
}
