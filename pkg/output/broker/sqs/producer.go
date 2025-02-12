package aws

import (
	"context"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	awsCloud "github.com/jhonquirama/my-portfolio/pkg/cloud/aws"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
)

type ProducerConfig interface {
	ProducerQueueName() string
	ProducerTimeout() int64
	ConsumerRegion() string
}

type Producer interface {
	SendMessage(ctx context.Context, message string) error
}

type producer struct {
	queueURL string
	timeout  int64
	client   sqs.SQS
}

func NewProducer(
	ctx context.Context,
	config ProducerConfig) (Producer, error) {
	var (
		client *sqs.SQS
		output *sqs.GetQueueUrlOutput
		err    error
	)

	session := awsCloud.NewSession(ctx, config.ConsumerRegion())

	client = sqs.New(&session)

	if output, err = client.GetQueueUrlWithContext(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(config.ProducerQueueName()),
	}); err != nil {
		return nil, err
	}

	return &producer{
		queueURL: *output.QueueUrl,
		timeout:  config.ProducerTimeout(),
		client:   *client,
	}, nil
}

func (c producer) SendMessage(ctx context.Context, message string) error {
	var (
		input  *sqs.SendMessageInput
		output *sqs.SendMessageOutput
		cancel context.CancelFunc
		err    error
	)

	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(c.timeout))
	defer cancel()

	input = &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(c.queueURL),
	}

	if output, err = c.client.SendMessageWithContext(ctx, input); err != nil {
		return customError.New(ctx, customError.UnknownError, customError.WithError(err))
	}

	customLogger.Info(ctx, "message sent successfully: ", customLogger.WithObject(output))

	return nil
}
