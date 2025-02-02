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
}
