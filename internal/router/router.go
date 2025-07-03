package router

import (
	"context"
	"hertzTemplate/internal/controller/system"
	"hertzTemplate/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/requestid"
	"github.com/yunlings/gtools/grandom"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func CustomizedRegister(h *server.Hertz) {
	h.Use(middleware.AccessLog)
	h.Use(
		requestid.New(requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
			return grandom.GenerateUUIDWithoutHyphen()
		}),
			// 自定义 request id 响应头键值
			requestid.WithCustomHeaderStrKey("X-Request-ID"),
		),
	)
	h.GET("/ping", system.Ping)

}
