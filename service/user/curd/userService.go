package curd

import (
	"fmt"
	"go_douyin/dao"
	"go_douyin/model"
	JWT "go_douyin/service/user/token"
	"go_douyin/utils/md5_encrypt"
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
	// 预先处理密码MD5加密
	user.Password = md5_encrypt.Base64Md5(user.Password)
	row := h.userMapper.Add(user)
	if row > 0 {
		return true
	} else {
		return false
	}
}

func (h *UserService) Login(username string, password string) (bool, model.User, string) {
	var user model.User = h.userMapper.Login(username, md5_encrypt.Base64Md5(password))
	// 比较结构体是否为空
	if reflect.DeepEqual(user, model.User{}) { //判断是否为空值
		//fmt.Println("user is empty")
		return false, user, ""
	} else {
		// 生成JWT
		userTokenFactory := JWT.CreateUserFactory()
		if userToken, err := userTokenFactory.GenerateToken(user.UserID, user.Username, 28800); err == nil {
			return true, user, userToken
		} else {
			return false, user, ""
		}

	}
}

func (h *UserService) GetInfo(userid uint64, token string) model.User {
	// 解析JWT
	userTokenFactory := JWT.CreateUserFactory()
	customClaims, _ := userTokenFactory.ParseToken(token)
	fmt.Print(customClaims)
	if userid == customClaims.UserID {
		return h.userMapper.GetInfo(userid)
	} else {
		// 返回空结构体
		return model.User{}
	}

}
