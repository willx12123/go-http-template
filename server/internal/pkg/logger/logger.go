package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/slog-zerolog"
)

var instance *slog.Logger

func Init() {
	zerologL := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	instance = slog.New(newCtxHandler(slogzerolog.Option{Logger: &zerologL}.NewZerologHandler()))
}

type LogFn func(ctx context.Context, msg string, attrs ...any)

func InfoContext(ctx context.Context, msg string, attrs ...any) {
	instance.InfoContext(ctx, msg, attrs...)
}

func ErrorContext(ctx context.Context, msg string, attrs ...any) {
	instance.ErrorContext(ctx, msg, attrs...)
}

func WarnContext(ctx context.Context, msg string, attrs ...any) {
	instance.WarnContext(ctx, msg, attrs...)
}

func DebugContext(ctx context.Context, msg string, attrs ...any) {
	instance.DebugContext(ctx, msg, attrs...)
}
