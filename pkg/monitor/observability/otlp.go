package observability

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tc "go.opentelemetry.io/otel/trace"
	"runtime"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	Service SpanType = "service"
	Client  SpanType = "client"
)

type (
	SpanType string

	Config interface {
		ApmServiceName() string
		ApmServerHost() string
		ApmSecretToken() string
		ApmLogFile() string
		ApmLogLevel() string
		ApmEnvironment() string
	}
)

type Provider struct {
	serviceName string
	exporterURL string
	Provider    *trace.TracerProvider
}

var tp tc.Tracer // nolint: gochecknoglobals

func NewProvider(ctx context.Context, cnf Config) (*Provider, error) {
	e, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(cnf.ApmServerHost()),
		otlptracehttp.WithHeaders(map[string]string{"content-type": "application/json"}),
		otlptracehttp.WithInsecure(), // as connection is not need to be secured in this project
	))
	if err != nil {
		return nil, err
	}

	// setup resource
	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cnf.ApmServiceName()),
		semconv.ServiceVersionKey.String("0.0.1"),
	)

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(e),
		trace.WithResource(r),
	)

	tp = tracerProvider.Tracer(cnf.ApmServiceName())

	return &Provider{
		serviceName: cnf.ApmServiceName(),
		exporterURL: cnf.ApmServerHost(),
		Provider:    tracerProvider,
	}, nil
}

func (p *Provider) RegisterAsGlobal() (func(ctx context.Context) error, error) {
	// set global provider
	otel.SetTracerProvider(p.Provider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return p.Provider.Shutdown, nil
}

func NewSpan(ctx context.Context, spanType SpanType) (context.Context, tc.Span) {
	ctx, span := tp.Start(ctx, string(spanType))
	return ctx, span
}

func GetCallerName(skip int) (string, int) {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", 0
	}

	f := runtime.FuncForPC(pc)
	if f == nil {
		return "", 0
	}

	_, line := f.FileLine(pc)

	return f.Name(), line
}
