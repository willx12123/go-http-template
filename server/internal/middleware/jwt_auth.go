package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/pkg/logger"
	"server/internal/service/authenticate"
	userSvc "server/internal/service/user"
)

func JwtAuth(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		logger.Default.Error("abort request because there is no Authorization Token")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claim, err := authenticate.DecodeToken(tokenStr)
	if err != nil {
		logger.Default.Error("abort request because token is invalid", zap.Error(err))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := userSvc.FindUser(ctx, claim.UserID)
	if err != nil {
		logger.Default.Error("token is valid but find user fail in db", zap.Uint("user_id", claim.UserID), zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	authenticate.SetCurrentUser(ctx, user)
	ctx.Next()
}
