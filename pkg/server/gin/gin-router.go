package gin

import (
	"github.com/gin-gonic/gin"

	healthHandler "github.com/jhonquirama/hexa-scaffolding-ms/internal/health/infrastructure/input/handler/http"
)

const (
	v1Group       string = "/v1/hexa-scaffolding-ms"
	getHealthPath string = "/health"
)

func (s *Server) routerConfig() {
	api := s.Engine.Group("/api")

	s.healthRoutes(api)
}

func (s *Server) healthRoutes(api *gin.RouterGroup) {
	routes := api.Group(v1Group)
	{
		newHealthHandler := healthHandler.NewHealthHandler(s.Container.HealthService())
		routes.GET(getHealthPath, newHealthHandler.GetHealth)
	}
}
