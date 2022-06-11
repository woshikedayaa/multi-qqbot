package inits

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multi-qqbot/controllors"
	"multi-qqbot/log"
)

func Route() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	//处理路由
	Dir := viper.GetStringSlice("user.qqRelativePath")
	Names := viper.GetStringSlice("user.qqName")
	if len(Names) != len(Dir) {
		log.Fatal("路径，人数 不一致 请检查 config/config.toml")
	}
	for i, v := range Dir {
		log.Info("注册路由", zap.String("名字", Names[i]), zap.String("路径", v))
		r.POST(v, controllors.MessageHandler(Names[i]))
	}
	//
	return r
}
