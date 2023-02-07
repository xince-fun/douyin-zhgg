package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

type RelationFollowerListService struct {
	ctx context.Context
}

// NewRelationFollowerListService new RelationFollowerListService
func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

func (s *RelationFollowerListService) RelationFollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	userIds := []int64{req.UserId}
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	u := users[0]
	userList, err := db.GetFansUsers(s.ctx, int64(u.ID))
	if err != nil {
		return nil, err
	}
	return pack.FollowList(userList), nil
}
