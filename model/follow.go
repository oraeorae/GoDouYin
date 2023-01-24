package model

import "time"

// 关注粉丝实体类
type Follow struct {
	FollowID     uint64    `json:"follow_id"`      // follow_id
	UserID       uint64    `json:"user_id"`        // user_id
	FollowUserID uint64    `json:"follow_user_id"` // follow_user_id
	CreateTime   time.Time `json:"create_time"`    // create_time
}

// 表名
func (u *Follow) TableName() string {
	return "tb_follow_list"
}
