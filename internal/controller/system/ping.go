package system

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
	//logger.InfoX(c, "这是自定义消息%s", c.Request.Header.Get("Host"))
	//response.Error(c, constant.CodeNotFound, nil)

}

// 自定义Http状态码
var statusCode = 200

// 设定Http状态码
func changeStatusCode() {
	statusCode = 999
}

// 定义存活检查接口
func CheckOnline(ctx context.Context, c *app.RequestContext) {
	c.JSON(statusCode, utils.H{
		"message": "online",
	})
}

// 定义下线接口
func Offline(ctx context.Context, c *app.RequestContext) {
	changeStatusCode()

	c.JSON(200, utils.H{
		"message": "offline",
	})
	time.Sleep(30 * time.Second)
}
