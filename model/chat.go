package model

import "time"

// 这里后面的记得改和以前的一样
type Chat struct {
	ChatID      uint64    `gorm:"column:chat_id;primaryKey;autoIncrement"`
	SenderID    uint64    `gorm:"column:sender_id;not null"`
	RecipientID uint64    `gorm:"column:recipient_id;not null"`
	Message     string    `gorm:"column:message;not null"`
	SendTime    time.Time `gorm:"column:send_time;not null;default:CURRENT_TIMESTAMP"`
}

func (Chat) TableName() string {
	return "tb_chat"
}
