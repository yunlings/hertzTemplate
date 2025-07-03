package middleware

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func AccessLog(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	c.Next(ctx)

	latency := time.Since(start)
	// 构造结构化日志字段
	logData := map[string]interface{}{
		//"level":     "info",
		//"time":      time.Now().Format("2006-01-02T15:04:05.000+0800"),
		"status":    c.Response.StatusCode(),
		"cost":      latency.String(),
		"method":    string(c.Request.Header.Method()),
		"full_path": string(c.Request.URI().PathOriginal()),
		"client_ip": c.ClientIP(),
		"host":      string(c.Request.Header.Peek("Host")),
		"traceId":   c.Response.Header.Get("X-Request-ID"),
	}
	jsonData, _ := json.Marshal(logData)
	hlog.CtxInfof(ctx, "%s", jsonData)
}
