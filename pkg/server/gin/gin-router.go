package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	usersHandler "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/input/handler/http"

	healthHandler "github.com/jhonquirama/my-portfolio/internal/health/infrastructure/input/handler/http"
)

const (
	v1Group string = "/v1/my-portfolio"
	//HS
	getHealthPath string = "/health"

	// USERS
	usersPathSingUp string = "/users/sing-up"
)

func (s *Server) routerConfig() {
	api := s.Engine.Group("/api")

	s.healthRoutes(api)
	s.usersRoutes(api)
}

func (s *Server) healthRoutes(api *gin.RouterGroup) {
	routes := api.Group(v1Group)
	{
		newHealthHandler := healthHandler.NewHealthHandler(s.Container.HealthService())
		routes.GET(getHealthPath, newHealthHandler.GetHealth)
	}
}

func (s *Server) portfolioRoutes(api *gin.RouterGroup) {
}

func (s *Server) usersRoutes(api *gin.RouterGroup) {
	routes := api.Group(v1Group)
	fmt.Println(routes)
	{
		newUsersHandler := usersHandler.NewUsersHandler(s.Container.UsersService())
		routes.POST(usersPathSingUp, newUsersHandler.UsersSingUp)
	}
}
