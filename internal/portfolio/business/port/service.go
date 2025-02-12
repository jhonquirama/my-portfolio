package port

import (
	"context"

	healthModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
)

//go:generate mockery --name PortfolioService
type PortfolioService interface {
	GetHealth(ctx context.Context, filter healthModel.GetHealth) (healthModel.Health, error)
}
