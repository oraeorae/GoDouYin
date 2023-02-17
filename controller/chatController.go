package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_douyin/model"
	"go_douyin/service/chat"
	"go_douyin/utils/response"
	"time"
)

type ChatController struct {
	chatService *chat.ChatService
}

func NewChatController() *ChatController {
	return &ChatController{
		chatService: chat.NewChatService(),
	}
}

// 发送信息
func (h *ChatController) AddChat(c *gin.Context) {
	// 获取请求参数
	var requestBody map[string]interface{}
	requestBody = make(map[string]interface{})
	// 解析请求体
	c.ShouldBindJSON(&requestBody)
	// 获取请求参数
	//在 HTTP POST 请求中，请求体中的数据通常是以字符串形式发送的。JSON 格式中的数字默认都是浮点型，默认都是 float64 类型
	user_id := requestBody["user_id"].(float64)
	to_user_id := requestBody["to_user_id"].(float64)
	content, _ := requestBody["content"].(string)
	var cc model.Chat
	cc.RecipientID = uint64(to_user_id)
	cc.Message = content
	cc.SenderID = uint64(user_id)
	cc.SendTime = time.Now()
	err := h.chatService.AddChat(cc)
	if err != nil {
		response.Success(c, "私信失败", gin.H{})
		fmt.Println(err)
		return
	}
	response.Success(c, "私信成功", gin.H{})
}

// 聊天记录
func (h *ChatController) ListChat(c *gin.Context) {
	// 获取请求参数
	ToUserId := c.Query("to_user_id")
	userId := c.Query("user_id")
	data, err := h.chatService.GetMessages(userId, ToUserId)
	if err != nil {
		response.Success(c, "获取失败", gin.H{})
		return
	}
	response.Success(c, "获取成功", gin.H{
		"message_list": data,
	})
}
