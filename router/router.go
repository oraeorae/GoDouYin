package router

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"go_douyin/controller"
	"go_douyin/middleware/ratelimit"
	"go_douyin/middleware/validator"
	"go_douyin/utils/cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	var xssMdlwr xss.XssMw
	router.Use(xssMdlwr.RemoveXss())
	userController := controller.NewUserController()
	followController := controller.NewFollowController()
	videoController := controller.NewVideoController()
	commentController := controller.NewCommentController()

	// 用户组：登录注册，获取个人信息
	v1 := router.Group("/douyin/user")
	{
		v1.POST("register", userController.Register)
		v1.POST("login", validator.LoginValidationMiddleware(), userController.Login)
		v1.GET("/", userController.GetInfo)
	}
	// 社交组：关注，粉丝相关信息
	v2 := router.Group("/douyin/relation")
	{
		v2.Use(ratelimit.RateLimiter())
		v2.POST("action", followController.FollowAction)
		v2.GET("follow/list", followController.FollowList)
		v2.GET("follower/list", followController.FansList)
		v2.GET("friend/list", followController.FriendsList)
	}
	// 用户组：登录注册，获取个人信息
	v3 := router.Group("/douyin/publish")
	{
		v3.POST("action", videoController.UploadFile)
	}
	// 社交组：评论功能
	v4 := router.Group("/douyin/comment")
	{
		v4.POST("action", commentController.AddComment)
	}
	//允许跨域
	router.Use(cors.Next())
	return router
}
