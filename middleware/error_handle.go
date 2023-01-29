package middleware

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/pkg/errors"
)

func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)
	if len(c.Errors) == 0 {
		return
	}
	hertzErr := c.Errors[0]
	err := hertzErr.Unwrap()
	logger.CtxErrorf(ctx, "%+v", err)
	err = errors.Unwrap(err)
	c.JSON(500, utils.H{
		"code":    500,
		"message": err.Error(),
	})
}
