package pack

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/kitex_gen/comment"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/consts"
)

func CommentInfo(commentOri *db.CommentOri, userdb *db.User) *comment.Comment {
	comment := &comment.Comment{
		Id: int64(commentOri.ID),
		User: &user.User{
			Id:            int64(userdb.ID),
			Name:          userdb.Username,
			FollowCount:   &userdb.FollowCount,
			FollowerCount: &userdb.FollowerCount,
			IsFollow:      false,
		},
		Content:    commentOri.Contents,
		CreateDate: commentOri.UpdatedAt.Format(consts.TimeFormat),
	}
	return comment
}

func CommentList(currentId int64, comments []*db.CommentOri, userMap map[int64]*db.User, followMap map[int64]*db.Follow) []*comment.Comment {
	commentList := make([]*comment.Comment, 0)
	for _, commentOri := range comments {
		commentUser, ok := userMap[commentOri.UserId]
		if !ok {
			commentUser = &db.User{
				Username:      "unknown user",
				FollowCount:   0,
				FollowerCount: 0,
			}
			commentUser.ID = 0
		}
		var isFollow bool = false
		if currentId != -1 {
			_, ok := followMap[commentOri.UserId]
			if ok {
				isFollow = true
			}
		}

		commentList = append(commentList, &comment.Comment{
			Id: int64(commentUser.ID),
			User: &user.User{
				Id:            int64(commentUser.ID),
				Name:          commentUser.Username,
				FollowCount:   &commentUser.FollowCount,
				FollowerCount: &commentUser.FollowerCount,
				IsFollow:      isFollow,
			},
			Content:    commentOri.Contents,
			CreateDate: commentOri.CreatedAt.Format(consts.TimeFormat),
		})
	}
	return commentList
}
