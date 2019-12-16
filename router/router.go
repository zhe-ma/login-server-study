package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/handler/user"
	"github.com/zhe-ma/login-server-study/router/middleware"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middleware.NoCache)
	engine.Use(middleware.Options)
	engine.Use(middleware.Secure)
	engine.Use(middlewares...)

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	engine.GET("/version", handler.GetVersion)
	engine.GET("/computer_info", handler.GetComputerInfo)

	engine.POST("v1/login", user.Login)

	userGroup := engine.Group("v1/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("", user.Create)
		userGroup.GET("/:id", user.Get)
		userGroup.DELETE("/:id", user.Delete)
		userGroup.PUT("/:id", user.Update)
		userGroup.GET("", user.List)
	}

	return engine
}
