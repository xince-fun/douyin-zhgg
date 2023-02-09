package pack

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/comment"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"errors"
)

//BuildCommentBaseResp build comment baseResp from error
func BuilBaseResp(err error) *comment.BaseResp {
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

func baseResp(err errno.ErrNo) *comment.BaseResp {
	return &comment.BaseResp{StatusCode: err.ErrCode, StatsuMsg: &err.ErrMsg}
}
