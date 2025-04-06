package logic

import (
	"context"
	"fmt"
	"time"

	healthModel "github.com/jhonquirama/my-portfolio/internal/health/business/model"
	healthPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
)

type healthService struct {
}

func NewHealthService() healthPort.HealthService {
	return &healthService{}
}

func (svc *healthService) GetHealth(_ context.Context, _ healthModel.GetHealth) (healthModel.Health, error) {

	for i := 0; i < 25; i++ {
		fmt.Println("WIP", i)
		time.Sleep(time.Second * 2)
	}
	return healthModel.Health{
		Status: "Ok",
	}, nil
}
