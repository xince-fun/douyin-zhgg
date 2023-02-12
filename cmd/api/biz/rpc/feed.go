package rpc

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed/feedservice"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/middleware"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := feedservice.NewClient(
		consts.FeedServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

func GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	resp, err := feedClient.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return resp, nil
}
