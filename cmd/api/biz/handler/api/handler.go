package api

import (
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
