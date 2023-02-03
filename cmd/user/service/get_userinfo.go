package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/jwt"
	"context"
)

type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{
		ctx: ctx,
	}
}

// GetUserInfo get user info
func (s *GetUserInfoService) GetUserInfo(req user.DouyinUserRequest) (*user.User, error) {
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	userIds := []int64{claims.UserId}
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	u := users[0]

	// isFollow功能还没实现
	isFollow := true

	userInfo := pack.UserInfo(u, isFollow)
	return userInfo, nil
}
