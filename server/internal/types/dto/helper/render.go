package helper

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"server/internal/pkg/logger"

	"server/internal/pkg/logger/attr"
)

func RenderOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RenderForbidden(ctx *gin.Context) {
	logger.ErrorContext(ctx, "RenderForbidden")
	ctx.JSON(http.StatusForbidden, nil)
}

func RenderBadRequest(ctx *gin.Context, err error) {
	logger.ErrorContext(ctx, "RenderBadRequest", attr.Err(err))
	ctx.JSON(http.StatusBadRequest, nil)
}

func RenderInternalServerError(ctx *gin.Context, err error, message string) {
	logger.ErrorContext(ctx, "RenderInternalServerError", attr.Err(err), slog.String("message", message))
	ctx.JSON(http.StatusInternalServerError, message)
}
