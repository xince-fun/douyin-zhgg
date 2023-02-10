package pack

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/feed"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
)

func VedioList(currentId int64, videoData []*db.Video, userMap map[int64]*db.User, favoriteMap map[int64]*db.Favorite, relationMap map[int64]*db.Follow) []*feed.Video {
	videoList := make([]*feed.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.User{
				Username:      "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.ID = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.ID)]
			if ok {
				isFavorite = true
			}
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}
		videoList = append(videoList, &feed.Video{
			Id: int64(video.ID),
			Author: &user.User{
				Id:            int64(videoUser.ID),
				Name:          videoUser.Username,
				FollowCount:   &videoUser.FollowCount,
				FollowerCount: &videoUser.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoList
}
