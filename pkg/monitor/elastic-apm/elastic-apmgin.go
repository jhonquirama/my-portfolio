package apm

import (
	"context"

	gin "github.com/gin-gonic/gin"
	apmgin "go.elastic.co/apm/module/apmgin/v2"
	apm "go.elastic.co/apm/v2"
)

const (
	userIDHeader       = "user_id"
	superadminIDHeader = "superadmin_id"
	superadminHeader   = "superadmin"
	emailHeader        = "email"
	usernameHeader     = "full_name"
)

func GinMiddleware(router *gin.Engine, tracer *apm.Tracer) gin.HandlerFunc {
	return apmgin.Middleware(router, apmgin.WithTracer(tracer), apmgin.WithPanicPropagation())
}

func SetAuthenticatedUserToTransactionCtx(c *gin.Context) {
	var (
		userIDAuthenticated string
		tx                  = TransactionFromContext(c.Request.Context())
	)

	tx.SetLabel(superadminHeader, c.GetHeader(superadminHeader))
	tx.SetLabel(userIDHeader, c.GetHeader(userIDHeader))
	tx.SetLabel(superadminIDHeader, c.GetHeader(superadminIDHeader))

	if superadminID := c.GetHeader(superadminIDHeader); superadminID != "" {
		userIDAuthenticated = superadminID
	} else if userID := c.GetHeader(userIDHeader); userID != "" {
		userIDAuthenticated = userID
	}

	if userIDAuthenticated != "" {
		tx.Transaction.Context.SetUserID(userIDAuthenticated)
		tx.Transaction.Context.SetUserEmail(c.Request.Header.Get(emailHeader))
		tx.Transaction.Context.SetUsername(c.Request.Header.Get(usernameHeader))
	}
}

func RequestContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()

	ctx = context.WithValue(ctx, userIDHeader, c.GetHeader(userIDHeader))             // nolint: staticcheck, revive
	ctx = context.WithValue(ctx, superadminIDHeader, c.GetHeader(superadminIDHeader)) // nolint: staticcheck, revive
	ctx = context.WithValue(ctx, superadminHeader, c.GetHeader(superadminHeader))     // nolint: staticcheck, revive

	return ctx
}
