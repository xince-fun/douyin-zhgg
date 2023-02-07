package pack

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
)

func PublishInfo(v *db.Video, userMap map[int64]*db.User) *feed.Video {
	if v == nil {
		return nil
	}
	user := userMap[v.UserId]
	return &feed.Video{
		Id:            int64(v.ID),
		Author:        UserInfo(user, true),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    true,
		Title:         v.Title,
	}
}

func PublishList(video []*db.Video, userMap map[int64]*db.User) []*feed.Video {
	res := make([]*feed.Video, 0)
	if len(video) == 0 {
		return res
	}
	for _, v := range video {
		if m := PublishInfo(v, userMap); m != nil {
			res = append(res, m)
		}
	}
	return res
}
