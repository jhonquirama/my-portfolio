package gin

import (
	"github.com/gin-gonic/gin"
	usersHandler "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/input/handler/http"

	healthHandler "github.com/jhonquirama/my-portfolio/internal/health/infrastructure/input/handler/http"
)

const (
	v1Group string = "/v1/my-portfolio"
	// HS
	getHealthPath string = "/health"

	// USERS
	usersPathSignUp        string = "/users/sign-up"
	usersPathConfirmSignUp string = "/users/confirm-sign-up"
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

func (s *Server) usersRoutes(api *gin.RouterGroup) {
	routes := api.Group(v1Group)
	{
		newUsersHandler := usersHandler.NewUsersHandler(s.Container.UsersService())
		routes.POST(usersPathSignUp, newUsersHandler.UsersSignUp)
		routes.POST(usersPathConfirmSignUp, newUsersHandler.UsersConfirmSignUp)
	}
}
