package qcloud

// 直接发送短信字段
type msgSmsField struct {
	//0:普通短信;1:营销短信（强调：要按需填值，不然会影响到业务的正常使用）
	Type interface{} `json:"type,omitempty"`
	// 短信内容，需要匹配审核通过的模板内容
	Msg string `json:"msg,omitempty"`
}

func newMsgSmsField(t int, msg string) *msgSmsField {
	return &msgSmsField{
		Type: t,
		Msg:  msg,
	}
}
