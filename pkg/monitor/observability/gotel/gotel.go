package gotel

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

//go:generate mockery --name TelemetryProvider

// TelemetryProvider is an interface for the Telemetry provider.
type TelemetryProvider interface {
	GetServiceName() string
	TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span)
	Shutdown(ctx context.Context)
	GetProvider() *trace.TracerProvider
}

type Config interface {
	ApmServiceName() string
	ApmServerHost() string
	ApmSecretToken() string
	ApmLogFile() string
	ApmLogLevel() string
	ApmEnvironment() string
}

// Telemetry is a wrapper around the OpenTelemetry logger, meter, and tracer.
type Telemetry struct {
	tp     *trace.TracerProvider
	tracer oteltrace.Tracer
	cfg    Config
}

// NewTelemetry creates a new Telemetry instance.
func NewTelemetry(ctx context.Context, cfg Config) (TelemetryProvider, error) {
	rp := newResource(cfg.ApmServiceName(), "0.0.1")

	ntp, err := newTracerProvider(ctx, rp, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer: %w", err)
	}
	tracer := ntp.Tracer(cfg.ApmServiceName())

	return &Telemetry{
		tp:     ntp,
		tracer: tracer,
		cfg:    cfg,
	}, nil
}

// newTracerProvider creates a new tracer provider with the OTLP gRPC exporter.
func newTracerProvider(ctx context.Context, res *resource.Resource, cnf Config) (*trace.TracerProvider, error) {
	exporter, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(cnf.ApmServerHost()),
		otlptracehttp.WithHeaders(map[string]string{"content-type": "application/json"}),
		otlptracehttp.WithInsecure(), // as connection is not need to be secured in this project
	))
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	// Create Resource
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	// set global provider
	otel.SetTracerProvider(tp)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}

// newResource creates a new OTEL resource with the service name and version.
func newResource(serviceName string, serviceVersion string) *resource.Resource {
	hostName, _ := os.Hostname()

	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(serviceVersion),
		semconv.HostName(hostName),
	)
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

func (t *Telemetry) GetProvider() *trace.TracerProvider {
	return t.tp
}

// GetServiceName returns the name of the service.
func (t *Telemetry) GetServiceName() string {
	return t.cfg.ApmServiceName()
}

// TraceStart starts a new span with the given name. The span must be ended by calling End.
func (t *Telemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span) { //nolint:ireturn
	//nolint: spancheck
	return t.tracer.Start(ctx, name)
}

// Shutdown shuts down the logger, meter, and tracer.
func (t *Telemetry) Shutdown(ctx context.Context) {
	err := t.tp.Shutdown(ctx)
	if err != nil {
		return
	}
}
