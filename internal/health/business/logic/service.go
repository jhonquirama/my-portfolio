package logic

import (
	"context"

	healthModel "github.com/jhonquirama/my-portfolio/internal/health/business/model"
	healthPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
)

type healthService struct {
	healthRepository healthPort.HealthRepository
}

func NewHealthService(
	healthRepository healthPort.HealthRepository,
) healthPort.HealthService {
	return &healthService{
		healthRepository: healthRepository,
	}
}

func (svc *healthService) GetHealth(_ context.Context, _ healthModel.GetHealth) (healthModel.Health, error) {
	return healthModel.Health{
		Status: "Ok",
	}, nil
}
