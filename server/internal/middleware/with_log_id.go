package middleware

import (
	"github.com/gin-gonic/gin"

	"server/internal/pkg/logger"
)

const headerFieldLogID = "X-LOG-ID"

func WithLogID(c *gin.Context) {
	logID := logger.WithLogID(c, c.GetHeader(headerFieldLogID))
	c.Next()
	c.Header(headerFieldLogID, logID)
}
