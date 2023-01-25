package curd

import (
	"go_douyin/dao"
	"go_douyin/global/variable"
	"go_douyin/model"
	"go_douyin/utils/redis"
	"strconv"
	"time"
)

type FollowService struct {
	followMapper *dao.FollowMapper
	redisClient  *redis.RedisClient
}

func NewFollowService() *FollowService {
	return &FollowService{
		followMapper: dao.NewFollowMapper(),
		redisClient:  redis.NewRedisClient(),
	}
}

// 关注
func (h *FollowService) FollowAction(follow model.Follow) bool {
	isFollow, _ := h.followMapper.Judge(follow)
	if isFollow > 0 {
		return false
	} else {
		follow.CreateTime = time.Now()
		row := h.followMapper.Add(follow)
		if row > 0 {
			return true
		} else {
			return false
		}
	}
}

// 取消关注
func (h *FollowService) CancalFollowAction(follow model.Follow) bool {
	isFollow, followId := h.followMapper.Judge(follow)
	if isFollow > 0 {
		follow.FollowID = followId
		row := h.followMapper.Delete(follow)
		if row > 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// 获取关注列表
func (h *FollowService) FollowList(userId uint64) []model.User {
	// 生成缓存key
	cacheKey := "follow_list_" + strconv.FormatUint(userId, 10)
	// 检查缓存键值是否存在（解决缓存击穿——可以避免大量的无效请求导致缓存服务器压力过大）
	if !variable.Filter.Test([]byte(cacheKey)) {
		// 缓存中没有数据，从数据库中获取
		followList := h.followMapper.FollowFindList(userId)
		// 将数据存储到缓存中
		h.redisClient.SetWithExpire(cacheKey, followList, 3600)
		// 添加到布隆过滤器
		variable.Filter.Add([]byte(cacheKey))
		return followList
	}
	// 从缓存中获取数据
	var followList []model.User
	err := h.redisClient.Get(cacheKey, &followList)
	if err == nil {
		// 缓存中有数据，直接返回
		return followList
	}
	// 获取分布式锁（防止缓存雪崩——在缓存中所有数据都失效时，请求量瞬间增加导致服务器压力过大，甚至瘫痪）
	lockKey := "follow_list_lock_" + strconv.FormatUint(userId, 10)
	isLock := h.redisClient.AcquireLock(lockKey, 30)
	if isLock != true {
		// 获取锁失败
		return nil
	}
	//在函数执行完成后，使用 defer 关键字来释放分布式锁
	defer h.redisClient.ReleaseLock(lockKey)

	// 缓存中没有数据，从数据库中获取
	followList = h.followMapper.FollowFindList(userId)
	// 注意这段代码是将空值缓存，防止缓存穿透
	if len(followList) == 0 {
		// 数据库中没有数据，将一个空的缓存设置一个较短的过期时间，防止缓存穿透
		h.redisClient.SetWithExpire(cacheKey, followList, 60)
		return nil
	}
	// 将数据存储到缓存中（因为关注功能数据更新频率快，所以60*15，15分钟后就过期）
	h.redisClient.SetWithExpire(cacheKey, followList, 900)
	return followList
}

// 获取粉丝列表
func (h *FollowService) FansList(userId uint64) []model.User {
	return h.followMapper.FansFindList(userId)
}

// 获取好友列表
func (h *FollowService) FriendsList(userId uint64) []model.User {
	return h.followMapper.FriendsFindList(userId)
}
