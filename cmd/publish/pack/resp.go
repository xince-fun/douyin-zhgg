package pack

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/publish"
	"ByteTech-7815/douyin-zhgg/pkg/errno"

	"errors"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *publish.BaseResp {
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

func baseResp(err errno.ErrNo) *publish.BaseResp {
	return &publish.BaseResp{StatusCode: err.ErrCode, StatsuMsg: &err.ErrMsg}
}
