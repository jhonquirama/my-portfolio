package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	customLogger "github.com/jhonquirama/hexagonal-scaffolding/pkg/log"
)

func NewSession(ctx context.Context, awsRegion string) session.Session {
	var (
		awsSession *session.Session
		err        error
	)

	if awsSession, err = session.NewSession(&aws.Config{
		Region: &awsRegion},
	); err != nil {
		customLogger.Fatal(ctx, err)
	}

	return *awsSession
}
