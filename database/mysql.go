package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

/**
 * @Description:全局 DB
 */
var (
	SqlDB *gorm.DB
)

/**
 * @Description: 初始化数据库
 * @return *gorm.DB
 */
func SqlClient() *gorm.DB {
	if SqlDB == nil || reflect.DeepEqual(SqlDB, gorm.DB{}) {
		// 声明连接字符串
		dsn := "root:12qwAS@(43.139.72.246:3377)/db_douyin?timeout=10s&readTimeout=30s&writeTimeout=60s"
		// 开启连接
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			//panic("failed to connect database")
		}
		SqlDB = db
		return db
	}
	return SqlDB
}
