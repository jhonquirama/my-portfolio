package logic

import (
	"context"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
	usersPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
)

const (
	apmService = "service"
)

type (
	Config       interface{}
	usersService struct {
		dbAuthRepository usersPort.DBAuthRepository
		cognitoAuth      usersPort.CognitoClientAuthRepository
		apm              gotel.TelemetryProvider
	}
)

func NewUsersService(dbAuthRepository usersPort.DBAuthRepository,
	cognitoAuth usersPort.CognitoClientAuthRepository, apm gotel.TelemetryProvider,
) usersPort.UsersService {
	return &usersService{
		dbAuthRepository: dbAuthRepository,
		cognitoAuth:      cognitoAuth,
		apm:              apm,
	}
}

func (svc *usersService) UsersSignUp(ctx context.Context, user usersModel.UsersSignUpInput) error {
	ctx, span := svc.apm.TraceStart(ctx, apmService)
	defer span.End()

	_, err := svc.cognitoAuth.SignUp(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (svc *usersService) UsersConfirmSignUp(ctx context.Context,
	user usersModel.UsersConfirmSignUpInputAndSignInInput) (usersModel.UsersSignInOutput, error) {
	ctx, span := svc.apm.TraceStart(ctx, apmService)
	defer span.End()

	err := svc.cognitoAuth.ConfirmSignUp(ctx, user)
	if err != nil {
		return usersModel.UsersSignInOutput{}, err
	}

	token, err := svc.UsersInitiateAuth(ctx, user)
	if err != nil {
		return usersModel.UsersSignInOutput{}, err
	}

	return token, nil
}

func (svc *usersService) UsersInitiateAuth(ctx context.Context,
	user usersModel.UsersConfirmSignUpInputAndSignInInput) (usersModel.UsersSignInOutput, error) {
	ctx, span := svc.apm.TraceStart(ctx, apmService)
	defer span.End()

	accessInfo, err := svc.cognitoAuth.UsersInitiateAuth(ctx, user)
	if err != nil {
		return usersModel.UsersSignInOutput{}, err
	}

	return accessInfo, nil
}
