package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/godruoyi/go-snowflake"
)

func PublishVideo(ctx context.Context, video *Video) error {
	video.ID = uint(snowflake.ID())
	if err := DB.WithContext(ctx).Create(&video).Error; err != nil {
		klog.Error("error occurred when public video " + err.Error())
		return err
	}
	return nil
}
