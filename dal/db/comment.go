package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId   int64  `json:"user_id" gorm:"index:idx_userid;not null"`
	VideoId  int64  `json:"video_id" gorm:"index:idx_videoid;not null"`
	Contents string `gorm:"type:varchar(255);not null"`
}

func (c *Comment) TableName() string {
	return consts.CommentTableName
}
