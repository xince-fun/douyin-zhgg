package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Favorite 可以用来找到用户喜欢的每个视频
type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"index:idx_userid;not null"`
	VideoId int64 `gorm:"index:idx_videoid;not null"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// 获取视频点赞信息  通过 用户ID 和 视频ID
func QueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*Favorite, error) {
	var favorites []*Favorite
	err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", currentId, videoIds).Find(&favorites).Error
	if err != nil {
		klog.Error("failed to query favorite record " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]*Favorite)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
}

// 增加点赞量，添加点赞记录
func CreateFavorite(ctx context.Context, favorite *Favorite, videoID int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("video").Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("add farorite count error " + err.Error())
			return err
		}
		err = tx.Table("favorite").Create(favorite).Error
		if err != nil {
			klog.Error("failed to create favorite record " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 减少点赞量，删除点赞记录
func DeleteFavorite(ctx context.Context, currentId int64, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var favorite *Favorite
		err := tx.Table("favorite").Where("user_id = ? AND video_id = ?", currentId, videoId).Delete(&favorite).Error
		if err != nil {
			klog.Error("delete favorite record fail " + err.Error())
			return err
		}

		err = tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			klog.Error("SubFavoriteCount error " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 通过用户Id 查询 其点赞视频号
func QueryFavoriteById(ctx context.Context, userId int64) ([]int64, error) {
	var favorites []*Favorite
	err := DB.WithContext(ctx).Table("favorite").Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		klog.Error("query favorite record fail " + err.Error())
		return nil, err
	}
	videoIds := make([]int64, 0)
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoId)
	}
	return videoIds, nil
}
