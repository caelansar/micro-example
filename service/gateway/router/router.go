package router

import (
	"github.com/gin-gonic/gin"
	"micro-example/service/gateway/handler"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/signup", handler.SignUpHandler)
		api.POST("/login", handler.LoginHandler)
	}
	return router
}