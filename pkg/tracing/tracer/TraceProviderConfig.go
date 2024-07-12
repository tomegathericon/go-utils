package tracer

import (
	"context"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type TraceProviderConfig struct {
	*otelResource
}

func NewTraceProviderConfig() *TraceProviderConfig {
	return &TraceProviderConfig{
		&otelResource{},
	}
}

func (pc *TraceProviderConfig) NewContext() context.Context {
	return context.WithValue(context.Background(), "tpc", pc)
}

type otelResource struct {
	serviceName, serviceVersion, hostName string
}

func (o *otelResource) ServiceName() string {
	return o.serviceName
}

func (o *otelResource) SetServiceName(serviceName string) {
	o.serviceName = serviceName
}

func (o *otelResource) ServiceVersion() string {
	return o.serviceVersion
}

func (o *otelResource) SetServiceVersion(serviceVersion string) {
	o.serviceVersion = serviceVersion
}

func (o *otelResource) HostName() string {
	return o.hostName
}

func (o *otelResource) SetHostName(hostName string) {
	o.hostName = hostName
}

func (t *TraceProviderConfig) NewSDKResource() (*resource.Resource, error) {
	resource, resourceErr := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceName(t.serviceName), semconv.ServiceVersion(t.serviceVersion), semconv.HostName(t.HostName())))
	if resourceErr != nil {
		return nil, resourceErr
	}
	return resource, nil
}
