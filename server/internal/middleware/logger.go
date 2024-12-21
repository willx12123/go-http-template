package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"server/internal/pkg/logger"
	"server/internal/pkg/logger/attr"
)

func Logger(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	latencyTime := time.Since(startTime).Milliseconds()
	statusCode := c.Writer.Status()

	var logFn logger.LogFn
	if statusCode >= 200 && statusCode < 300 {
		logFn = logger.InfoContext
	} else if statusCode >= 500 {
		logFn = logger.ErrorContext
	} else {
		logFn = logger.InfoContext
	}
	logFn(
		c,
		"[middleware.Logger] api request",
		attr.Int("status", statusCode),
		slog.String("method", c.Request.Method),
		slog.String("uri", c.Request.RequestURI),
		slog.Int64("cost", latencyTime),
		slog.String("ip", c.ClientIP()),
	)
}
