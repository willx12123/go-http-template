package srver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"server/internal/middleware"
	"server/internal/pkg/config"
	"server/internal/pkg/validator"
)

func Init() {
	isProd := config.IsProd()
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	binding.Validator = validator.NewGinValidator()

	r.Use(
		middleware.Logger,
		gin.Recovery(),
	)
	if !isProd {
		initCors(r)
	}
	routes(r)
	if isProd {
		initFrontendResource(r)
	}
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func initCors(r *gin.Engine) {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"http://localhost:5173"}
	r.Use(cors.New(conf))
}
