package pack

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/favorite"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"errors"
)

//BuildCommentBaseResp build comment baseResp from error
func BuilBaseResp(err error) *favorite.BaseResp {
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

func baseResp(err errno.ErrNo) *favorite.BaseResp {
	return &favorite.BaseResp{StatusCode: err.ErrCode, StatsuMsg: &err.ErrMsg}
}
