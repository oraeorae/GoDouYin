package test

import (
	"go_douyin/dao"
	"go_douyin/database"
	"testing"
)

//测试统一返回接口
func TestResponse(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	userdao := dao.NewUserMapper()
	t.Fatal(userdao.FindAll())
}
