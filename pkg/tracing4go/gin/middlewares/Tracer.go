package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/tracing4go/tracer"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func OpenTelemetryTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := tracer.NewTracer()
		t.SetSpanName(c.FullPath())
		t.SetContext(context.Background())
		t.StartTrace()
		t.Span().SetAttributes(semconv.HTTPStatusCode(c.Writer.Status()))
		t.Span().SetAttributes(semconv.HTTPRoute(c.FullPath()))
		t.EndTrace()
		c.Request = c.Request.WithContext(t.Context())
		c.Next()
	}
}
