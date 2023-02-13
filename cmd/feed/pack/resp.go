package pack

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"errors"
)

func BuildBaseResp(err error) *feed.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *feed.BaseResp {
	return &feed.BaseResp{StatusCode: err.ErrCode, StatsuMsg: &err.ErrMsg}
}
