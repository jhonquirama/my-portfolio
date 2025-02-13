package logic

import (
	"context"
	portfolioModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
	portfolioPort "github.com/jhonquirama/my-portfolio/internal/portfolio/business/port"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

type (
	Config           interface{}
	portfolioService struct {
		portfolioRepository portfolioPort.PortfolioRepository
	}
)

func NewPortfolioService(
	portfolioRepository portfolioPort.PortfolioRepository,
) portfolioPort.PortfolioService {
	return &portfolioService{
		portfolioRepository: portfolioRepository,
	}
}

func (svc *portfolioService) GetHealth(ctx context.Context,
	filter portfolioModel.GetHealth) (portfolioModel.Health, error) {
	_ = filter
	span, _ := apm.NewSpan(ctx, apm.Service)
	defer span.End()
	return portfolioModel.Health{
		Status: "Ok",
	}, nil
}
