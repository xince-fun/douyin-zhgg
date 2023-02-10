package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/favorite"
	"context"
	"errors"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction implement the like and unlike operations
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	videos, err := db.QueryVideoByVideoId(s.ctx, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}
	//若ActionType（操作类型）等于1，则向favorite表创建一条记录，同时向video表的目标video增加点赞数
	//若ActionType等于2，则向favorite表删除一条记录，同时向video表的目标video减少点赞数
	//若ActionType不等于1和2，则返回错误
	if req.ActionType == 1 {
		favorite := &db.Favorite{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		}

		err := db.CreateFavorite(s.ctx, favorite, req.VideoId)
		if err != nil {
			return err
		}
	}
	if req.ActionType == 2 {
		err := db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}

	}
	if req.ActionType != 1 && req.ActionType != 2 {
		return errors.New("action type no equal 1 and 2")
	}
	return nil
}
