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
	// 读取配置
	config.Init()
	// 注意初始化数据库
	database.SqlClient()
	r := router.SetupRouter()
	r.Run(":8081")
}
