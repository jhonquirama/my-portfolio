package logic

import (
	"context"

	healthModel "github.com/jhonquirama/hexa-scaffolding-ms/internal/health/business/model"
	healthPort "github.com/jhonquirama/hexa-scaffolding-ms/internal/health/business/port"
	apm "github.com/jhonquirama/hexa-scaffolding-ms/pkg/monitor/elastic-apm"
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

func (svc *healthService) GetHealth(ctx context.Context, filter healthModel.GetHealth) (healthModel.Health, error) {
	_ = filter
	span, _ := apm.NewSpan(ctx, apm.Service)
	defer span.End()
	return healthModel.Health{
		Status: "Ok",
	}, nil
}
