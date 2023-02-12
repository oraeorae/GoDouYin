package test

import (
	"fmt"
	"go_douyin/utils/kafka_client"
	"testing"
)

// 消息队列测试
func TestKafka(t *testing.T) {
	// 创建工具类
	kafkaClient := kafka_client.NewKafkaClient([]string{"43.139.72.246:9092"}, "test-topic")
	defer kafkaClient.Close()
	// 生产消息测试
	err := kafkaClient.ProduceMessage("Hello, Kafka!")
	if err != nil {
		t.Errorf("Failed to produce message: %v", err)
	}

	// 消费消息测试
	message, err := kafkaClient.ConsumeMessage()
	if err != nil {
		t.Errorf("Failed to consume message: %v", err)
	}
	fmt.Println("Consumed message:", message)
}
