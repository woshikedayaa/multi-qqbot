package requester

import (
	"encoding/json"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"multi-qqbot/log"
	"multi-qqbot/models"
	"net/http"
	"strings"
)

func SendMessage(msg string, proto *models.Message) {
	//处理消息
	//将消息转为json格式
	bs, err := json.Marshal(struct {
		MessageType string `json:"message_type"`
		UserId      int64  `json:"user_id"`
		GroupId     int64  `json:"group_id"`
		Message     string `json:"message"`
	}{UserId: proto.UserId, GroupId: proto.GroupId, Message: msg, MessageType: proto.MessageType})
	if err != nil {
		log.Error(err)
	}
	//发送消息
	resp, err := http.Post(viper.GetString("baseUrl")+"/send_msg", "application/json", strings.NewReader(string(bs)))

	if err != nil {
		log.Error(err)
		return
	}

	if resp.StatusCode != 200 {
		log.Warn("发送消息失败", zap.String("消息", string(bs)), zap.Any("响应", resp))
		return
	}

	log.Info("发送消息成功", zap.ByteString("消息json", bs))
	return
}
