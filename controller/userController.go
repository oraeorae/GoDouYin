package controller

import (
	"github.com/gin-gonic/gin"
	"go_douyin/model"
	"go_douyin/service"
	"go_douyin/utils/response"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{UserService: service.NewUserService()}
}

func (h *UserController) Register(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	if h.UserService.Register(user) {
		response.Success(c, "注册成功", gin.H{})
		//c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": "注册成功"})
	} else {
		response.Fail(c, -1, "注册失败", gin.H{})
	}
}

func (h *UserController) Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	isLogin, userDB := h.UserService.Login(user.Username, user.Password)
	if isLogin {
		response.Success(c, "登录成功", gin.H{
			"user_id": userDB.UserID,
		})
	} else {
		response.Fail(c, -1, "登录失败", gin.H{})
	}
}
