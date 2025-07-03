package response

import (
	"hertzTemplate/pkg/constant"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`                 // 业务状态码
	Message   string      `json:"message"`              // 响应消息
	Data      interface{} `json:"data,omitempty"`       // 响应数据
	Timestamp int64       `json:"timestamp"`            // 时间戳
	RequestID string      `json:"request_id,omitempty"` // 请求ID
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Code      int                    `json:"code"`                 // 业务状态码
	Message   string                 `json:"message"`              // 错误消息
	Errors    map[string]interface{} `json:"errors,omitempty"`     // 详细错误信息
	Timestamp int64                  `json:"timestamp"`            // 时间戳
	RequestID string                 `json:"request_id,omitempty"` // 请求ID
	Path      string                 `json:"path,omitempty"`       // 请求路径
	Method    string                 `json:"method,omitempty"`     // 请求方法
}

// getRequestID 获取请求ID
func getRequestID(c *app.RequestContext) string {
	//return string(c.GetHeader(constant.HeaderRequestID))
	return c.Response.Header.Get("X-Request-ID")
}

// Success 成功响应
func Success(c *app.RequestContext, data interface{}) {
	response := &Response{
		Code:      constant.CodeSuccess,
		Message:   constant.GetMessage(constant.CodeSuccess),
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	}
	c.JSON(consts.StatusOK, response)
}

// SuccessWithMessage 成功响应（自定义消息）
func SuccessWithMessage(c *app.RequestContext, message string, data interface{}) {
	response := &Response{
		Code:      constant.StatusOK,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	}
	c.JSON(consts.StatusOK, response)
}

func Error(c *app.RequestContext, code int, err map[string]interface{}) {
	response := &ErrorResponse{
		Code:      code,
		Message:   constant.GetMessage(code),
		Errors:    err,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	}
	c.JSON(consts.StatusOK, response)
}
