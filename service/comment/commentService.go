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
	comment.Content = variable.Trie.Filter(comment.Content)
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
