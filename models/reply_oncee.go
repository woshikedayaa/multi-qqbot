package models

type ReplyOnce struct {
	EnableObject int64       `json:"enable"`
	EmitTime     string      `json:"emitTime"`
	Reply        ReplySingle `json:"reply"`
}
