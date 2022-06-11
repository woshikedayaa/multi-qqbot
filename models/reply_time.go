package models

type ReplyTime struct {
	EnableObject  int64         `json:"enable"`
	EmitStartTime string        `json:"emitStartTime"`
	EmitEndTime   string        `json:"emitEndTime"`
	Reply         []ReplySingle `json:"reply"`
}
