package main

import (
	"ByteTech-7815/douyin-zhgg/cmd/favorite/pack"
	"ByteTech-7815/douyin-zhgg/cmd/favorite/service"
	favorite "ByteTech-7815/douyin-zhgg/kitex_gen/favorite"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp = new(favorite.DouyinFavoriteActionResponse)

	if req.UserId == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp.BaseResp = pack.BuilBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)

	if err != nil {
		resp.BaseResp = pack.BuilBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuilBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.DouyinFavoriteListResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuilBaseResp(errno.ParamErr)
		return resp, nil
	}

	vedioList, err := service.NewFavoriteListService(ctx).FavoriteList(req)

	if err != nil {
		resp.BaseResp = pack.BuilBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp.BaseResp = pack.BuilBaseResp(errno.ParamErr)
	resp.VideoList = vedioList
	return resp, nil
}
