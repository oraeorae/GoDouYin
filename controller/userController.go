package controller

import (
	"github.com/gin-gonic/gin"
	"go_douyin/model"
	"go_douyin/service"
	"go_douyin/utils/response"
	"net/http"
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
		response.Success(c, "注册成功", gin.H{
			"id": '1',
		}, "token")
		//c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": "注册成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": "OK", "data": "注册失败"})
	}
}

func (h *UserController) Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	if h.UserService.Login(user.Username, user.Password) {
		response.Success(c, "登录成功", gin.H{
			"id": '1',
		}, "token")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": "OK", "data": "登录失败"})
	}
}
