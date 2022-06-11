package inits

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multi-qqbot/log"
	"multi-qqbot/models"
	"multi-qqbot/requester"
	"multi-qqbot/utils"
	"strconv"
)

func InitReadReplies() {

	log.Info("开始读取消息文件")
	defer log.Sync()

	//初始化map
	for _, v := range viper.GetStringSlice("user.qqNumber") {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Error(err)
		}
		//
		models.TimeMaps[i] = make(map[string]*models.ReplyTime)
		models.AlwaysMaps[i] = make(map[string]*models.ReplyAlways)
	}

	//Time消息
	for _, v := range utils.ReadJsonFromFileDir("./config/replies/time") {
		trp := &models.ReplyTime{}
		err := json.Unmarshal(v, trp)
		if err != nil {
			log.Error(err)
		}

		for _,value :=range trp.Reply{
			log.Info("Time消息", zap.String("触发消息", value.EmitMsg), zap.String("回复消息", value.ReplyMsg))
			models.TimeMaps[trp.EnableObject][value.EmitMsg] = trp
		}

		//log.Info("Time消息", zap.String("触发消息", trp.Reply[i].EmitMsg), zap.String("回复消息", trp.Reply[i].ReplyMsg))
		//models.TimeMaps[trp.EnableObject][trp.Reply[i].EmitMsg] = trp
	}

	//Always消息
	for _, v := range utils.ReadJsonFromFileDir("./config/replies/always") {
		arp := &models.ReplyAlways{}
		err := json.Unmarshal(v, arp)
		if err != nil {
			log.Error(err)
		}

		for _,value :=range arp.Reply{
			log.Info("Always消息", zap.String("触发消息", value.EmitMsg), zap.String("回复消息", value.ReplyMsg))
			models.AlwaysMaps[arp.EnableObject][value.EmitMsg] = arp
		}

	}

	//Once消息
	for _, v := range utils.ReadJsonFromFileDir("./config/replies/once") {
		orp := &models.ReplyOnce{}
		err := json.Unmarshal(v, orp)
		if err != nil {
			log.Error(err)
		}
		log.Info("Once消息", zap.String("触发时间", orp.EmitTime), zap.String("回复消息", orp.Reply.ReplyMsg))
		nowTamp := carbon.Now().Timestamp()
		futher := carbon.Parse(utils.CreateFullTime(orp.EmitTime))
		after := int64(0)
		if nowTamp >= futher.Timestamp() {
			after = futher.Tomorrow().Timestamp() - nowTamp
		}
		if nowTamp < futher.Timestamp() {
			after = futher.Timestamp() - nowTamp
		}
		log.Info("Once消息距离下一次发送 ", zap.Int64("time", after))
		utils.CreateNewDailyTimer(after, func() {
			//依次发送消息
			for _, v := range viper.GetStringSlice("user.onceGroupList") {
				id, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					log.Error(err)
				}
				requester.SendMessage(orp.Reply.ReplyMsg, &models.Message{
					MessageType: "group",
					GroupId:     id,
				})
			}
		})
	}
	log.Info("读取消息文件完成", zap.Any("Time", models.TimeMaps), zap.Any("Always", models.AlwaysMaps))
}
