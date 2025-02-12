package kafka

import (
	"context"
	"encoding/json"

	kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	customError "github.com/jhonquirama/hexagonal-scaffolding/pkg/error"
	apm "github.com/jhonquirama/hexagonal-scaffolding/pkg/monitor/elastic-apm"
)

type (
	Config interface {
		Host() string
	}
	Producer interface {
		SendMessage(ctx context.Context, topic string, message interface{}) error
	}

	producer struct {
		producer *kafka.Producer
	}
)

func NewProducer(c Config) (Producer, error) {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":  c.Host(),
			"enable.idempotence": true,
			"acks":               "all",
			"retries":            5,
			"compression.type":   "snappy",
			"linger.ms":          20,
			"batch.size":         32768,
		})

	if err != nil {
		return nil, err
	}

	return &producer{
		producer: p,
	}, nil
}

func (p *producer) SendMessage(ctx context.Context, topic string, message interface{}) error {
	var (
		kafkaMessage *kafka.Message
		value        []byte
		err          error
	)

	span, ctx := apm.NewSpan(ctx, apm.Kafka)
	defer span.End()

	if message == nil {
		err = customError.New(ctx, customError.UnknownError)
		return err
	}

	if value, err = json.Marshal(message); err != nil {
		return customError.New(ctx, customError.UnknownError, customError.WithError(err))
	}

	kafkaMessage = &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}

	traceParentKey, traceParentValue := span.TraceHeader()
	kafkaMessage.Headers = append(kafkaMessage.Headers,
		kafka.Header{
			Key:   traceParentKey,
			Value: []byte(traceParentValue),
		})

	deliveryChan := make(chan kafka.Event)
	if err = p.producer.Produce(kafkaMessage, deliveryChan); err != nil {
		return customError.New(ctx, customError.UnknownError, customError.WithError(err))
	}

	result := <-deliveryChan
	if resultMessage := result.(*kafka.Message); resultMessage.TopicPartition.Error != nil {
		return customError.New(ctx, customError.UnknownError, customError.WithError(resultMessage.TopicPartition.Error))
	}

	return nil
}
