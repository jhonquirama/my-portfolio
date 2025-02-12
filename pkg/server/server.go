package server

import (
	"context"

	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/container"
	customLogger "github.com/jhonquirama/hexa-scaffolding-ms/pkg/log"
	apm "github.com/jhonquirama/hexa-scaffolding-ms/pkg/monitor/elastic-apm"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/output/notify/slack"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/server/gin"
	lambda "github.com/jhonquirama/hexa-scaffolding-ms/pkg/server/gin-lambda"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/settings"
)

type Server struct {
	GinServer *gin.Server
	GinLambda *lambda.GinLambdaServer
	Config    settings.Config
}

func NewServer(ctx context.Context) (*Server, error) {
	var (
		cnf    settings.Config
		cnt    container.Container
		tracer *apm.Tracer
		err    error
	)

	if cnf, err = settings.NewConfigFromFile(); err != nil {
		if cnf, err = settings.NewConfigFromSystemManager(ctx); err != nil {
			return nil, err
		}
	}

	if cnt, err = container.NewContainer(ctx, cnf); err != nil {
		return nil, err
	}

	if tracer, err = apm.NewAPM(cnf.ApmConfig()); err != nil {
		return nil, err
	}

	slackConfig := slack.NewSlackClient(cnf.Slack())

	tracer.Flush(nil)

	customLogger.SetEnvironment(cnf.ApmConfig().ApmEnvironment(), slackConfig)

	ginServer := gin.NewGinServer(tracer, cnt)

	return &Server{
		GinServer: ginServer,
		GinLambda: lambda.NewGinLambdaServer(ginServer),
		Config:    cnf,
	}, nil
}
