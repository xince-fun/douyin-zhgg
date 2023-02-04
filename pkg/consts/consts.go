package consts

import "time"

const (
	MySQLDefaultDSN = "zhgg:zhgg@tcp(127.0.0.1:33069)/zhgg-dy?charset=utf8mb4&parseTime=true&loc=Local"
	UserServiceAddr = ":9000" //User服务地址

	ETCDAddress = "127.0.0.1:2379"

	UserTableName     = "user"
	VideoTableName    = "video"
	CommentTableName  = "comment"
	FavoriteTableName = "favorite"
	RelationTableName = "relation"

	// jwt
	SecretKey           = "secret key"
	IdentityKey         = "id"
	TokenExpireDuration = time.Hour * 24

	// service name
	ApiServiceName  = "api"
	UserServiceName = "user"

	TCP = "tcp"
)
