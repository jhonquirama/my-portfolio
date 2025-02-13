package handler

import (
	CustomCode "github.com/jhonquirama/my-portfolio/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"

	UsersPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	ioModel "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/input/handler/http/iomodel"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

type UsersHandler struct {
	service UsersPort.UsersService
}

func NewUsersHandler(service UsersPort.UsersService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

// swagger:route GET /my-portfolio/users/sing-up register new users
//
// # Create new user
//
// Responses:
//
//	201:
//
// 400: ErrorResponse
func (h *UsersHandler) UsersSingUp(c *gin.Context) {
	var (
		ctx           = apm.RequestContext(c)
		userSingUpReq ioModel.UsersSingUpInput
	)

	if err := c.BindJSON(&userSingUpReq); err != nil {
		c.Errors = append(c.Errors,
			c.Error(CustomCode.New(ctx, CustomCode.RequestBodyValidation, CustomCode.WithMessage(err.Error()))))
		return
	}

	err := h.service.UsersSingUp(ctx, ioModel.MapUsersSingUpIOModelToSingUpModel(userSingUpReq))
	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(http.StatusCreated, nil)
}
