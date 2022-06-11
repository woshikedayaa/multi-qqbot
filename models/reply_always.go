package models

type ReplyAlways struct {
	EnableObject int64         `json:"enable"`
	Reply        []ReplySingle `json:"reply"`
}
