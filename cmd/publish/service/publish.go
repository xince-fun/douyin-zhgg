package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/publish"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/ffmpeg"
	"ByteTech-7815/douyin-zhgg/pkg/oss"
	"bytes"
	"context"
	"github.com/godruoyi/go-snowflake"
	"strconv"
	"strings"
	"time"
)

type PublishService struct {
	ctx context.Context
}

// NewPublishService new public service
func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

// Publish upload video
func (s *PublishService) Publish(req *publish.DouyinPublishActionRequest) error {
	videoData := req.Data
	title := req.Title

	fileId := snowflake.ID()

	fileName := strings.Join([]string{strconv.Itoa(int(fileId)), ".mp4"}, "")

	// 上传视频到本地
	filePath := strings.Join([]string{oss.Path, "/public/", fileName}, "")
	err := oss.PublishVideoToPublic(videoData, filePath)
	if err != nil {
		return errno.PublishVideoToPublicErr
	}

	// 上传视频到oss
	objectKey := strings.Join([]string{"video/", fileName}, "")
	err = oss.PublishVideoToOss(objectKey, filePath)
	if err != nil {
		return errno.PublishVideoToOssErr
	}

	// 获取视频播放地址
	signedURL, err := oss.QueryOssVideoURL(objectKey)
	if err != nil {
		return errno.GetOssVideoUrlErr
	}
	videoUrl := strings.Split(signedURL, "?")[0]

	// 获取视频封面
	coverName := strings.Join([]string{strconv.Itoa(int(fileId)), ".jpg"}, "")
	cover, err := ffmpeg.GetVideoCover(filePath, consts.FrameNum)
	if err != nil {
		return errno.GetVideoCoverErr
	}

	// 上传封面到oss
	objectKey = strings.Join([]string{"cover/", coverName}, "")
	coverReader := bytes.NewReader(cover)
	err = oss.PublishCoverToOss(objectKey, coverReader)
	if err != nil {
		return errno.PublishCoverToOssErr
	}

	// 获取视频封面地址
	signedURL, err = oss.QueryOssCoverURL(objectKey)
	if err != nil {
		return errno.GetOssCoverUrlErr
	}
	coverUrl := strings.Split(signedURL, "?")[0]

	err = db.PublishVideo(s.ctx, &db.Video{
		UserId:        req.UserId,
		Title:         title,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		UpdatedAt:     time.Now()})
	if err != nil {
		return err
	}
	return nil
}
