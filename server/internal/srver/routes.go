package srver

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"server/internal/handler"
)

func routes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(gzip.Gzip(gzip.DefaultCompression))

	{
		apiV1 := api.Group("/v1")

		apiV1.POST("/register", handler.User.Register)
		apiV1.POST("/login", handler.User.Login)
	}
}