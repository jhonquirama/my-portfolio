package container

import (
	"context"

	healthLogic "github.com/jhonquirama/my-portfolio/internal/health/business/logic"
	healthPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
	config "github.com/jhonquirama/my-portfolio/pkg/settings"
)

type (
	Container interface {
		HealthService() healthPort.HealthService
	}
)
type health struct {
	service healthPort.HealthService
}

type container struct {
	health health
}

func NewContainer(_ context.Context, _ config.Config) (Container, error) {
	healthService := healthLogic.NewHealthService(nil)

	return &container{health: health{service: healthService}}, nil
}

func (c container) HealthService() healthPort.HealthService {
	return c.health.service
}
