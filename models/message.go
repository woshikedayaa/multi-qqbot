package models

import (
	"encoding/json"
	"multi-qqbot/log"
	"strconv"
)

// Message 私聊可以
type Message struct {
	Time        int64      `json:"time"`
	SelfId      int64      `json:"self_id"`
	PostType    string     `json:"post_type"`
	MessageType string     `json:"message_type"`
	SubType     string     `json:"sub_type"`
	TempSource  int        `json:"temp_source"`
	MessageId   int        `json:"message_id"`
	GroupId     int64      `json:"group_id"`
	UserId      int64      `json:"user_id"`
	Message     string     `json:"message"`
	RawMessage  string     `json:"raw_message"`
	Font        int        `json:"font"`
	Sender      SenderType `json:"sender"`
}
type PrivateMessage Message
type GroupMessage Message

//SenderType 发送者类型
type SenderType struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
}

func (m *Message) GetUserIdString() string {
	if m.UserId == 0 {
		log.Errors("未初始化的消息")
		return ""
	}
	return strconv.FormatInt(m.UserId, 10)
}

func (s *SenderType) GetUserIdString() string {
	if s.UserId == 0 {
		log.Errors("未初始化的sender")
		return ""
	}
	return strconv.FormatInt(s.UserId, 10)
}

func NewMessage() *Message {
	return &Message{
		Time:        0,
		SelfId:      0,
		PostType:    "",
		MessageType: "group",
		SubType:     "",
		TempSource:  0,
		MessageId:   0,
		UserId:      0,
		Message:     "",
		RawMessage:  "",
		Font:        0,
		Sender: SenderType{
			UserId:   0,
			Nickname: "",
			Sex:      "unknow",
			Age:      0,
		},
	}
}

func NewMessageByJson(bs []byte) *Message {
	message := NewMessage()
	err := json.Unmarshal(bs, message)
	if err != nil {
		log.Error(err)
	}
	return message
}
