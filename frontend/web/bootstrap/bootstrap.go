package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/sso/modules/accounts/repositories"
	"github.com/rushairer/sso/modules/applications/repositories"
	"github.com/rushairer/sso/modules/auth/handlers"
	"github.com/rushairer/sso/modules/auth/services"
)

// 在这里初始化各种服务对象
func SetupServer(server *gin.Engine) {
	// 初始化认证相关的服务和处理器
	accountRepo := repositories.NewAccountRepository()
	applicationRepo := repositories.NewApplicationRepository()
	authService := services.NewAuthService(accountRepo, applicationRepo)
	authHandler := handlers.NewAuthHandler(authService, accountRepo, applicationRepo)

	// 认证相关路由
	authGroup := server.Group("/auth")
	{
		authGroup.GET("/authorize", gin.WrapF(authHandler.HandleAuthorize))
		authGroup.POST("/login", gin.WrapF(authHandler.HandleLogin))
		authGroup.POST("/token", gin.WrapF(authHandler.HandleToken))
	}

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
