package logic

import (
	"context"
	portfolioModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
	portfolioPort "github.com/jhonquirama/my-portfolio/internal/portfolio/business/port"
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

func (svc *portfolioService) GetHealth(_ context.Context,
	_ portfolioModel.GetHealth) (portfolioModel.Health, error) {
	return portfolioModel.Health{
		Status: "Ok",
	}, nil
}
