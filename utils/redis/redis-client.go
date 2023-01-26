package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_douyin/global/variable"
	"time"
)

type RedisClient struct {
	Pool *redis.Pool
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		Pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", variable.Config.GetString("redis.host")+":"+variable.Config.GetString("redis.port"),
					redis.DialPassword(variable.Config.GetString("redis.auth")))
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

func (c *RedisClient) Set(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	conn := c.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, jsonData)
	return err
}

func (c *RedisClient) SetWithExpire(key string, value interface{}, expire int) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	conn := c.Pool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, jsonData, "EX", expire)
	return err
}

func (c *RedisClient) Get(key string, value interface{}) error {
	conn := c.Pool.Get()
	defer conn.Close()

	jsonData, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, value)
}

// 设置分布式锁(SET key value NX PX millseconds)
func (c *RedisClient) AcquireLock(lockKey string, timeout int) bool {
	conn := c.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", lockKey, 1, "EX", timeout, "NX")
	if err == nil {
		return true
	}
	return false
}

// 释放锁
func (c *RedisClient) ReleaseLock(lockKey string) {
	conn := c.Pool.Get()
	defer conn.Close()
	conn.Do("DEL", lockKey)
}
