package dao

import (
	"go_douyin/database"
	"go_douyin/model"
)

// UserMapper 自定义UserMapper的类型，于user实体类对应即可
type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (UserMapper) Login(username string, password string) model.User {
	var user model.User
	//查询数据
	database.SqlDB.Where("username = ?", username).Where("password = ?", password).Find(&user)
	return user
}

func (UserMapper) GetInfo(userid uint64) model.User {
	var user model.User
	//查询数据
	database.SqlDB.Where("user_id = ?", userid).Find(&user)
	return user
}

func (UserMapper) FindAll() []model.User {
	var users []model.User
	//查询数据
	database.SqlDB.Find(&users)
	return users
}

func (UserMapper) Add(user model.User) int64 {
	//新增数据
	res := database.SqlDB.Create(&user)
	return res.RowsAffected
}

func (UserMapper) Update(user model.User) int64 {
	//更新数据
	res := database.SqlDB.Save(user)
	return res.RowsAffected
}

func (UserMapper) Delete(user model.User) int64 {
	//删除数据
	res := database.SqlDB.Delete(&model.User{}, user.UserID)
	return res.RowsAffected
}
