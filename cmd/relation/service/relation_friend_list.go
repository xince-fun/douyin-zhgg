package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

type RelationFriendListService struct {
	ctx context.Context
}

// NewRelationFriendListService new RelationFriendListService
func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) RelationFriendList(req *relation.DouyinRelationFriendListRequest) ([]*user.User, error) {
	userIds := []int64{req.UserId}
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	u := users[0]

	userList, err := db.GetFriendsUsers(s.ctx, int64(u.ID))
	if err != nil {
		return nil, err
	}
	return pack.UserList(userList), nil
}
