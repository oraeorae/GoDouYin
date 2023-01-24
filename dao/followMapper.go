package dao

import (
	"go_douyin/database"
	"go_douyin/model"
)

//(以下代码可以防止SQL注入)
// FollowMapper 自定义FollowMapper的类型，于follow实体类对应即可
type FollowMapper struct{}

func NewFollowMapper() *FollowMapper {
	return &FollowMapper{}
}

// 显示关注列表
func (FollowMapper) FollowFindList(userId uint64) []model.User {
	var followedUsers []model.User
	//查询数据（联表查询）
	database.SqlDB.Table("tb_follow_list").Select("tb_users.*").Joins("JOIN tb_users ON tb_follow_list.follow_user_id = tb_users.user_id").Where("tb_follow_list.user_id = ?", userId).Scan(&followedUsers)
	return followedUsers
}

// 显示粉丝列表
func (FollowMapper) FansFindList(userId uint64) []model.User {
	var fansUsers []model.User
	//查询数据（联表查询）
	database.SqlDB.Table("tb_follow_list").Select("tb_users.*").Joins("JOIN tb_users ON tb_follow_list.user_id = tb_users.user_id").Where("tb_follow_list.follow_user_id = ?", userId).Scan(&fansUsers)
	return fansUsers
}

// 显示好友列表（互相关注才算）
// SQL语句：
//SELECT tb_users.username FROM tb_users JOIN tb_follow_list ON tb_users.user_id = tb_follow_list.follow_user_id WHERE tb_follow_list.user_id = [指定user_id]
func (FollowMapper) FriendsFindList(userId uint64) []model.User {
	var friendsUsers []model.User
	database.SqlDB.Raw("SELECT u.user_id, u.username FROM tb_users u JOIN tb_follow_list f1 ON u.user_id = f1.follow_user_id JOIN tb_follow_list f2 ON u.user_id = f2.user_id WHERE f1.user_id = ? AND f2.follow_user_id = ?", userId, userId).Scan(&friendsUsers)
	return friendsUsers
}

// 关注
func (FollowMapper) Add(follow model.Follow) int64 {
	//新增数据
	res := database.SqlDB.Create(&follow)
	return res.RowsAffected
}

// 取消关注
func (FollowMapper) Delete(follow model.Follow) int64 {
	//删除数据
	res := database.SqlDB.Delete(&model.Follow{}, follow.FollowID)
	return res.RowsAffected
}

// 判断是否已关注或者返回已关注的id
func (FollowMapper) Judge(follow model.Follow) (int64, uint64) {
	var followUser model.Follow
	res := database.SqlDB.Where("user_id = ?", follow.UserID).Where("follow_user_id = ?", follow.FollowUserID).Find(&followUser)
	return res.RowsAffected, followUser.FollowID
}

func (FollowMapper) FindAll() []model.Follow {
	var follows []model.Follow
	//查询数据
	database.SqlDB.Find(&follows)
	return follows
}

func (FollowMapper) Update(follow model.Follow) int64 {
	//更新数据
	res := database.SqlDB.Save(follow)
	return res.RowsAffected
}
