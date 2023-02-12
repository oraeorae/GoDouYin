package main

import (
	"go_douyin/config"
	"go_douyin/dao"
	"go_douyin/database"
	"go_douyin/global/variable"
	"go_douyin/router"
)

func main() {
	// 1.初始化配置，读取配置
	config.Init()
	// 2.初始化全局变量
	variable.Init()
	// 注意初始化数据库
	database.SqlClient()

	// 5.创建协程监听消息队列
	go dao.ListenComment()

	variable.ZapLog.Info("程序正在运行")
	r := router.SetupRouter()
	r.Run(":8081")
}
