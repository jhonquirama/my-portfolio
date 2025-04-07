package gin

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jhonquirama/my-portfolio/pkg/container"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"net/http"
	"sync"
	"time"
)

type (
	Server struct {
		Engine    *gin.Engine
		Tracer    *trace.TracerProvider
		Container container.Container
		Svc       *http.Server
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

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup, quit chan struct{}, address string) {
	defer wg.Done()

	defer func(Tracer *trace.TracerProvider, ctx context.Context) {
		err := Tracer.Shutdown(ctx)
		if err != nil {
			return
		}
	}(s.Tracer, ctx)

	srv := &http.Server{
		Addr:           address,
		Handler:        s.Engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe(): %s", err)
		}
	}()

	log.Printf("HTTP server running on %s", address)

	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}

	log.Println("HTTP server exited properly")
}
