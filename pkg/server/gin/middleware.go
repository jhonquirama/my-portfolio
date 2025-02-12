package gin

import (
	"github.com/gin-gonic/gin"

	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
	ginMiddleware "github.com/jhonquirama/my-portfolio/pkg/server/gin/middleware"
)

func (s *Server) middlewareConfig() {
	s.Engine.Use(apm.GinMiddleware(s.Engine, s.Apm.Tracer), gin.Recovery())

	s.Engine.Use(gin.Recovery())

	s.Engine.Use(func(c *gin.Context) {
		apm.SetAuthenticatedUserToTransactionCtx(c)
	})

	s.Engine.Use(func(c *gin.Context) {
		ginMiddleware.MiddlewareError(c)
	})

	s.Engine.Use(func(c *gin.Context) {
		ginMiddleware.MiddlewareTracking(c)
	})
}
