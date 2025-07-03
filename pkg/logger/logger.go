package logger

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func InitLoggerZap() {
	// 设置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // 自定义时间格式为 ISO8601
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("log.file_name"),
		MaxSize:    viper.GetInt("log.max_size"),    // 一个文件最大可达 20M。
		MaxBackups: viper.GetInt("log.max_backups"), // 最多同时保存 5 个文件。
		MaxAge:     viper.GetInt("log.max_age"),     // 一个文件最多可以保存 7天。
		Compress:   viper.GetBool("log.compress"),   // 用 gzip 压缩。
	}

	enc := zapcore.NewJSONEncoder(encoderConfig)

	if viper.GetBool("log.stdout") == true {
		logger := hertzzap.NewLogger(hertzzap.WithCoreEnc(enc), hertzzap.WithCoreWs(zapcore.AddSync(os.Stdout)))
		logger.SetLevel(hlog.Level(2))
		hlog.SetLogger(logger)
	} else {
		logger := hertzzap.NewLogger(hertzzap.WithCoreEnc(enc))
		logger.SetOutput(lumberjackLogger)
		logger.SetLevel(hlog.Level(2))
		hlog.SetLogger(logger)
	}

}
func getRequestID(c *app.RequestContext) string {
	//return string(c.GetHeader(constant.HeaderRequestID))
	return c.Response.Header.Get("X-Request-ID")
}
func InfoX(c *app.RequestContext, format string, v ...interface{}) {

	logData := map[string]interface{}{
		"traceId": getRequestID(c),
		"custmsg": fmt.Sprintf(format, v...),
	}
	jsonData, _ := json.Marshal(logData)

	hlog.Infof("%s", jsonData)

}
