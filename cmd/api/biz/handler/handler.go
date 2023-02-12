package handler

import (
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	Data       []interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data ...interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Data:       data,
	})
}

type UserResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

// SendUserResponse pack user response
func SendUserResponse(c *app.RequestContext, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     userId,
		Token:      token,
	})
}

type UserInfoResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       interface{} `json:"user"`
}

// SendUserInfoResponse pack user info response
func SendUserInfoResponse(c *app.RequestContext, err error, user interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       user,
	})
}

func SendRelationActionResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

type FollowListResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	FollowList []interface{} `json:"follow_list"`
}

func SendRelationListResponse(c *app.RequestContext, err error, followList []*user.User) {
	Err := errno.ConvertErr(err)
	// 把对用户的指针数组转换成interface数组
	var followListInterface []interface{}
	for _, v := range followList {
		followListInterface = append(followListInterface, v)
	}
	c.JSON(consts.StatusOK, FollowListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		FollowList: followListInterface,
	})
}

type VideoListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  interface{} `json:"video_list,omitempty"`
}

func SendVideoListResponse(c *app.RequestContext, err error, videoList interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, VideoListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
	})
}

type FeedInfoResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Feed       interface{} `json:"feed"`
}

func SendFeedResponse(c *app.RequestContext, err error, feed interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, FeedInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		Feed:       feed,
	})
}
