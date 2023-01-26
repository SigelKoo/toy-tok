package core_feed

import (
	"context"
	"toy-tok/biz/model/core_feed"

	"github.com/pkg/errors"

	"github.com/cloudwego/hertz/pkg/app"
)

const LISTMAX int = 30

func GetVideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_feed.ToytokFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		_ = c.Error(errors.WithStack(err))
		return
	}
	statusCode := int32(0)
	statusMsg := ""
	videos := []*core_feed.Video{}
	latestTime := int64(0)
	resp := &core_feed.ToytokFeedResponse{
		StatusCode: &statusCode,
		StatusMsg:  &statusMsg,
		VideoList:  videos,
		NextTime:   &latestTime,
	}
	c.JSON(200, resp)
}
