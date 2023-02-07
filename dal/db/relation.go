package db

import (
	"ByteTech-7815/douyin-zhgg/pkg/consts"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	FollowerID int64 `json:"follower_id" gorm:"index;type:bigint(20);not null"`
	FolloweeID int64 `json:"followee_id" gorm:"index;type:bigint(20);not null"`
}

func TableName() string {
	return consts.RelationTableName
}

func CreateFollow(ctx context.Context, followerID int64, followeeID int64) error {
	follow := &Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
	return DB.WithContext(ctx).Create(&follow).Error
}

func DeleteFollow(ctx context.Context, followerID int64, followeeID int64) error {
	follow := &Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
	return DB.WithContext(ctx).Delete(&follow).Error
}

// GetFansID 返回粉丝id列表
func GetFansID(ctx context.Context, followeeID int64) ([]int64, error) {
	follows := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("followee_id = ?", followeeID).Find(&follows).Error; err != nil {
		return nil, err
	}

	res := make([]int64, 0)
	for _, follow := range follows {
		res = append(res, follow.FollowerID)
	}
	return res, nil
}

// GetFollowingIDs 返回关注用户的id列表
func GetFollowingIDs(ctx context.Context, userID int64) ([]int64, error) {
	follows := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("follower_id = ?", userID).Find(&follows).Error; err != nil {
		return nil, err
	}

	res := make([]int64, 0)
	for _, follow := range follows {
		res = append(res, follow.FolloweeID)
	}
	return res, nil
}

// GetFollowingUsers 返回关注用户列表
func GetFollowingUsers(ctx context.Context, userID int64) ([]*User, error) {
	followingIDs, err := GetFollowingIDs(ctx, userID)
	if err != nil {
		klog.Error("error occurred when get following ids", err)
	}
	return QueryUserById(ctx, followingIDs)
}

// GetFollowCount 返回关注数
func GetFollowCount(ctx context.Context, followerID int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("follower_id = ?", followerID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetFollowerCount 返回粉丝数
func GetFollowerCount(ctx context.Context, followeeID int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("followee_id = ?", followeeID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// IsFollowing 判断是否关注
func IsFollowing(ctx context.Context, followerID int64, followeeID int64) bool {
	follow := &Follow{}
	if err := DB.WithContext(ctx).Where("follower_id = ? AND followee_id = ?", followerID, followeeID).First(&follow).Error; err != nil {
		return false
	}
	return true
}

// GetFriendsID 返回好友id列表
// 两个用户互相关注即为好友
func GetFriendsID(ctx context.Context, userID int64) ([]int64, error) {
	// 获取关注列表
	follows := make([]*Follow, 0)
	if err := DB.WithContext(ctx).Where("follower_id = ?", userID).Find(&follows).Error; err != nil {
		return nil, err
	}

	res := make([]int64, 0)
	// 遍历关注列表，判断我的关注是否关注了我
	for _, follow := range follows {
		if IsFollowing(ctx, follow.FolloweeID, userID) {
			res = append(res, follow.FolloweeID)
		}
	}
	return res, nil
}
