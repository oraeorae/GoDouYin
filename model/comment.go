package model

import (
	"time"
)

// 评论实体类
type Comment struct {
	CommentID  uint64    `json:"comment_id"`  // comment_id
	UserID     uint64    `json:"user_id"`     // user_id
	VideoID    uint64    `json:"video_id"`    // video_id
	Content    string    `json:"content"`     // content
	CreateTime time.Time `json:"create_time"` // create_time
}
