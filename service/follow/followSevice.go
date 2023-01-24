package curd

import (
	"go_douyin/dao"
	"go_douyin/model"
	"time"
)

type FollowService struct {
	followMapper *dao.FollowMapper
}

func NewFollowService() *FollowService {
	return &FollowService{
		followMapper: dao.NewFollowMapper(),
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
	return h.followMapper.FollowFindList(userId)
}

// 获取粉丝列表
func (h *FollowService) FansList(userId uint64) []model.User {
	return h.followMapper.FansFindList(userId)
}

// 获取好友列表
func (h *FollowService) FriendsList(userId uint64) []model.User {
	return h.followMapper.FriendsFindList(userId)
}
