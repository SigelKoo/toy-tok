package mysql

import (
	"context"
	"time"
	"toy-tok/biz/model/gorm_model"
	"toy-tok/util"

	"github.com/pkg/errors"
)

func FindUserIsExists(ctx context.Context, userName string) (bool, error) {
	existsUser := &gorm_model.UserInfo{}
	// 可以用Count改进
	queryRes := DB.WithContext(ctx).Model(&gorm_model.UserInfo{}).Where("user_name = ?", userName).Find(existsUser)
	if queryRes.Error != nil {
		return false, queryRes.Error
	}
	resNum := queryRes.RowsAffected
	if resNum == 0 {
		return false, nil
	}
	return true, nil
}

func CreateUser(ctx context.Context, userName, password string) (int64, error) {
	isExists, err := FindUserIsExists(ctx, userName)
	if err != nil {
		return 0, err
	}
	if isExists {
		return 0, errors.New("user already exists")
	}
	passwordSHA := util.GetPasswordHash(password)
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

func FindUserPasswordIsCorrect(ctx context.Context, userName, password string) (int64, error) {
	isExists, err := FindUserIsExists(ctx, userName)
	if err != nil {
		return 0, err
	}
	if !isExists {
		return 0, errors.New("user not exists")
	}
	passwordSHA := util.GetPasswordHash(password)
	existsUser := &gorm_model.UserInfo{}
	queryRes := DB.WithContext(ctx).Model(&gorm_model.UserInfo{}).Where("user_name = ? and password = ?", userName, passwordSHA).Find(&existsUser)
	if queryRes.Error != nil {
		return 0, queryRes.Error
	}
	resNum := queryRes.RowsAffected
	if resNum == 0 {
		return 0, nil
	}
	return existsUser.ID, nil
}
