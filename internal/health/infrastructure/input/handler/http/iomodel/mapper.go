package iomodel

import (
	healthModel "github.com/jhonquirama/hexa-scaffolding-ms/internal/health/business/model"
)

func ToGetHealthModel() healthModel.GetHealth {
	return healthModel.GetHealth{}
}

func HealthModelToIO(health healthModel.Health) Health {
	return Health{
		Status: health.Status,
	}
}
