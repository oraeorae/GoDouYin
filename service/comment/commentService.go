package comment

import (
	"encoding/json"
	"fmt"
	"go_douyin/dao"
	"go_douyin/global/variable"
	"go_douyin/model"
)

type CommentService struct {
	commentMapper *dao.CommentMapper
}

func NewCommentService() *CommentService {
	return &CommentService{
		commentMapper: dao.NewCommentMapper(),
	}
}

// AddComment 添加评论
func (h *CommentService) AddComment(comment model.Comment) error {
	// 进行敏感词过滤
	comment.Content = variable.Trie.Replace(comment.Content, variable.Root)
	fmt.Println(comment.Content)
	// 序列化评论数据
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	// 将评论放到消息队列
	err = variable.Kafka.ProduceMessage(string(commentJSON))
	return err
}

// 预加载评论列表
func (h *CommentService) PreloadCommentList(video_id uint64) error {
	fmt.Printf("执行预加载")
	// 将预加载的视频id放到消息队列
	err := variable.Kafka_preload.ProduceMessage(string(video_id))
	return err
}
