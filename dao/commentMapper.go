package dao

import (
	"encoding/json"
	"fmt"
	"go_douyin/global/variable"
	"go_douyin/model"
)

type CommentMapper struct{}

func NewCommentMapper() *CommentMapper {
	return &CommentMapper{}
}

func ListenComment() {
	for {
		// 获取消息
		message, err := variable.Kafka.ConsumeMessage()
		if err != nil {
			fmt.Printf("获取消息失败：%v\n", err)
			continue
		}
		// 反序列化评论数据
		var comment model.Comment
		err = json.Unmarshal([]byte(message), &comment)
		if err != nil {
			fmt.Printf("反序列化评论数据失败：%v\n", err)
			continue
		}
		// 存储评论数据
		err = SaveComment(comment)
		if err != nil {
			fmt.Printf("存储评论数据失败：%v\n", err)
			continue
		}
		fmt.Printf("成功存储评论数据：%+v\n", comment)
	}
}

func SaveComment(comment model.Comment) error {
	// 这里是将评论数据存储到数据库的代码，具体实现方式取决于你使用的数据库类型
	fmt.Println("正在存储数据库", comment)
	return nil
}
