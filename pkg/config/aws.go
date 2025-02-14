package config

type (
	aws struct {
		Region         string         `json:"region" validate:"required"`
		Cognito        cognito        `json:"cognito" validate:"required"`
		Dynamodb       dynamodbConfig `mapstructure:"dynamodb_config" json:"dynamodb_config" validate:"required"`
		DynamodbClient dynamodbClient `mapstructure:"dynamodb_client" json:"dynamodb_client" validate:"required"`
	}

	AwsConfig interface {
		GetRegion() string
	}
)

func (a *aws) GetRegion() string {
	return a.Region
}
