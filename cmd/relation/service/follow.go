package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"context"
)

type FollowService struct {
	ctx context.Context
}

// NewFollowService new FollowService
func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}

func (s *FollowService) Follow(req *relation.DouyinRelationActionRequest) error {
	// 检查是否已经关注
	if db.IsFollowing(s.ctx, req.UserId, req.ToUserId) == false {
		// 未关注， 则添加关注
		return db.CreateFollow(s.ctx, req.UserId, req.ToUserId)
	}
	return nil
}
