package port

import (
	"context"

	healthModel "github.com/jhonquirama/my-portfolio/internal/health/business/model"
)

//go:generate mockery --name HealthService --filename hs_service_mock.go --output ../../../../pkg/mocks/services
type HealthService interface {
	GetHealth(ctx context.Context, filter healthModel.GetHealth) (healthModel.Health, error)
}
