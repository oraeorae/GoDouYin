package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_douyin/model"
	"go_douyin/service/comment"
	"go_douyin/utils/response"
	"time"
)

type CommentController struct {
	commentService *comment.CommentService
}

func NewCommentController() *CommentController {
	return &CommentController{commentService: comment.NewCommentService()}
}

// 评论列表（此处只是个demo）
func (h *CommentController) AddComment(c *gin.Context) {
	// 获取请求参数
	var requestBody map[string]interface{}
	requestBody = make(map[string]interface{})
	// 解析请求体
	c.ShouldBindJSON(&requestBody)
	fmt.Println(requestBody)
	// 获取请求参数
	//在 HTTP POST 请求中，请求体中的数据通常是以字符串形式发送的。JSON 格式中的数字默认都是浮点型，默认都是 float64 类型
	video_id := requestBody["video_id"].(float64)
	comment_text, _ := requestBody["comment_text"].(string)
	var cc model.Comment
	cc.Content = comment_text
	// 正常是从token解析，这里是demo
	cc.UserID = 1
	cc.VideoID = uint64(video_id)
	cc.CreateTime = time.Now()
	err := h.commentService.AddComment(cc)
	if err != nil {
		response.Success(c, "评论失败", gin.H{})
		return
	}

	response.Success(c, "评论成功", gin.H{})
}

// 预加载demo，想测试结果可以放在上面
func (h *CommentController) DemoPreload(c *gin.Context) {
	h.commentService.PreloadCommentList('1')
}
