package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_douyin/service/video"
	"net/http"
)

type VideoController struct {
}

func NewVideoController() *VideoController {
	return &VideoController{}
}
func (h *VideoController) UploadFile(c *gin.Context) {
	file, header, _ := c.Request.FormFile("data")
	video.SaveFile(file, header)
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println(token, title)
	//service.UploadService(param1, param2, filename)
	c.String(http.StatusOK, "File uploaded successfully")
}
