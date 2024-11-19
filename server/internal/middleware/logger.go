package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/pkg/logger"
)

func Logger(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	latencyTime := time.Since(startTime).Milliseconds()
	statusCode := c.Writer.Status()

	var logFn func(msg string, fields ...zap.Field)
	if statusCode >= 200 && statusCode < 300 {
		logFn = logger.Default.Info
	} else if statusCode >= 500 {
		logFn = logger.Default.Error
	} else {
		logFn = logger.Default.Warn
	}
	logFn(
		"[middleware.Logger] api request",
		zap.Int("status", statusCode),
		zap.String("method", c.Request.Method),
		zap.String("uri", c.Request.RequestURI),
		zap.Int64("cost", latencyTime),
		zap.String("ip", c.ClientIP()),
	)
}
