package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/comment"
	"context"
	"errors"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService new CommentListService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

// CommentList get comment list info
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	//验证视频Id是否存在
	videos, err := db.QueryVideoByVideoId(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("video not exist")
	}

	//获取一系列评论信息
	comments, err := db.QueryCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	//获取评论信息的用户id
	userIds := make([]int64, 0)
	for _, comment := range comments {
		userIds = append(userIds, comment.UserId)
	}

	//获取一系列用户信息
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	commentList := pack.CommentList(comments, userMap)
	return commentList, nil
}
