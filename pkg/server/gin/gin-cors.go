package gin

import (
	"github.com/gin-contrib/cors"
)

func (s *Server) corsConfig() {
	s.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Accept",
			"Origin",
			"Content-Type",
			"Authorization",
			"Content-Length",
			"Cache-Control",
			"Pragma",
			"Expires",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
