package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/comment"
	"context"
	"errors"
	"sync"
)

type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	var comment *comment.Comment
	var err error
	if req.ActionType == 1 {
		comment, err = s.CreateCommentAction(req)
	}

	if req.ActionType == 2 {
		comment, err = s.DeleteCommentAction(req)
	}
	return comment, err

}

func (s *CommentActionService) CreateCommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	videos, err := db.QueryVideoByVideoId(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("videoId not exists")
	}

	commentOri := &db.CommentOri{
		UserId:   *req.CommentId,
		VideoId:  req.VideoId,
		Contents: *req.CommentText,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	var user *db.User
	var commentErr, userErr error
	//创建评论记录，增加评论数
	go func() {
		defer wg.Done()
		err := db.CreateComment(s.ctx, commentOri)
		if err != nil {
			commentErr = err
			return
		}
	}()
	//获取用户信息
	go func() {
		defer wg.Done()
		users, err := db.QueryUserById(s.ctx, []int64{*req.CommentId})
		if err != nil {
			userErr = err
			return
		}
		user = users[0]
	}()
	wg.Wait()
	if commentErr != nil {
		return nil, commentErr
	}
	if userErr != nil {
		return nil, userErr
	}
	comment := pack.CommentInfo(commentOri, user)
	return comment, nil
}

func (s *CommentActionService) DeleteCommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	videos, err := db.QueryVideoByVideoId(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("videoId not exist")
	}
	comments, err := db.QueryCommentByCommentIds(s.ctx, []int64{*req.CommentId})
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return nil, errors.New("commentId not exist")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var commentOri *db.CommentOri
	var userRaw *db.User
	var commentErr, userErr error
	//删除评论记录并减少视频评论数
	go func() {
		defer wg.Done()
		commentOri, err = db.DeleteComment(s.ctx, *req.CommentId)
		if err != nil {
			commentErr = err
			return
		}
	}()
	//获取用户信息
	go func() {
		defer wg.Done()
		users, err := db.QueryUserById(s.ctx, []int64{*req.CommentId})
		if err != nil {
			userErr = err
			return
		}
		userRaw = users[0]
	}()
	wg.Wait()
	if commentErr != nil {
		return nil, commentErr
	}
	if userErr != nil {
		return nil, userErr
	}

	comment := pack.CommentInfo(commentOri, userRaw)
	return comment, nil
}
