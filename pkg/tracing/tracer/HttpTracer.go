package tracer

import (
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

type HttpTracer struct {
	route  string
	status int
	*Tracer
}

func (h *HttpTracer) Status() int {
	return h.status
}

func (h *HttpTracer) SetStatus(status int) {
	h.status = status
}

func (h *HttpTracer) Route() string {
	return h.route
}

func (h *HttpTracer) SetRoute(route string) {
	h.route = route
}

func NewHttpTracer() *HttpTracer {
	return &HttpTracer{
		Tracer: NewTracer(),
	}
}

func (h *HttpTracer) Start() {
	h.SetSpanName(h.Route())
	h.StartTrace()
	h.Span().SetAttributes(semconv.HTTPStatusCode(h.Status()))
	h.Span().SetAttributes(semconv.HTTPRoute(h.Route()))
	h.EndTrace()
}
