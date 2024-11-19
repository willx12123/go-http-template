package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/pkg/logger"
)

func RenderOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RenderForbidden(ctx *gin.Context) {
	logger.Default.Error("RenderForbidden")
	ctx.JSON(http.StatusForbidden, nil)
}

func RenderBadRequest(ctx *gin.Context, err error) {
	logger.Default.Error("RenderBadRequest", zap.Error(err))
	ctx.JSON(http.StatusBadRequest, nil)
}

func RenderInternalServerError(ctx *gin.Context, err error, message string) {
	logger.Default.Error("RenderInternalServerError", zap.Error(err), zap.String("message", message))
	ctx.JSON(http.StatusInternalServerError, message)
}
