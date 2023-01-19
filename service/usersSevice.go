package service

import (
	"go_douyin/dao"
	"go_douyin/model"
	"time"
)

type UserService struct {
	userMapper *dao.UserMapper
}

func NewUserService() *UserService {
	return &UserService{
		userMapper: dao.NewUserMapper(),
	}
}

func (h *UserService) Register(user model.User) bool {
	user.CreateTime = time.Now()
	row := h.userMapper.Add(user)
	if row > 0 {
		return true
	} else {
		return false
	}
}

func (h *UserService) Login(username string, password string) bool {
	row := h.userMapper.Login(username, password)
	if row > 0 {
		return true
	} else {
		return false
	}
}
