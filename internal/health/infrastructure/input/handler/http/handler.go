package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userModel "github.com/jhonquirama/hexagonal-scaffolding/internal/health/business/model"
	userPort "github.com/jhonquirama/hexagonal-scaffolding/internal/health/business/port"
	ioModel "github.com/jhonquirama/hexagonal-scaffolding/internal/health/infrastructure/input/handler/http/iomodel"
	apm "github.com/jhonquirama/hexagonal-scaffolding/pkg/monitor/elastic-apm"
)

type HealthHandler struct {
	service userPort.HealthService
}

func NewHealthHandler(
	service userPort.HealthService,
) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

// swagger:route GET /hexagonal-scaffolding/health Health get-health
//
//	Get health.
//
//	Responses:
//		200: Health
//		400: ErrorResponse
func (h *HealthHandler) GetHealth(g *gin.Context) {
	var (
		ctx    = apm.RequestContext(g)
		health userModel.Health
		err    error
	)

	if health, err = h.service.GetHealth(ctx, ioModel.ToGetHealthModel()); err != nil {
		g.Errors = append(g.Errors, g.Error(err))
		return
	}

	g.JSON(http.StatusOK, ioModel.HealthModelToIO(health))
}
