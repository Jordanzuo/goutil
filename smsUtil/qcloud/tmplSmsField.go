package qcloud

// 发送模板短信字段
type tmplSmsField struct {
	// 签名 (前缀)
	Sign string `json:"sign,omitempty"`
	// 模板id
	Tpl_id int `json:"tpl_id,omitempty"`
	// 模板参数
	Params []string `json:"params,omitempty"`
}

func newTmplSmsField(sign string, id int, params []string) *tmplSmsField {
	return &tmplSmsField{
		Sign:   sign,
		Tpl_id: id,
		Params: params,
	}
}
