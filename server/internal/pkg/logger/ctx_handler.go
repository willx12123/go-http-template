package logger

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const fieldLogID = "log_id"

func WithLogID(ctx *gin.Context, logID string) string {
	if logID == "" {
		logID = uuid.NewString()
	}
	ctx.Set(fieldLogID, slog.String(fieldLogID, logID))
	return logID
}

type ctxHandler struct {
	slog.Handler
}

func newCtxHandler(h slog.Handler) slog.Handler {
	return ctxHandler{h}
}

func (h ctxHandler) Handle(ctx context.Context, r slog.Record) error {
	if attr, ok := ctx.Value(fieldLogID).(slog.Attr); ok {
		r.AddAttrs(attr)
	}
	return h.Handle(ctx, r)
}
