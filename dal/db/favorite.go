package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
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
