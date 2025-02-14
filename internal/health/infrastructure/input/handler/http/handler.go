package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userModel "github.com/jhonquirama/my-portfolio/internal/health/business/model"
	userPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
	ioModel "github.com/jhonquirama/my-portfolio/internal/health/infrastructure/input/handler/http/iomodel"
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

// swagger:route GET /my-portfolio/health Health get-health
//
//	Get health.
//
//	Responses:
//		200: Health
//		400: ErrorResponse
func (h *HealthHandler) GetHealth(g *gin.Context) {
	var (
		health userModel.Health
		err    error
	)

	if health, err = h.service.GetHealth(g.Request.Context(), ioModel.ToGetHealthModel()); err != nil {
		g.Errors = append(g.Errors, g.Error(err))
		return
	}

	g.JSON(http.StatusOK, ioModel.HealthModelToIO(health))
}
