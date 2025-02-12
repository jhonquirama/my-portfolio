package apm

import (
	apmgrpc "go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptor(tracer *Tracer) grpc.UnaryServerInterceptor {
	return apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery(), apmgrpc.WithTracer(tracer.Tracer))
}

func NewStreamServerInterceptor() grpc.StreamServerInterceptor {
	return apmgrpc.NewStreamServerInterceptor()
}

func NewUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return apmgrpc.NewUnaryClientInterceptor()
}

func NewStreamClientInterceptor() grpc.StreamClientInterceptor {
	return apmgrpc.NewStreamClientInterceptor()
}
