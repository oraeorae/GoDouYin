package router

import (
	"github.com/gin-gonic/gin"
	"go_douyin/controller"
	"go_douyin/utils/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	userController := controller.NewUserController()
	v1 := router.Group("/douyin/user")
	{
		v1.POST("register", userController.Register)
		v1.POST("login", userController.Login)
		v1.GET("/", userController.GetInfo)
	}
	//允许跨域
	router.Use(cors.Next())
	return router
}
