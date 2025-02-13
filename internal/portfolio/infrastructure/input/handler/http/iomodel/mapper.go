package iomodel

import (
	healthModel "github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
)

func ToGetHealthModel() healthModel.GetHealth {
	return healthModel.GetHealth{}
}

func HealthModelToIO(health healthModel.Health) Health {
	return Health{
		Status: health.Status,
	}
}
