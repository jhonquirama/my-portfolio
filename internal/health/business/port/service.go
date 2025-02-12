package port

import (
	"context"

	healthModel "github.com/jhonquirama/hexa-scaffolding-ms/internal/health/business/model"
)

//go:generate mockery --name HealthService
type HealthService interface {
	GetHealth(ctx context.Context, filter healthModel.GetHealth) (healthModel.Health, error)
}
