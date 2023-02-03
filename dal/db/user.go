package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
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
