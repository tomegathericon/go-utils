package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/tracing/tracer"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
)

func OpenTelemetryTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := tracer.NewTracer()
		t.SetSpanName(c.FullPath())
		t.SetContext(c.Request.Context())
		t.StartTrace()
		t.Span().SetAttributes(semconv.HTTPStatusCode(c.Writer.Status()))
		t.Span().SetAttributes(semconv.HTTPRoute(c.FullPath()))
		t.EndTrace()
		c.Request = c.Request.WithContext(t.Context())
		c.Next()
	}
}
