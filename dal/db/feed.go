package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        int64     `json:"user_id" gorm:"index,unique;not null"`
	PlayUrl       string    `json:"play_url" gorm:"not null"`
	CoverUrl      string    `json:"cover_url" gorm:"not null"`
	Title         string    `json:"title" gorm:"not null"`
	FavoriteCount int64     `gorm:"default:0"`
	CommentCount  int64     `gorm:"default:0"`
	UpdatedAt     time.Time `json:"update_time" gorm:"index,sort:desc"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// QueryVideoByLatestTime query list of video info by latest create time
func QueryVideoByLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	res := make([]*Video, 0)
	t := time.UnixMilli(latestTime)
	if err := DB.WithContext(ctx).Limit(consts.LimitVideoNum).Where("update_time < ?", t).Find(&res).Error; err != nil {
		klog.Error("error occurred when query video by latest create time " + err.Error())
		return nil, err
	}
	return res, nil
}

// QueryVideoByVideoId query list of video info by video id
func QueryVideoByVideoId(ctx context.Context, videoIds []int64) ([]*Video, error) {
	var videos []*Video
	err := DB.WithContext(ctx).Where("id in (?)", videoIds).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByVideoId error " + err.Error())
		return nil, err
	}
	return videos, nil
}

// QueryVideoByUserId query list of video info by userid
func QueryVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		klog.Error("error occurred when query video by userid " + err.Error())
		return nil, err
	}
	return res, nil
}
