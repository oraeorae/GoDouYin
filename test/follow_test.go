package test

import (
	"go_douyin/dao"
	"go_douyin/database"
	"go_douyin/model"
	"testing"
	"time"
)

//关注
func TestFollowAction(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	follow := model.Follow{}
	// 因为有外键的限制，所以必须都是存在的
	follow.UserID = 51
	follow.FollowUserID = 1
	follow.CreateTime = time.Now()
	t.Log(followdao.Add(follow))
}

//取消关注
func TestCancalFollowAction(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	follow := model.Follow{}
	// 因为有外键的限制，所以必须都是存在的
	follow.FollowID = 6
	follow.UserID = 5
	follow.FollowUserID = 51
	follow.CreateTime = time.Now()
	t.Log(followdao.Delete(follow))
}

//取消关注
func TestJudgeFollowAction(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	follow := model.Follow{}
	// 因为有外键的限制，所以必须都是存在的
	follow.UserID = 5
	follow.FollowUserID = 51
	t.Log(followdao.Judge(follow))
}

//测试关注的联表语句是否正常
func TestFollow(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	t.Log(followdao.FollowFindList(1))
}

//测试粉丝的联表语句是否正常
func TestFans(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	t.Log(followdao.FansFindList(1))
}

//测试好友的联表语句是否正常
func TestFriend(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	followdao := dao.NewFollowMapper()
	t.Log(followdao.FriendsFindList(1))
}
