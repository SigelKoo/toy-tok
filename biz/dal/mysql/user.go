package mysql

import (
	"context"
	"time"
	"toy-tok/biz/model/gorm_model"
	"toy-tok/util"
)

func CreateUser(ctx context.Context, userName, password string) (int64, error) {
	passwordSHA := util.GetKeccakStr(password)
	newUserInfo := &gorm_model.UserInfo{
		Name:          userName,
		UserName:      userName,
		Password:      passwordSHA,
		FollowCount:   0,
		FollowerCount: 0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	insertRes := DB.WithContext(ctx).Model(&gorm_model.UserInfo{}).Create(newUserInfo)
	if insertRes.Error != nil {
		return newUserInfo.ID, insertRes.Error
	}
	return newUserInfo.ID, nil
}
