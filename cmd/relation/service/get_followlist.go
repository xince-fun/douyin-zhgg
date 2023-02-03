package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/jwt"
	"context"
)

type GetFollowListService struct {
	ctx context.Context
}

// NewFollowListService new GetFollowListService
func NewFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(req relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	userIds := []int64{claims.UserId}
	if len(userIds) == 0 {
		return nil, errno.UserNotExistErr
	}
	id := userIds[0]
	users, err := db.GetFollowingUsers(s.ctx, id)
	if err != nil {
		return nil, err
	}
	return pack.FollowList(users), nil
}
