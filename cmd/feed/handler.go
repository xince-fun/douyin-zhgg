package main

import (
	"ByteTech-7815/douyin-zhgg/cmd/feed/pack"
	"ByteTech-7815/douyin-zhgg/cmd/feed/service"
	feed "ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetUserFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	resp = new(feed.DouyinFeedResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp.VedioList, err = service.NewGetUserFeedService(ctx).GetUserFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	*resp.NextTime = 0
	return resp, nil
}
