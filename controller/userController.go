package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_douyin/middleware/validator"
	"go_douyin/model"
	"go_douyin/service/user/curd"
	"go_douyin/utils/response"
	"reflect"
	"strconv"
)

type UserController struct {
	UserService *curd.UserService
}

func NewUserController() *UserController {
	return &UserController{UserService: curd.NewUserService()}
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
	var login validator.Login
	login = c.MustGet("login").(validator.Login)
	isLogin, userDB, token := h.UserService.Login(login.Username, login.Password)
	if isLogin {
		response.Success(c, "登录成功", gin.H{
			"user_id": userDB.UserID,
			"token":   token,
		})
	} else {
		response.Fail(c, -1, "登录失败", gin.H{})
	}
}

func (h *UserController) GetInfo(c *gin.Context) {
	// 获取参数
	userid, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	fmt.Println(c.Query("user_id"))
	fmt.Println(userid)
	// 获取参数
	token := c.Query("token")
	user := h.UserService.GetInfo(userid, token)
	if !reflect.DeepEqual(user, model.User{}) {
		response.Success(c, "获取成功", gin.H{
			"id":             user.UserID,
			"name":           user.Username,
			"follow_count":   user.FollowCount,
			"follower_count": user.FollowerCount,
			"is_follow":      false,
		},
		)
	} else {
		response.Fail(c, -1, "获取失败", gin.H{})
	}
}
