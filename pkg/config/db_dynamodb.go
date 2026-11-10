package config // nolint: dupl

type (
	dynamodbClient struct {
		MaxRetriesValue            int `mapstructure:"max_retries" json:"max_retries" validate:"required"`
		MaxBackoffDelaySecondValue int `mapstructure:"max_backoff_delay_second" json:"max_backoff_delay_second" validate:"required"` // nolint: lll
	}

	DynamodbClientConfig interface {
		MaxRetries() int
		MaxBackoffDelaySecond() int
	}

	// DynamodbConfig = ticketDynamodb.Config
)

func (db *dynamodbClient) MaxRetries() int {
	return db.MaxRetriesValue
}
func (db *dynamodbClient) MaxBackoffDelaySecond() int {
	return db.MaxBackoffDelaySecondValue
}

type (
	dynamodbConfig struct {
	}

	DynamodbConfig interface {
	}
)
