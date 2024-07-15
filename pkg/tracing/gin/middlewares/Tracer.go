package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/tracing/tracer"
)

func OpenTelemetryTracing(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("oc", ctx)
		hcp := tracer.NewHttpContextPropagator()
		hcp.SetAction(tracer.Extract)
		hcp.SetHeader(c.Request.Header)
		hcp.SetContext(c.Request.Context())
		if err := hcp.Propagate(); err != nil {
			panic(err)
		}
		h := tracer.NewHttpTracer()
		h.SetStatus(c.Writer.Status())
		h.SetRoute(c.FullPath())
		h.SetContext(hcp.Context())
		h.Start()
		c.Request = c.Request.WithContext(h.Context())
		c.Next()
	}
}
