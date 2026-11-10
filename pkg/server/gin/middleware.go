package gin

import (
	"github.com/gin-gonic/gin"
	ginMiddleware "github.com/jhonquirama/my-portfolio/pkg/server/gin/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *Server) middlewareConfig() {
	s.Engine.Use(otelgin.Middleware("App", otelgin.WithTracerProvider(s.Tracer)))

	s.Engine.Use(gin.Recovery())

	s.Engine.Use(func(c *gin.Context) {
		ginMiddleware.MiddlewareError(c)
	})

	s.Engine.Use(func(c *gin.Context) {
		ginMiddleware.MiddlewareTracking(c)
	})
}
