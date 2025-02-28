package container

import (
	"context"
	healthLogic "github.com/jhonquirama/my-portfolio/internal/health/business/logic"
	healthPort "github.com/jhonquirama/my-portfolio/internal/health/business/port"
	portfolioLogic "github.com/jhonquirama/my-portfolio/internal/portfolio/business/logic"
	portfolioPort "github.com/jhonquirama/my-portfolio/internal/portfolio/business/port"
	"github.com/jhonquirama/my-portfolio/internal/users/business/logic"
	usersPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	cognito2 "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/output/client/cognito"
	dynamodb2 "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/output/data/dynamodb"
	"github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito"
	"github.com/jhonquirama/my-portfolio/pkg/config"
	"github.com/jhonquirama/my-portfolio/pkg/data/dynamodb"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
)

type (
	Container interface {
		HealthService() healthPort.HealthService
		PortfolioService() portfolioPort.PortfolioService
		UsersService() usersPort.UsersService
		Tracer() gotel.TelemetryProvider
	}

	tracer struct {
		tracer gotel.TelemetryProvider
	}
	health struct {
		service healthPort.HealthService
	}
	portfolio struct {
		service portfolioPort.PortfolioService
	}
	users struct {
		service usersPort.UsersService
	}
	container struct {
		health    health
		portfolio portfolio
		users     users
		tracer    tracer
	}
)

func NewContainer(ctx context.Context, cnf config.Config) (Container, error) {
	telemetry, err := gotel.NewTelemetry(ctx, cnf.ApmConfig())
	if err != nil {
		return nil, err
	}

	cognitoClient, err := cognito.NewCognito(ctx)
	if err != nil {
		return nil, err
	}

	cognitoRepository := cognito2.NewCognitoRepository(telemetry, cognitoClient, cnf.CognitoConfig())

	dynamoDB, err := dynamodb.NewDynamoDB(ctx, cnf.DynamodbClientConfig())
	if err != nil {
		return nil, err
	}

	dynamodbRepository := dynamodb2.NewAuthDynamoRepository(cnf.DynamodbConfig(), dynamoDB)

	healthService := healthLogic.NewHealthService()
	portfolioService := portfolioLogic.NewPortfolioService()
	usersService := logic.NewUsersService(dynamodbRepository, cognitoRepository, telemetry)

	return &container{
		health:    health{service: healthService},
		portfolio: portfolio{service: portfolioService},
		users:     users{service: usersService},
		tracer:    tracer{tracer: telemetry}}, nil
}

func (c container) HealthService() healthPort.HealthService {
	return c.health.service
}
func (c container) PortfolioService() portfolioPort.PortfolioService {
	return c.portfolio.service
}
func (c container) UsersService() usersPort.UsersService {
	return c.users.service
}
func (c container) Tracer() gotel.TelemetryProvider {
	return c.tracer.tracer
}
