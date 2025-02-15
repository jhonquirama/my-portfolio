package logic

import (
	"context"

	healthModel "github.com/jhonquirama/my-portfolio/internal/health/business/model"
	healthPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
)

type healthService struct {
}

func NewHealthService() healthPort.HealthService {
	return &healthService{}
}

func (svc *healthService) GetHealth(_ context.Context, _ healthModel.GetHealth) (healthModel.Health, error) {
	return healthModel.Health{
		Status: "Ok",
	}, nil
}
