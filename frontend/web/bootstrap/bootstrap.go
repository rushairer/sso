package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 在这里初始化各种服务对象
func SetupServer(server *gin.Engine) {

	// routers
	testGroup := server.Group("/test")
	{
		testGroup.GET(
			"/alive",
			func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		)
	}

	// UI
	{
		server.Static("/_next", "./frontend/web/public/_next")

		server.StaticFile("/", "./frontend/web/public/index.html")

		server.StaticFile("/favicon.ico", "./frontend/web/public/favicon.ico")
		server.StaticFile("/404.html", "./frontend/web/public/404.html")
		server.StaticFile("/placeholder.svg", "./frontend/web/public/placeholder.svg")
	}
}
