package port

import (
	"context"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

//go:generate mockery --name PortfolioRepository
type (
	UsersRepository interface {
	}

	DBAuthRepository interface {
	}

	CognitoClientAuthRepository interface {
		SignUp(ctx context.Context, data usersModel.UsersSingUpInput) (usersModel.UsersSingUpOutput, error)
	}
)
