package jwt

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserName string `json:"username"`
	UserId   int64  `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(username string, userid int64) (string, error) {
	claims := &CustomClaims{
		UserName: username,
		UserId:   userid,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(consts.SecretKey))
}

// ParseToken 解析JWT
func ParseToken(signedToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(consts.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errno.TokenInvalidErr
}
