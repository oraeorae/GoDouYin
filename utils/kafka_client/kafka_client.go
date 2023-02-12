package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
)

// KafkaClient 封装了kafka-go库中的生产者和消费者
type KafkaClient struct {
	// producer 用于生产消息
	producer *kafka.Writer
	// consumer 用于消费消息
	consumer *kafka.Reader
}

// NewKafkaClient 创建一个Kafka工具类
func NewKafkaClient(brokers []string, topic string) *KafkaClient {
	kafkaClient := &KafkaClient{}
	// 创建生产者
	kafkaClient.producer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	// 创建消费者
	kafkaClient.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return kafkaClient
}

// ProduceMessage 生产消息
func (k *KafkaClient) ProduceMessage(message string) error {
	// 封装消息
	kafkaMessage := kafka.Message{
		Value: []byte(message),
	}
	// 发送消息
	return k.producer.WriteMessages(context.Background(), kafkaMessage)
}

// ConsumeMessage 消费消息
func (k *KafkaClient) ConsumeMessage() (string, error) {
	// 获取消息
	m, err := k.consumer.FetchMessage(context.Background())
	if err != nil {
		return "", err
	}
	return string(m.Value), nil
}

// Close 关闭Kafka工具类，释放资源
func (k *KafkaClient) Close() {
	k.producer.Close()
	k.consumer.Close()
}
