package pack

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
)

// UserInfo pack user info
func UserInfo(u *db.User, isFollow bool) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      isFollow,
	}
}
