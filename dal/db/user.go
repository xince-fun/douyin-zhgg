package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username" gorm:"index,unique;type:varchar(32);not null"; `
	Password      string `json:"password" gorm:"type:varchar(32);not null"`
	FollowCount   int64  `gorm:"default:0"`
	FollowerCount int64  `gorm:"default:0"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// QueryUserByName query list of user info by name
func QueryUserByName(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		klog.Error("error occurred when query user by username " + err.Error())
		return nil, err
	}
	return res, nil
}

// QueryUserById query list of user info by id
func QueryUserById(ctx context.Context, userIds []int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id in ?", userIds).Find(&res).Error; err != nil {
		klog.Error("error occurred when query user by userId " + err.Error())
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, username string, password string) (int64, error) {
	user := &User{
		Username:      username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	if err := DB.WithContext(ctx).Create(&user).Error; err != nil {
		klog.Error("error occurred when create user " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}
