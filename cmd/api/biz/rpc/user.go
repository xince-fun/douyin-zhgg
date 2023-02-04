package rpc

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user/userservice"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"ByteTech-7815/douyin-zhgg/pkg/middleware"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		consts.UserServiceName,
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
	userClient = c
}

// RegisterUser register user info
func RegisterUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return 0, "", err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return resp.UserId, resp.Token, nil
}

// LoginUser check user info
func LoginUser(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, string, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return 0, "", err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, "", errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return resp.UserId, resp.Token, nil
}

// UserInfo get user info
func UserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.User, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatsuMsg)
	}
	return resp.User, nil
}
