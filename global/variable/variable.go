package variable

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/syyongx/go-wordsfilter"
	"github.com/willf/bloom"
	"go.uber.org/zap"
	"go_douyin/global/my_errors"
	"go_douyin/utils/kafka_client"
	_ "gorm.io/gorm"
	"log"
	"os"
)

// 全局变量（注意首字母大写）
var (
	BasePath string // 定义项目的根目录
	//EventDestroyPrefix = "Destroy_"            //  程序退出时需要销毁的事件前缀
	//ConfigKeyPrefix    = "Config_"             //  配置文件键值缓存时，键的前缀
	//DateFormat         = "2006-01-02 15:04:05" //  设置全局日期时间格式

	// 全局日志指针
	ZapLog *zap.Logger

	// 全局配置文件
	Config *viper.Viper

	// 创建布隆过滤器
	Filter *bloom.BloomFilter

	// 全局消息队列
	// 评论的队列
	Kafka *kafka_client.KafkaClient
	// 预加载的队列
	Kafka_preload *kafka_client.KafkaClient
	// 私信聊天的队列
	Kafka_chat *kafka_client.KafkaClient

	//全局敏感词过滤
	Trie *wordsfilter.WordsFilter
	Root map[string]*wordsfilter.Node

	//ConfigYml       ymlconfig_interf.YmlConfigInterf // 全局配置文件指针
	//ConfigGormv2Yml ymlconfig_interf.YmlConfigInterf // 全局配置文件指针

	//gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	//GormDbMysql      *gorm.DB // 全局gorm的客户端连接
	//GormDbSqlserver  *gorm.DB // 全局gorm的客户端连接
	//GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

	//雪花算法全局变量
	//SnowFlake snowflake_interf.InterfaceSnowFlake

	//websocket
	//WebsocketHub              interface{}
	//WebsocketHandshakeSuccess = `{"code":200,"msg":"ws连接成功","data":""}`
	//WebsocketServerPingMsg    = "Server->Ping->Client"

	//casbin 全局操作指针
	//Enforcer *casbin.SyncedEnforcer
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	//2.检查storage/logs 目录是否存在
	if _, err := os.Stat(BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}

}

func Init() {
	//1.检查配置文件以及日志目录等非编译性的必要条件
	//checkRequiredFolders()
	//2.初始化布隆过滤器
	Filter = bloom.New(1000000, 5)
	// 3.创建监听评论的消息队列（后面改到配置那里）
	Kafka = kafka_client.NewKafkaClient([]string{"43.139.72.246:9092"}, "comment-topic")
	Kafka_preload = kafka_client.NewKafkaClient([]string{"43.139.72.246:9092"}, "comment-preload-topic")
	Kafka_chat = kafka_client.NewKafkaClient([]string{"43.139.72.246:9092"}, "chat-topic")

	// 4.创建敏感词过滤树
	fmt.Println("创建敏感词前缀树")
	Trie = wordsfilter.New()
	Root, _ = Trie.GenerateWithFile(BasePath + "/config/sensitive_words.txt")
}
