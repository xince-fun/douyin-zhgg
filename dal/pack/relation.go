package pack

import "ByteTech-7815/douyin-zhgg/dal/db"
import "ByteTech-7815/douyin-zhgg/kitex_gen/user"

func FollowList(users []*db.User) []*user.User {
	if users == nil {
		return nil
	}
	list := make([]*user.User, 0, len(users))
	for _, u := range users {
		list = append(list, UserInfo(u, true))
	}
	return list
}
