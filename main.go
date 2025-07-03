package main

import (
	"fmt"
	"hertzTemplate/internal/router"
	"hertzTemplate/pkg/config"
	"hertzTemplate/pkg/logger"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
	"github.com/spf13/viper"
)

func main() {

	//加载配置
	if err := config.InitConfig(); err != nil {
		fmt.Printf("配置加载失败:err=%v\n", err)
	}
	//日志加载
	logger.InitLoggerZap()

	h := server.Default(server.WithExitWaitTime(10*time.Second), server.WithHostPorts(fmt.Sprintf(":%s", viper.GetString("server.port"))), server.WithTracer(prometheus.NewServerTracer(":9093", "/metrics", prometheus.WithEnableGoCollector(true))))

	router.CustomizedRegister(h)

	h.Spin()
}
