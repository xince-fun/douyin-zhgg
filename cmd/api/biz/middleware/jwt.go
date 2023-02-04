package middleware

import (
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/handler"
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/model/api"
	"ByteTech-7815/douyin-zhgg/cmd/api/biz/rpc"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{UserId: int64(claims["id"].(float64))}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req api.DouyinUserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			userId, err := rpc.LoginUser(context.Background(), &user.DouyinUserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			c.Set("USER_ID", userId)
			return userId, err
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, message string, time time.Time) {
			userId, _ := c.Get("USER_ID")
			handler.SendUserResponse(c, errno.Success, userId.(int64), message)
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			handler.SendResponse(c, errno.AuthorizationFailedErr)
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
}
