package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/internal/pkg/logger"
	"server/internal/pkg/logger/attr"
	"server/internal/service/authenticate"
	userSvc "server/internal/service/user"
)

func JwtAuth(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		logger.ErrorContext(ctx, "abort request because there is no Authorization Token")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claim, err := authenticate.DecodeToken(tokenStr)
	if err != nil {
		logger.ErrorContext(ctx, "abort request because token is invalid", attr.Err(err))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := userSvc.FindUser(ctx, claim.UserID)
	if err != nil {
		logger.ErrorContext(ctx, "token is valid but find user fail in db", attr.Uint("user_id", claim.UserID), attr.Err(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	authenticate.SetCurrentUser(ctx, user)
	ctx.Next()
}
