package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	awsCloud "github.com/jhonquirama/my-portfolio/pkg/cloud/aws"
)

type (
	Consumer interface {
		Receive(ctx context.Context) ([]*sqs.Message, error)
		Delete(ctx context.Context, receiptHandle string) error
	}

	ConsumerConfig interface {
		SqsMaxNumberOfMessages() int64
		SqsMessageAttributeValue() string
		SqsQueueName() string
		SqsTimeout() int64
		SqsRegion() string
	}

	consumer struct {
		maxNumberOfMessages   int64
		messageAttributeValue string
		queueURL              string
		timeout               int64
		client                sqs.SQS
	}
)

func NewConsumer(
	ctx context.Context,
	config ConsumerConfig) (Consumer, error) {
	var (
		client *sqs.SQS
		output *sqs.GetQueueUrlOutput
		err    error
	)

	session := awsCloud.NewSession(ctx, config.SqsRegion())

	client = sqs.New(&session)

	if output, err = client.GetQueueUrlWithContext(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(config.SqsQueueName()),
	}); err != nil {
		return nil, err
	}

	return &consumer{
		maxNumberOfMessages:   config.SqsMaxNumberOfMessages(),
		messageAttributeValue: config.SqsMessageAttributeValue(),
		queueURL:              *output.QueueUrl,
		timeout:               config.SqsTimeout(),
		client:                *client,
	}, nil
}

func (c *consumer) Receive(ctx context.Context) ([]*sqs.Message, error) {
	var (
		receiveMessageOutput *sqs.ReceiveMessageOutput
		err                  error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(c.timeout+5))
	defer cancel()

	if receiveMessageOutput, err = c.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(c.queueURL),
		MaxNumberOfMessages:   aws.Int64(c.maxNumberOfMessages),
		WaitTimeSeconds:       aws.Int64(c.timeout),
		MessageAttributeNames: aws.StringSlice([]string{c.messageAttributeValue}),
	}); err != nil {
		return nil, err
	}

	return receiveMessageOutput.Messages, nil
}

func (c *consumer) Delete(ctx context.Context, receiptHandle string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(c.timeout))
	defer cancel()

	if _, err := c.client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}); err != nil {
		return err
	}

	return nil
}
