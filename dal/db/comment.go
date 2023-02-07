package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
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

//创建评论：创建评论记录，增加评论数
func CreateComment(ctx context.Context, comment *Comment) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("comment").Create(&comment).Error
		if err != nil {
			klog.Error("failed to create comment " + err.Error())
			return err
		}
		err = tx.Table("vedio").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			klog.Error("add comment count error " + err.Error())
			return err
		}
		err = tx.Table("comment").First(&comment).Error
		if err != nil {
			klog.Error("failed to find comment %v, %v", comment, err.Error())
			return err
		}
		return nil
	})
	return nil
}

//删除评论：删除评论记录，减少评论数，返回该评论
func DeleteComment(ctx context.Context, commentId int64) (*Comment, error) {
	var commentRaw *Comment
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("comment").Where("id = ?", commentId).First(&commentRaw).Error
		if err == gorm.ErrRecordNotFound {
			klog.Errorf("not find comment %v, %v", commentRaw, err.Error())
			return err
		}
		if err != nil {
			klog.Errorf("find comment %v fail, %v", commentRaw, err.Error())
			return err
		}
		err = tx.Table("comment").Where("id = ?", commentId).Delete(&Comment{}).Error
		if err != nil {
			klog.Error("delete comment fail " + err.Error())
			return err
		}
		err = tx.Table("video").Where("id = ?", commentRaw.VideoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			klog.Error("add comment count error " + err.Error())
			return err
		}
		return nil
	})
	return commentRaw, nil
}

//通过评论id查询一组评论信息
func QueryCommentByCommentIds(ctx context.Context, commentIds []int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Table("comment").Where("id In ?", commentIds).Find(&comments).Error
	if err != nil {
		klog.Error("failed to query comment by comment id " + err.Error())
		return nil, err
	}
	return comments, nil
}

//通过视频id号倒序返回一组评论信息
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Table("comment").Order("updated_at desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		klog.Error("failed to query comment by video id " + err.Error())
		return nil, err
	}
	return comments, nil
}
