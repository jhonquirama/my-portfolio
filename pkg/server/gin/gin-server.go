package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jhonquirama/my-portfolio/pkg/container"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
	"go.opentelemetry.io/otel/sdk/trace"
)

type (
	Server struct {
		Engine    *gin.Engine
		Tracer    *trace.TracerProvider
		Apm       gotel.TelemetryProvider
		Container container.Container
	}
)

func NewGinServer(tracer gotel.TelemetryProvider, container container.Container) *Server {
	server := &Server{
		Engine:    gin.Default(),
		Tracer:    tracer.GetProvider(),
		Container: container,
	}

	server.corsConfig()
	server.middlewareConfig()
	server.routerConfig()

	return server
}

func (s *Server) Run(address string) error {
	defer func(Tracer *trace.TracerProvider, ctx context.Context) {
		err := Tracer.Shutdown(ctx)
		if err != nil {
			return
		}
	}(s.Tracer, context.Background())
	return s.Engine.Run(address)
}
