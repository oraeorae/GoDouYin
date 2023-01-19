package test

import (
	"go_douyin/dao"
	"go_douyin/database"
	"go_douyin/model"
	"testing"
	"time"
)

//单元测试
func TestPrintf(t *testing.T) {
	t.Log("TestA")
}

//测试数据库连接是否正常：新增数据
func TestDB(t *testing.T) {
	var user model.User
	user.Username = "admin"
	user.Password = "admin"
	user.CreateTime = time.Now()
	database.SqlDB = database.SqlClient()
	res := database.SqlDB.Create(&user)
	t.Fatal("在影响的行数", res.RowsAffected)
}

//测试dao层是否正常使用
func TestDao(t *testing.T) {
	//注意必须先初始化
	database.SqlClient()
	userdao := dao.NewUserMapper()
	t.Fatal(userdao.FindAll())
}
