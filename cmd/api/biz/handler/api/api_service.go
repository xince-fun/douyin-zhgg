// Code generated by hertz generator.

package api

import (
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/handler"
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/middleware"
	api "ByteTech-7815/douyin-zhgg/cmd/api/biz/model/api"
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/rpc"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUserFeed .
// @router /douyin/feed/ [GET]
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	userId, token, err := rpc.RegisterUser(context.Background(), &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	handler.SendUserResponse(c, errno.Success, userId, token)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}

	user, err := rpc.UserInfo(ctx, &user.DouyinUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	handler.SendUserInfoResponse(c, errno.Success, user)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinPublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinPublishListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinCommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router /douyin/relation/action/ [GET]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	userId, err := middleware.JwtMiddleware.Authenticator(ctx, c)
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	err = rpc.RelationAction(ctx, &relation.DouyinRelationActionRequest{
		UserId:     userId.(int64),
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		handler.SendResponse(c, errno.ConvertErr(err))
		return
	}
	handler.SendResponse(c, errno.Success)
}

// RelationFollowList .
// @router /douyin/relatioin/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}
