package main

import (
	"go_douyin/config"
	"go_douyin/database"
	"go_douyin/global/variable"
	"go_douyin/router"
)

func main() {
	// 初始化全局变量
	variable.Init()
	// 注意初始化数据库
	database.SqlClient()
	// 4.初始化配置，读取配置
	config.Init()

	variable.ZapLog.Info("程序正在运行")
	r := router.SetupRouter()
	r.Run(":8081")
}
