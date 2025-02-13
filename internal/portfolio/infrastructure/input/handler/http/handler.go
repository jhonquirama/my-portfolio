package handler

import (
	userModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
	"net/http"

	"github.com/gin-gonic/gin"

	portfolioPort "github.com/jhonquirama/my-portfolio/internal/portfolio/business/port"
	ioModel "github.com/jhonquirama/my-portfolio/internal/portfolio/infrastructure/input/handler/http/iomodel"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

type PortfolioHandler struct {
	service portfolioPort.PortfolioService
}

func NewPortfolioHandler(service portfolioPort.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
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
func (h *PortfolioHandler) GetHealth(g *gin.Context) {
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
