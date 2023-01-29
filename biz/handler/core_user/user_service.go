package core_user

import (
	"context"
	"toy-tok/biz/dal/mysql"
	"toy-tok/biz/model/core_user"
	"toy-tok/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

const (
	successHTTPCode  = 200
	failedServerCode = 500
)

var (
	successCode     = int32(0)
	successMsg      = "register user success"
	bindFailedCode  = int32(1)
	bindFailedMsg   = "request bind and validate failed"
	gormFailedCode  = int32(2)
	gormFailedMsg   = "gorm create user failed"
	tokenFailedCode = int32(3)
	tokenFailedMsg  = "jwt create token failed"
	failedToken     = ""
	failedUserID    = int64(0)
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_user.ToytokUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		c.JSON(failedServerCode, &core_user.ToytokUserRegisterResponse{
			StatusCode: &bindFailedCode,
			StatusMsg:  &bindFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		})
		return
	}
	userID, err := mysql.CreateUser(ctx, *req.Username, *req.Password)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		c.JSON(failedServerCode, &core_user.ToytokUserRegisterResponse{
			StatusCode: &gormFailedCode,
			StatusMsg:  &gormFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		})
		return
	}
	token, err := util.CreateAccessToken(userID)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		c.JSON(failedServerCode, &core_user.ToytokUserRegisterResponse{
			StatusCode: &tokenFailedCode,
			StatusMsg:  &tokenFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		})
		return
	}
	resp := &core_user.ToytokUserRegisterResponse{
		StatusCode: &successCode,
		StatusMsg:  &successMsg,
		UserId:     &userID,
		Token:      &token,
	}
	c.JSON(successHTTPCode, resp)
}
