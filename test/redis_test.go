package test

import (
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

type User struct {
	Username string
	Age      int
}

// 测试出来不行，但是放到主函数就可以，哪里出错呢
func TestRedisClient(t *testing.T) {
	_ = redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "43.139.72.246"+":"+"6381",
				redis.DialPassword("000415"))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}
