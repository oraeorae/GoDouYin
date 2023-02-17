package dao

import (
	"encoding/json"
	"fmt"
	"go_douyin/global/variable"
	"go_douyin/model"
)

type ChatMapper struct{}

func NewChatMapper() *ChatMapper {
	return &ChatMapper{}
}

// 监听私信消息队列
func ListenChat() {
	for {
		// 获取消息
		message, err := variable.Kafka_chat.ConsumeMessage()
		if err != nil {
			fmt.Printf("获取消息失败：%v\n", err)
			continue
		}
		// 反序列化私信数据
		var chat model.Chat
		err = json.Unmarshal([]byte(message), &chat)
		if err != nil {
			fmt.Printf("反序列化私信数据失败：%v\n", err)
			continue
		}
		// 存储私信数据
		err = SaveChat(chat)
		if err != nil {
			fmt.Printf("存储私信数据失败：%v\n", err)
			continue
		}
		fmt.Printf("成功存储私信数据：%+v\n", chat)
	}
}

// 监听预加载私信消息队列
func ListenPreloadChatList() {
	for {
		fmt.Printf("正在预加载……")
		// 获取消息
		message, err := variable.Kafka_preload.ConsumeMessage()
		if err != nil {
			fmt.Printf("获取消息失败：%v\n", err)
			continue
		}
		// 此次应写存储到缓存的函数
		fmt.Printf("正在缓存" + message + "视频id的私信")
	}
}

func SaveChat(chat model.Chat) error {
	// 这里是将私信数据存储到数据库的代码，具体实现方式取决于你使用的数据库类型
	// 这里还可以补个缓存
	fmt.Println("正在存储数据库，同时更新进缓存", chat)
	return nil
}
