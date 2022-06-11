package models

type ReplySingle struct {
	EmitMsg  string `json:"emit_msg"`
	ReplyMsg string `json:"reply_msg"`
}

func (r *ReplySingle) Handle(m *Message) string {
	//如果这个指令有内置的其他作用 就先回复其他作用的

	//否则直接返回消息
	//先判断这个消息是不是在触发区间

	return r.ReplyMsg
}
