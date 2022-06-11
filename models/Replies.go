package models

var (
	TimeMapFlag   = "time"
	TimeMaps      = make(map[int64]map[string]*ReplyTime)
	AlwaysMapFlag = "always"
	AlwaysMaps    = make(map[int64]map[string]*ReplyAlways)
)
