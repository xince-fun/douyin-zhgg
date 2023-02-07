package rpc

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/kitex_gen/publish"
	"ByteTech-7815/douyin-zhgg/kitex_gen/publish/publishservice"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/middleware"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		consts.PublishServiceName,
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
	publishClient = c
}

// Publish video
func Publish(ctx context.Context, req *publish.DouyinPublishActionRequest) error {
	resp, err := publishClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return nil
}

func PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) ([]*feed.Video, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return resp.VideoList, nil

}
