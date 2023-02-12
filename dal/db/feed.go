package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"

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
	Time          int64  `gorm:"not null"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

func QueryVideoByTime(ctx context.Context, time int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Limit(30).Where("Time <= ?", time).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
