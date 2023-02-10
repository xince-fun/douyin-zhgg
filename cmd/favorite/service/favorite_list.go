package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/favorite"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"context"
	"errors"
	"sync"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList get video information that users like
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*feed.Video, error) {
	//检查用户是否存在
	user, err := db.QueryUserById(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("user not exist")
	}

	//获取目标用户的点赞视频id号
	videoIds, err := db.QueryFavoriteById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	//获取点赞视频的信息
	videoData, err := db.QueryVideoByVideoId(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}

	//获取点赞视频的用户id号
	userIds := make([]int64, 0)
	for _, video := range videoData {
		userIds = append(userIds, video.UserId)
	}

	//获取点赞视频的用户信息
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var favoriteMap map[int64]*db.Favorite
	var relationMap map[int64]*db.Follow

	//if user not logged in
	if req.UserId == -1 {
		favoriteMap = nil
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error
		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteMap, err = db.QueryFavoriteByIds(s.ctx, req.UserId, videoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()
		//获取关注信息
		go func() {
			defer wg.Done()
			relationMap, err = db.QueryRelationByIds(s.ctx, req.UserId, userIds)
			if err != nil {
				relationErr = err
				return
			}
		}()

		wg.Wait()
		if favoriteErr != nil {
			return nil, favoriteErr
		}
		if relationErr != nil {
			return nil, relationErr
		}
	}
	videoList := pack.VedioList(req.UserId, videoData, userMap, favoriteMap, relationMap)
	return videoList, nil
}
