package apm

import (
	"context"

	apm "go.elastic.co/apm/v2"
)

const (
	Handler    SpanType = "controller"
	Service    SpanType = "service"
	Repository SpanType = "repository"
	Client     SpanType = "client"
	Framework  SpanType = "framework"
	Dynamo     SpanType = "dynamodb"
	Kafka      SpanType = "kafka"
	Queries    SpanType = "queries"
)

type (
	SpanType string

	Span struct {
		*apm.Span
	}
)

func NewSpan(
	ctx context.Context,
	spanType SpanType,
) (*Span, context.Context) {
	funcName, _ := GetCallerName(1)
	span, ctx := apm.StartSpan(ctx, funcName, string(spanType))
	return &Span{span}, ctx
}

func (s *Span) End() {
	s.Span.End()
}

func (s *Span) SetLabel(key string, value interface{}) {
	labels := getLabels(key, value)
	for k, v := range labels {
		s.Span.Context.SetLabel(k, v)
	}
}

func (s *Span) TraceHeader() (string, string) {
	traceParenKey := TraceParentHeader()
	traceParentValue := FormatTraceParentHeader(s.Span.TraceContext())

	return traceParenKey, traceParentValue
}

func SpanFromContext(ctx context.Context) *Span {
	return &Span{apm.SpanFromContext(ctx)}
}

func SpanIDContext(ctx context.Context) interface{} {
	span := SpanFromContext(ctx)

	if span == nil || span.Span == nil || span.Span.SpanData == nil {
		return nil
	}

	return span.Span.TraceContext().Span
}
