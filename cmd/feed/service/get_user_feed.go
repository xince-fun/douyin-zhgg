package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
	"errors"
)

type GetUserFeedService struct {
	ctx context.Context
}

func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

func (s *GetUserFeedService) GetUserFeed(req *feed.DouyinFeedRequest) ([]*feed.Video, error) {
	resp := make([]*feed.Video, 0)

	// TODO：没有处理时间戳没有传递的情况
	videos, err := db.QueryVideoByTime(s.ctx, *req.LatestTime)
	if err != nil {
		return resp, err
	}
	if len(videos) == 0 {
		return resp, errors.New("no video")
	}

	for _, video := range videos {
		userId := video.UserId
		users, err := db.QueryUserById(s.ctx, []int64{userId})
		if err != nil {
			return resp, err
		}
		if len(users) == 0 {
			return nil, errno.UserNotExistErr
		}

		// follower 没实现
		isFollow := true
		userInfo := pack.UserInfo(users[0], isFollow)
		resp = append(resp, &feed.Video{
			Id:            int64(video.ID),
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    true, //没有实现favorite
			Title:         video.Title,
		})
	}
	return resp, nil
}
