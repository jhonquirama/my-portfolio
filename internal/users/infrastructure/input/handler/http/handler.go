package handler

import (
	CustomCode "github.com/jhonquirama/my-portfolio/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"

	UsersPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	ioModel "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/input/handler/http/iomodel"
)

type UsersHandler struct {
	service UsersPort.UsersService
}

func NewUsersHandler(service UsersPort.UsersService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

// swagger:route GET /my-portfolio/users/sign-up register new users
//
// # Create new user
//
// Responses:
//
//	201:
//
// 400: ErrorResponse
func (h *UsersHandler) UsersSignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var userSignUpReq ioModel.UsersSignUpInput

	if err := c.BindJSON(&userSignUpReq); err != nil {
		c.Errors = append(c.Errors,
			c.Error(CustomCode.New(ctx, CustomCode.RequestBodyValidation, CustomCode.WithMessage(err.Error()))))
		return
	}

	err := h.service.UsersSignUp(ctx, ioModel.MapUsersSignUpIOModelToSignUpModel(userSignUpReq))
	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// swagger:route GET /my-portfolio/users/confirm-sign-up confirm new users
//
// # confirm new user
//
// Responses:
//
//	200:
//
// 400: ErrorResponse
func (h *UsersHandler) UsersConfirmSignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var userConfirmSignUpReq ioModel.UsersConfirmSignUpInput

	if err := c.BindJSON(&userConfirmSignUpReq); err != nil {
		c.Errors = append(c.Errors,
			c.Error(CustomCode.New(ctx, CustomCode.RequestBodyValidation, CustomCode.WithMessage(err.Error()))))
		return
	}

	response, err := h.service.UsersConfirmSignUp(ctx,
		ioModel.MapUsersConfirmSignUpIOModelToSignUpModel(userConfirmSignUpReq))
	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, ioModel.MapUsersSignInModelToSignInIOModel(response))
}
