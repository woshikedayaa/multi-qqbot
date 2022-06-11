package inits

import (
	"github.com/spf13/viper"
	"multi-qqbot/log"
)

func InitConfig() {
	log.Info("读取配置文件 config/config.toml")
	//主要的配置文件
	viper.AddConfigPath("config")
	viper.SetConfigFile("config.toml")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	//读配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	//存一些全局变量
	viper.Set("baseUrl", "http://127.0.0.1:"+viper.GetString("core.botPort"))
}
