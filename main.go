package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multi-qqbot/inits"
	"multi-qqbot/log"
	"multi-qqbot/utils"
)

func main() {
	//
	log.Info(utils.ProjectName + " 开始服务")
	gin.SetMode(gin.ReleaseMode)
	//读取配置文件
	inits.InitConfig()
	inits.InitReadReplies()
	//配置路由
	route := inits.Route()
	//启动服务
	log.Warn("服务器将启动在:127.0.0.1:" + viper.GetString("core.serverPort"))
	err := route.Run(":" + viper.GetString("core.serverPort"))
	if err != nil {
		log.Fatal("启动服务失败！", zap.String("错误", err.Error()))
	}
	//
}
