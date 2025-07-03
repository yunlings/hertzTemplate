package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() error {

	viper.SetConfigName("config")   //指定配置文件名称
	viper.SetConfigType("yaml")     //指定配置文件类型(专用于从远程获取配置信息时指定配置类型)
	viper.AddConfigPath("./config") //指定查找配置文件路径
	//viper.SetConfigFile(cnf)      //命令行参数
	err := viper.ReadInConfig() //读取配置
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed ,err:%v\n", err)
		return err
	}

	viper.WatchConfig()                            //监视配置配置变更
	viper.OnConfigChange(func(in fsnotify.Event) { //回调函数
		fmt.Println("配置文件被修改了")
		err := viper.ReadInConfig() //读取配置
		if err != nil {
			//读取配置信息失败
			fmt.Printf("viper.ReadInConfig() failed ,err:%v\n", err)
			return
		}
	})
	return nil
}
