package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jhonquirama/my-portfolio/pkg/container"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability"
)

type (
	Server struct {
		Engine    *gin.Engine
		Apm       *observability.Provider
		Container container.Container
	}
)

func NewGinServer(tracer *observability.Provider, container container.Container) *Server {
	server := &Server{
		Engine:    gin.Default(),
		Apm:       tracer,
		Container: container,
	}

	server.corsConfig()
	server.middlewareConfig()
	server.routerConfig()

	return server
}

func (s *Server) Run(address string) error {
	// register tracing provider as a global provider
	stopTracingProvider, err := s.Apm.RegisterAsGlobal()
	if err != nil {
		return err
	}

	defer func() {
		if err = stopTracingProvider(context.Background()); err != nil {
			return
		}
	}()

	return s.Engine.Run(address)
}
