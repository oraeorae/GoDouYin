package router

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"go_douyin/controller"
	"go_douyin/utils/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	var xssMdlwr xss.XssMw
	router.Use(xssMdlwr.RemoveXss())
	userController := controller.NewUserController()
	followController := controller.NewFollowController()
	// 用户组：登录注册，获取个人信息
	v1 := router.Group("/douyin/user")
	{
		v1.POST("register", userController.Register)
		v1.POST("login", userController.Login)
		v1.GET("/", userController.GetInfo)
	}
	// 社交组：关注，粉丝相关信息
	v2 := router.Group("/douyin/relation")
	{
		v2.POST("action", followController.FollowAction)
		v2.GET("follow/list", followController.FollowList)
		v2.GET("follower/list", followController.FansList)
		v2.GET("friend/list", followController.FriendsList)
	}
	//允许跨域
	router.Use(cors.Next())
	return router
}
