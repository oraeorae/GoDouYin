package config

import (
	"github.com/spf13/viper"
	"go_douyin/global/variable"
	"go_douyin/utils/zap_factory"
	"log"
	"os"
)

func Init() {
	variable.Config = viper.New()
	//获取项目的执行路径
	var err error
	variable.BasePath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	variable.Config.AddConfigPath(variable.BasePath + "\\config") //设置读取的文件路径
	variable.Config.SetConfigName("config")                       //设置读取的文件名
	variable.Config.SetConfigType("yml")                          //设置文件的类型
	//尝试进行配置读取
	if err := variable.Config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	//3.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(zap_factory.ZapLogHandler)
}
