package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/server/gin"
)

type (
	GinLambdaServer struct {
		Lambda *ginadapter.GinLambda
	}
)

func NewGinLambdaServer(ginServer *gin.Server) *GinLambdaServer {
	server := &GinLambdaServer{
		Lambda: ginadapter.New(ginServer.Engine),
	}

	return server
}

func (s *GinLambdaServer) Start() error {
	lambda.Start(s.handler)
	return nil
}

func (s *GinLambdaServer) handler(
	ctx context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return s.Lambda.ProxyWithContext(ctx, request)
}
