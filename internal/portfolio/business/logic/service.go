package logic

import (
	"context"
	portfolioModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
	portfolioPort "github.com/jhonquirama/my-portfolio/internal/portfolio/business/port"
)

type (
	Config           interface{}
	portfolioService struct {
	}
)

func NewPortfolioService() portfolioPort.PortfolioService {
	return &portfolioService{}
}

func (svc *portfolioService) GetHealth(_ context.Context,
	_ portfolioModel.GetHealth) (portfolioModel.Health, error) {
	return portfolioModel.Health{
		Status: "Ok",
	}, nil
}
