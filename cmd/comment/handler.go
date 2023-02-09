package main

import (
	"ByteTech-7815/douyin-zhgg/cmd/comment/pack"
	"ByteTech-7815/douyin-zhgg/cmd/comment/service"
	comment "ByteTech-7815/douyin-zhgg/kitex_gen/comment"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
	"unicode/utf8"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)
	if utf8.RuneCountInString(*req.CommentText) == 0 || req.VideoId == 0 {
		resp.BaseResp = (*comment.BaseResp)(pack.BuilBaseResp(errno.ParamErr))
		return resp, nil
	}
	comment, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.BaseResp = pack.BuilBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuilBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)
	if req.VideoId == 0 {
		resp.BaseResp = (*comment.BaseResp)(pack.BuilBaseResp(errno.ParamErr))
		return resp, nil
	}
	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuilBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuilBaseResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}
