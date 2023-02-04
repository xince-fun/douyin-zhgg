package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
	"crypto/md5"
	"fmt"
	"io"
)

type RegisterUserService struct {
	ctx context.Context
}

// NewRegisterUserService new RegisterUserService
func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{ctx: ctx}
}

// RegisterUser register user info
func (s *RegisterUserService) RegisterUser(req *user.DouyinUserRegisterRequest) (int64, error) {
	users, err := db.QueryUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	userId, err := db.CreateUser(s.ctx, req.Username, password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
