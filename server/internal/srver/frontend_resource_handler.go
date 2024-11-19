package srver

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func initFrontendResource(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") {
			c.JSON(http.StatusNotFound, "Page not found")
			return
		}

		serveIndexHTML := func() {
			http.ServeFile(c.Writer, c.Request, "frontend/index.html")
		}

		if path == "/" {
			serveIndexHTML()
			return
		}

		filePath := "frontend" + path
		if strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".css") {
			brFilePath := filePath + ".br"
			if _, err := os.Stat(brFilePath); err == nil {
				c.Header("Content-Encoding", "br")
				c.Header("Content-Type", contentTypeByExtension(path))
				http.ServeFile(c.Writer, c.Request, brFilePath)
				return
			}
		}
		if _, err := os.Stat(filePath); err == nil {
			http.ServeFile(c.Writer, c.Request, filePath)
			return
		}
		serveIndexHTML()
	})
}

func contentTypeByExtension(path string) string {
	if strings.HasSuffix(path, ".css") {
		return "text/css; charset=utf-8"
	}
	if strings.HasSuffix(path, ".js") {
		return "application/javascript; charset=utf-8"
	}
	return "text/plain; charset=utf-8" // 默认情况，理论上不应该触发
}
