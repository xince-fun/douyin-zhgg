package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/kitex_gen/publish"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

func (s *PublishListService) PublishList(req *publish.DouyinPublishListRequest) ([]*feed.Video, error) {
	users, err := db.QueryUserById(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}

	videoData, err := db.QueryVideoByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*db.User, 0)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	videoList := pack.PublishList(videoData, userMap)

	return videoList, nil
}
