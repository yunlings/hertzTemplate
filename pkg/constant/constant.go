package constant

// HTTP状态码
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
)

// 业务错误码
const (
	// 通用错误码 (10000-19999)
	CodeSuccess            = 10000 // 成功
	CodeInternalError      = 10001 // 内部错误
	CodeInvalidParams      = 10002 // 参数错误
	CodeNotFound           = 10003 // 资源不存在
	CodeUnauthorized       = 10004 // 未授权
	CodeForbidden          = 10005 // 禁止访问
	CodeTooManyRequests    = 10006 // 请求过于频繁
	CodeServiceUnavailable = 10007 // 服务不可用
	CodeTimeout            = 10008 // 请求超时
	CodeConflict           = 10009 // 资源冲突
	CodeValidationFailed   = 10010 // 数据验证失败
	CodeDatabaseError      = 10011 // 数据库错误
	CodeRedisError         = 10012 // Redis错误
	CodeCacheError         = 10013 // 缓存错误
	CodeNetworkError       = 10014 // 网络错误
	CodeFileError          = 10015 // 文件操作错误
	CodeConfigError        = 10016 // 配置错误
	CodeEncryptionError    = 10017 // 加密错误
	CodeDecryptionError    = 10018 // 解密错误
	CodeSerializationError = 10019 // 序列化错误

	// 认证相关错误码 (20000-29999)
	CodeTokenInvalid        = 20001 // Token无效
	CodeTokenExpired        = 20002 // Token过期
	CodeTokenMalformed      = 20003 // Token格式错误
	CodeTokenNotFound       = 20004 // Token不存在
	CodeLoginFailed         = 20005 // 登录失败
	CodePasswordIncorrect   = 20006 // 密码错误
	CodeAccountLocked       = 20007 // 账户被锁定
	CodeAccountDisabled     = 20008 // 账户被禁用
	CodePermissionDenied    = 20009 // 权限不足
	CodeRoleNotFound        = 20010 // 角色不存在
	CodeRefreshTokenInvalid = 20011 // 刷新Token无效
	CodeCaptchaRequired     = 20012 // 需要验证码
	CodeCaptchaInvalid      = 20013 // 验证码错误
	CodeTwoFactorRequired   = 20014 // 需要二次验证
	CodeSessionExpired      = 20015 // 会话过期

	// 用户相关错误码 (30000-39999)
	CodeUserNotFound       = 30001 // 用户不存在
	CodeUserExists         = 30002 // 用户已存在
	CodeUsernameExists     = 30003 // 用户名已存在
	CodeEmailExists        = 30004 // 邮箱已存在
	CodePhoneExists        = 30005 // 手机号已存在
	CodeUserCreateFailed   = 30006 // 用户创建失败
	CodeUserUpdateFailed   = 30007 // 用户更新失败
	CodeUserDeleteFailed   = 30008 // 用户删除失败
	CodePasswordTooWeak    = 30009 // 密码强度不够
	CodePasswordSame       = 30010 // 新密码与旧密码相同
	CodeEmailNotVerified   = 30011 // 邮箱未验证
	CodePhoneNotVerified   = 30012 // 手机号未验证
	CodeProfileIncomplete  = 30013 // 个人信息不完整
	CodeAvatarUploadFailed = 30014 // 头像上传失败
	CodeUserStatusInvalid  = 30015 // 用户状态无效

	// 业务相关错误码 (40000-49999)
	CodeResourceNotFound     = 40001 // 资源不存在
	CodeResourceExists       = 40002 // 资源已存在
	CodeResourceCreateFailed = 40003 // 资源创建失败
	CodeResourceUpdateFailed = 40004 // 资源更新失败
	CodeResourceDeleteFailed = 40005 // 资源删除失败
	CodeResourceLocked       = 40006 // 资源被锁定
	CodeResourceExpired      = 40007 // 资源已过期
	CodeQuotaExceeded        = 40008 // 配额超限
	CodeOperationNotAllowed  = 40009 // 操作不被允许
	CodeDependencyExists     = 40010 // 存在依赖关系

	// 第三方服务错误码 (50000-59999)
	CodeThirdPartyError    = 50001 // 第三方服务错误
	CodePaymentFailed      = 50002 // 支付失败
	CodeSMSFailed          = 50003 // 短信发送失败
	CodeEmailFailed        = 50004 // 邮件发送失败
	CodeUploadFailed       = 50005 // 文件上传失败
	CodeDownloadFailed     = 50006 // 文件下载失败
	CodeAPICallFailed      = 50007 // API调用失败
	CodeWebhookFailed      = 50008 // Webhook调用失败
	CodeNotificationFailed = 50009 // 通知发送失败
	CodeSyncFailed         = 50010 // 数据同步失败
)

// 错误信息映射
var CodeMessages = map[int]string{
	// 通用错误
	CodeSuccess:            "操作成功",
	CodeInternalError:      "内部服务器错误",
	CodeInvalidParams:      "请求参数错误",
	CodeNotFound:           "资源不存在",
	CodeUnauthorized:       "未授权访问",
	CodeForbidden:          "禁止访问",
	CodeTooManyRequests:    "请求过于频繁，请稍后重试",
	CodeServiceUnavailable: "服务暂时不可用",
	CodeTimeout:            "请求超时",
	CodeConflict:           "资源冲突",
	CodeValidationFailed:   "数据验证失败",
	CodeDatabaseError:      "数据库操作失败",
	CodeRedisError:         "缓存服务错误",
	CodeCacheError:         "缓存操作失败",
	CodeNetworkError:       "网络连接错误",
	CodeFileError:          "文件操作失败",
	CodeConfigError:        "配置错误",
	CodeEncryptionError:    "数据加密失败",
	CodeDecryptionError:    "数据解密失败",
	CodeSerializationError: "数据序列化失败",

	// 认证相关错误
	CodeTokenInvalid:        "Token无效",
	CodeTokenExpired:        "Token已过期",
	CodeTokenMalformed:      "Token格式错误",
	CodeTokenNotFound:       "Token不存在",
	CodeLoginFailed:         "登录失败",
	CodePasswordIncorrect:   "密码错误",
	CodeAccountLocked:       "账户已被锁定",
	CodeAccountDisabled:     "账户已被禁用",
	CodePermissionDenied:    "权限不足",
	CodeRoleNotFound:        "角色不存在",
	CodeRefreshTokenInvalid: "刷新Token无效",
	CodeCaptchaRequired:     "需要验证码",
	CodeCaptchaInvalid:      "验证码错误",
	CodeTwoFactorRequired:   "需要二次验证",
	CodeSessionExpired:      "会话已过期",

	// 用户相关错误
	CodeUserNotFound:       "用户不存在",
	CodeUserExists:         "用户已存在",
	CodeUsernameExists:     "用户名已存在",
	CodeEmailExists:        "邮箱已存在",
	CodePhoneExists:        "手机号已存在",
	CodeUserCreateFailed:   "用户创建失败",
	CodeUserUpdateFailed:   "用户更新失败",
	CodeUserDeleteFailed:   "用户删除失败",
	CodePasswordTooWeak:    "密码强度不够",
	CodePasswordSame:       "新密码不能与旧密码相同",
	CodeEmailNotVerified:   "邮箱未验证",
	CodePhoneNotVerified:   "手机号未验证",
	CodeProfileIncomplete:  "个人信息不完整",
	CodeAvatarUploadFailed: "头像上传失败",
	CodeUserStatusInvalid:  "用户状态无效",

	// 业务相关错误
	CodeResourceNotFound:     "资源不存在",
	CodeResourceExists:       "资源已存在",
	CodeResourceCreateFailed: "资源创建失败",
	CodeResourceUpdateFailed: "资源更新失败",
	CodeResourceDeleteFailed: "资源删除失败",
	CodeResourceLocked:       "资源已被锁定",
	CodeResourceExpired:      "资源已过期",
	CodeQuotaExceeded:        "配额已超限",
	CodeOperationNotAllowed:  "操作不被允许",
	CodeDependencyExists:     "存在依赖关系，无法删除",

	// 第三方服务错误
	CodeThirdPartyError:    "第三方服务错误",
	CodePaymentFailed:      "支付失败",
	CodeSMSFailed:          "短信发送失败",
	CodeEmailFailed:        "邮件发送失败",
	CodeUploadFailed:       "文件上传失败",
	CodeDownloadFailed:     "文件下载失败",
	CodeAPICallFailed:      "API调用失败",
	CodeWebhookFailed:      "Webhook调用失败",
	CodeNotificationFailed: "通知发送失败",
	CodeSyncFailed:         "数据同步失败",
}

// GetMessage 根据错误码获取对应的消息
func GetMessage(code int) string {
	if message, exists := CodeMessages[code]; exists {
		return message
	}
	return "未知错误"
}

// SetMessage 设置错误码对应的消息（用于自定义消息）
func SetMessage(code int, message string) {
	CodeMessages[code] = message
}

// GetAllMessages 获取所有错误码消息映射
func GetAllMessages() map[int]string {
	return CodeMessages
}

// 常用常量
const (
	// 默认值
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100

	// 时间格式
	TimeFormat    = "2006-01-02 15:04:05"
	DateFormat    = "2006-01-02"
	TimeFormatISO = "2006-01-02T15:04:05Z07:00"

	// 缓存键前缀
	CacheKeyUser         = "user:"
	CacheKeyToken        = "token:"
	CacheKeyRefreshToken = "refresh_token:"
	CacheKeySession      = "session:"
	CacheKeyRateLimit    = "rate_limit:"
	CacheKeyConfig       = "config:"

	// 缓存过期时间
	CacheExpireShort  = 300   // 5分钟
	CacheExpireMedium = 1800  // 30分钟
	CacheExpireLong   = 3600  // 1小时
	CacheExpireDay    = 86400 // 1天

	// 用户状态
	UserStatusDisabled = 0 // 禁用
	UserStatusEnabled  = 1 // 启用

	// 性别
	GenderUnknown = 0 // 未知
	GenderMale    = 1 // 男
	GenderFemale  = 2 // 女

	// JWT相关
	JWTHeaderKey = "Authorization"
	JWTPrefix    = "Bearer "

	// 请求头
	HeaderRequestID    = "X-Request-ID"
	HeaderUserAgent    = "User-Agent"
	HeaderRealIP       = "X-Real-IP"
	HeaderForwardedFor = "X-Forwarded-For"
	HeaderContentType  = "Content-Type"

	// 内容类型
	ContentTypeJSON = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeXML  = "application/xml"
	ContentTypeText = "text/plain"

	// 环境
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)
