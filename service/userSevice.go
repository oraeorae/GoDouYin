package service

import (
	"go_douyin/dao"
	"go_douyin/model"
	"reflect"
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

func (h *UserService) Login(username string, password string) (bool, model.User) {
	var user model.User = h.userMapper.Login(username, password)
	if reflect.DeepEqual(user, model.User{}) { //判断是否为空值
		//fmt.Println("user is empty")
		return false, user
	} else {
		return true, user
	}

}

func (h *UserService) GetInfo(userid uint64) model.User {
	return h.userMapper.GetInfo(userid)
}
