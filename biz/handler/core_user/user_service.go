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
	successCode          = int32(0)
	registerSuccessMsg   = "register user success"
	loginSuccessMsg      = "user login success"
	bindFailedCode       = int32(1)
	bindFailedMsg        = "request bind and validate failed"
	gormCreateFailedCode = int32(2)
	gormCreateFailedMsg  = "gorm create user failed"
	tokenFailedCode      = int32(3)
	tokenFailedMsg       = "jwt create token failed"
	passwordFailedCode   = int32(4)
	passwordFailedMsg    = "user password error"
	gormFindFailedCode   = int32(5)
	gormFindFailedMsg    = "gorm find user failed"
	failedToken          = ""
	failedUserID         = int64(0)
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_user.ToytokUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &bindFailedCode,
			StatusMsg:  &bindFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	userID, err := mysql.CreateUser(ctx, *req.Username, *req.Password)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &gormCreateFailedCode,
			StatusMsg:  &gormCreateFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	token, err := util.CreateAccessToken(userID)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &tokenFailedCode,
			StatusMsg:  &tokenFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	successResp := &core_user.ToytokUserRegisterResponse{
		StatusCode: &successCode,
		StatusMsg:  &registerSuccessMsg,
		UserId:     &userID,
		Token:      &token,
	}
	c.JSON(successHTTPCode, successResp.String())
	return
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_user.ToytokUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &bindFailedCode,
			StatusMsg:  &bindFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	userID, err := mysql.FindUserPasswordIsCorrect(ctx, *req.Username, *req.Password)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &gormFindFailedCode,
			StatusMsg:  &gormFindFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	if userID == 0 {
		failedResp := &core_user.ToytokUserLoginResponse{
			StatusCode: &passwordFailedCode,
			StatusMsg:  &passwordFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(successHTTPCode, failedResp.String())
		return
	}
	token, err := util.CreateAccessToken(userID)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		failedResp := &core_user.ToytokUserRegisterResponse{
			StatusCode: &tokenFailedCode,
			StatusMsg:  &tokenFailedMsg,
			UserId:     &failedUserID,
			Token:      &failedToken,
		}
		c.JSON(failedServerCode, failedResp.String())
		return
	}
	successResp := &core_user.ToytokUserLoginResponse{
		StatusCode: &successCode,
		StatusMsg:  &loginSuccessMsg,
		UserId:     &userID,
		Token:      &token,
	}
	c.JSON(successHTTPCode, successResp.String())
	return
}
