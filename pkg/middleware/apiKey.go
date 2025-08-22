package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var validAccessKeys = map[string]bool{
	"ssl-Qqqxsar4NFJWXI2T": true,
}

func CheckApiKey(ctx context.Context, c *app.RequestContext) {
	apiKey := c.GetHeader("X-API-KEY")
	if isValid := validAccessKeys[string(apiKey)]; !isValid {
		c.JSON(consts.StatusForbidden, utils.H{"code": consts.StatusForbidden, "error": "invalid api key"})
		c.Abort()
		return
	}
	c.Next(ctx)

}
