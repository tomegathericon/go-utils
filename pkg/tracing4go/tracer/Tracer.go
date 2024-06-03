package tracer

import (
	"context"
	"github.com/tomegathericon/go-utils/pkg/tracing4go/tracer/models"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"os"
)

type Tracer struct {
	name, version, spanName string
	context                 context.Context
	span                    oteltrace.Span
}

func (t *Tracer) Span() oteltrace.Span {
	return t.span
}

func (t *Tracer) SetSpan(span oteltrace.Span) {
	t.span = span
}

func (t *Tracer) Context() context.Context {
	return t.context
}

func (t *Tracer) SetContext(context context.Context) {
	t.context = context
}

func NewTracer() *Tracer {
	return &Tracer{
		name:    string(tracerName),
		version: string(version),
	}
}

func (t *Tracer) Name() string {
	return t.name
}

func (t *Tracer) Version() string {
	return t.version
}

func (t *Tracer) SpanName() string {
	return t.spanName
}

func (t *Tracer) SetSpanName(spanName string) {
	t.spanName = spanName
}

func NewHttpTraceProvider(cfg *models.TraceProviderConfig) *trace.TracerProvider {
	hostName, hostNameErr := os.Hostname()
	if hostNameErr != nil {
		log.Error(hostNameErr.Error())
	}
	cfg.SetHostName(hostName)
	resource, resourceErr := cfg.NewSDKResource()
	if resourceErr != nil {
		log.Error(resourceErr.Error())
	}
	httpExporter, httpExporterErr := otlptracehttp.New(context.Background())
	if httpExporterErr != nil {
		log.Error(httpExporterErr.Error())
	}
	tp := trace.NewTracerProvider(trace.WithResource(resource), trace.WithSampler(trace.AlwaysSample()), trace.WithBatcher(httpExporter))
	return tp
}

func (t *Tracer) StartTrace() {
	tp := otel.GetTracerProvider()
	tracer := tp.Tracer(t.Name(), oteltrace.WithInstrumentationVersion(t.Version()))
	ctx, span := tracer.Start(t.Context(), t.SpanName())
	t.SetContext(ctx)
	t.SetSpan(span)
}

func (t *Tracer) EndTrace() {
	t.Span().AddEvent(t.SpanName())
	defer t.Span().End()
}
