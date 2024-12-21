package middleware

import (
	"github.com/gin-gonic/gin"

	"server/internal/pkg/logger"
)

func WithLogID(c *gin.Context) {
	logger.WithLogID(c)
	c.Next()
}
