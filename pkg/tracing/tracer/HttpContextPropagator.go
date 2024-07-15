package tracer

import (
	"errors"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

type HttpContextPropagator struct {
	*contextPropagator
	header http.Header
	prop   propagation.TraceContext
}

func (p *HttpContextPropagator) Header() http.Header {
	return p.header
}

func (p *HttpContextPropagator) SetHeader(header http.Header) {
	p.header = header
}

func NewHttpContextPropagator() *HttpContextPropagator {
	return &HttpContextPropagator{
		contextPropagator: newContextPropagator(),
		header:            nil,
	}
}

func (p *HttpContextPropagator) Propagate() error {
	switch p.Action() {
	case "extract":
		p.SetContext(p.prop.Extract(p.Context(), propagation.HeaderCarrier(p.Header())))
		break
	case "inject":
		p.prop.Inject(p.Context(), propagation.HeaderCarrier(p.Header()))
		break
	default:
		return errors.New("invalid action. Chose between \"extract\" and \"inject\"")
	}
	return nil
}
