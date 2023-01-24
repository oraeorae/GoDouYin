package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_douyin/model"
	"go_douyin/service/follow"
	JWT "go_douyin/service/user/token"
	"go_douyin/utils/response"
)

type FollowController struct {
	followService *curd.FollowService
}

func NewFollowController() *FollowController {
	return &FollowController{followService: curd.NewFollowService()}
}

// 关系操作
func (h *FollowController) FollowAction(c *gin.Context) {
	var requestBody map[string]interface{}
	requestBody = make(map[string]interface{})
	// 解析请求体
	c.ShouldBindJSON(&requestBody)
	// 获取请求参数
	//在 HTTP POST 请求中，请求体中的数据通常是以字符串形式发送的。JSON 格式中的数字默认都是浮点型，默认都是 float64 类型
	token := requestBody["token"].(string)
	toUserId, _ := requestBody["to_user_id"].(float64)
	// 1-关注，2-取消关注
	actionType, _ := requestBody["action_type"].(float64)

	// 解析JWT
	userTokenFactory := JWT.CreateUserFactory()
	customClaims, _ := userTokenFactory.ParseToken(token)

	if actionType == 1 {
		// 关注
		var follow model.Follow
		follow.UserID = customClaims.UserID
		follow.FollowUserID = uint64(toUserId)
		if h.followService.FollowAction(follow) {
			response.Success(c, "关注成功", gin.H{})
		} else {
			response.Fail(c, -1, "关注失败", gin.H{})
		}

	} else if actionType == 2 {
		// 取消关注
		var follow model.Follow
		follow.UserID = customClaims.UserID
		follow.FollowUserID = uint64(toUserId)
		if h.followService.CancalFollowAction(follow) {
			response.Success(c, "取消关注成功", gin.H{})
		} else {
			response.Fail(c, -1, "取消关注失败", gin.H{})
		}
	} else {
		fmt.Println(actionType)
		response.Fail(c, -1, "操作失败", gin.H{})
	}
}

// 用户关注列表
func (h *FollowController) FollowList(c *gin.Context) {
	// 获取请求参数
	token := c.Query("token")
	userId := c.Query("user_id")
	// 解析JWT
	userTokenFactory := JWT.CreateUserFactory()
	customClaims, _ := userTokenFactory.ParseToken(token)
	fmt.Println(userId)
	followList := h.followService.FollowList(customClaims.UserID)
	response.Success(c, "获取成功", gin.H{
		"user_list": followList,
	})
}

// 用户粉丝列表
func (h *FollowController) FansList(c *gin.Context) {
	// 获取请求参数
	token := c.Query("token")
	userId := c.Query("user_id")
	// 解析JWT
	userTokenFactory := JWT.CreateUserFactory()
	customClaims, _ := userTokenFactory.ParseToken(token)
	fmt.Println(userId)
	followList := h.followService.FansList(customClaims.UserID)
	response.Success(c, "获取成功", gin.H{
		"user_list": followList,
	})
}

// 用户好友列表
func (h *FollowController) FriendsList(c *gin.Context) {
	// 获取请求参数
	token := c.Query("token")
	userId := c.Query("user_id")
	userTokenFactory := JWT.CreateUserFactory()
	customClaims, _ := userTokenFactory.ParseToken(token)
	fmt.Println(userId)
	followList := h.followService.FriendsList(customClaims.UserID)
	response.Success(c, "获取成功", gin.H{
		"user_list": followList,
	})
}
