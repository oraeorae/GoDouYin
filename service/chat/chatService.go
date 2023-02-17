package chat

import (
	"encoding/json"
	"fmt"
	"go_douyin/global/variable"
	"go_douyin/model"
	"go_douyin/utils/kafka_client"
)

type ChatService struct {
	//chatMapper *dao.ChatMapper
	kafkaClients *kafka_client.KafkaClient
}

func NewChatService() *ChatService {
	return &ChatService{
		kafkaClients: &kafka_client.KafkaClient{},
	}
}

// AddChat 发送消息
func (h *ChatService) AddChat(chat model.Chat) error {
	// 序列化评论数据
	chatJSON, err := json.Marshal(chat)
	if err != nil {
		return err
	}
	// 将私信放到消息队列
	err = variable.Kafka_chat.ProduceMessage(string(chatJSON))
	return err
}

// GetMessages 获取聊天记录
func (h *ChatService) GetMessages(senderID string, recipientID string) ([]string, error) {
	fmt.Println("此处直接获取数据库或者缓存的数据即可")
	return nil, nil
}
