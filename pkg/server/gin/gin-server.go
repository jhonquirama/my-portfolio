package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/container"
	apm "github.com/jhonquirama/hexa-scaffolding-ms/pkg/monitor/elastic-apm"
)

type (
	Server struct {
		Engine    *gin.Engine
		Apm       *apm.Tracer
		Container container.Container
	}
)

func NewGinServer(tracer *apm.Tracer, container container.Container) *Server {
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
	defer s.Apm.Close()
	return s.Engine.Run(address)
}
