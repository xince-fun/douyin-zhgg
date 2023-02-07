package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

const (
	addFollow    = 1
	cancelFollow = 2
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 检查要关注的用户是否存在
	user, err := db.QueryUserById(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return err
	}
	if len(user) == 0 {
		return errno.UserNotExistErr
	}
	if req.ActionType == addFollow {
		if db.IsFollowing(s.ctx, req.UserId, req.ToUserId) == false {
			return db.CreateFollow(s.ctx, req.UserId, req.ToUserId)
		}
	} else if req.ActionType == cancelFollow {
		if db.IsFollowing(s.ctx, req.UserId, req.ToUserId) == true {
			return db.DeleteFollow(s.ctx, req.UserId, req.ToUserId)
		}
	}
	return nil
}
