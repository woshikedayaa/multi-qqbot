package controllors

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"go.uber.org/zap"
	"multi-qqbot/log"
	"multi-qqbot/models"
	"multi-qqbot/requester"
	"multi-qqbot/utils"
)

//MessageHandler 消息处理器
func MessageHandler(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//s := utils.ReadJsonFromReaderByString(c.Request.Body)
		message := models.NewMessageByJson(utils.ReadJsonFromReaderByBytes(c.Request.Body))
		//log.Info(fmt.Sprintf("%#v",message))
		if message.PostType == "message" {
			//日志记录
			log.Info(name,
				zap.String("收到", message.Message),
				zap.Int64("时间", message.Time),
				zap.String("消息类型", message.MessageType),
				zap.Int64("发送来源", message.UserId),
			)
			defer log.Sync()
			//开始处理消息并且回复
			mapType, TimeRep, AlwaRep := checkInAnyMap(message)
			if mapType == "" {
				return
			}
			switch mapType {

			case models.TimeMapFlag:
				{

					if carbon.Parse(utils.CreateFullTime(TimeRep.EmitStartTime)).Timestamp() < carbon.Now().Timestamp() && carbon.Parse(utils.CreateFullTime(TimeRep.EmitEndTime)).Timestamp() > carbon.Now().Timestamp() {

						for _, v := range TimeRep.Reply {
							if v.EmitMsg == message.Message {
								log.Info("触发消息"+message.Message)
								requester.SendMessage(v.ReplyMsg, message)
							}
						}

					}
				}
			case models.AlwaysMapFlag:
				{
					for _, v := range AlwaRep.Reply {
						if v.EmitMsg == message.Message {
							log.Info("触发消息"+message.Message)
							requester.SendMessage(v.ReplyMsg, message)
						}
					}
				}

			}
			//
		}
	}
}

func checkInAnyMap(msg *models.Message) (MapType string, tRep *models.ReplyTime, aRep *models.ReplyAlways) {
	if trep, existTime := models.TimeMaps[msg.SelfId][msg.Message]; existTime == true {
		return models.TimeMapFlag, trep, nil
	}

	if arep, exist := models.AlwaysMaps[msg.SelfId][msg.Message]; exist == true {
		return models.AlwaysMapFlag, nil, arep
	}

	return "", nil, nil
}

//MiddleWareCheck 废弃 没啥用
func MiddleWareCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() != 200 {
			log.Errors("出现错误！", zap.Any("Context", c))
		}
		return
	}
}
