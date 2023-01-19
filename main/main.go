package main

import (
	"go_douyin/database"
	"go_douyin/router"
)

func main() {
	// 注意初始化数据库
	database.SqlClient()
	r := router.SetupRouter()
	r.Run(":8081")
}
