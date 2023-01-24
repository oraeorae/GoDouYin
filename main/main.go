package main

import (
	"go_douyin/config"
	"go_douyin/database"
	"go_douyin/router"
)

func main() {
	// 读取配置
	config.Init()
	// 注意初始化数据库
	database.SqlClient()
	r := router.SetupRouter()
	r.Run(":8081")
}
