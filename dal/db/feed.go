package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        int64  `json:"user_id" gorm:"index,unique;not null"`
	PlayUrl       string `json:"play_url" gorm:"not null"`
	CoverUrl      string `json:"cover_url" gorm:"not null"`
	Title         string `json:"title" gorm:"not null"`
	FavoriteCount int64  `gorm:"default:0"`
	CommentCount  int64  `gorm:"default:0"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}
