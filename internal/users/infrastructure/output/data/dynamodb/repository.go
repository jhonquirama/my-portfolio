package dynamodb

import (
	authPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	noSql "github.com/jhonquirama/my-portfolio/pkg/data/dynamodb"
)

type (
	Config interface {
	}
	dbAuthRepository struct {
		dbClient noSql.Dynamodb
	}
)

func NewAuthDynamoRepository(c Config, dbClient noSql.Dynamodb) authPort.DBAuthRepository {
	return &dbAuthRepository{
		dbClient: dbClient,
	}
}
