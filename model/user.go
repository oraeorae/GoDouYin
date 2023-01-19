package model

import (
	"database/sql"
	"time"
)

// 实体类
type User struct {
	UserID        uint64        `json:"user_id"`        // user_id
	Username      string        `json:"username"`       // username
	Password      string        `json:"password"`       // password
	FollowCount   sql.NullInt64 `json:"follow_count"`   // follow_count
	FollowerCount sql.NullInt64 `json:"follower_count"` // follower_count
	CreateTime    time.Time     `json:"create_time"`    // create_time
}

// 表名
func (u *User) TableName() string {
	return "tb_users"
}
