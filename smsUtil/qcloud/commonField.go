package qcloud

// 公有字段
type commonField struct {
	// 签名
	Sig string `json:"sig"`
	// Unix时间戳
	Time int64 `json:"time"`
	// 通道扩展码，可选字段，默认没有开通(需要填空)。
	Extend string `json:"extend"`
	// 用户的session内容，腾讯server回包中会原样返回
	Ext string `json:"ext"`

	// 接收方
	// 单发短信时填mobile
	// 群发时填[]mobile
	Tel interface{} `json:"tel"`
}

func newCommonField(sig string, time int64, extend string, ext string, tel interface{}) *commonField {
	return &commonField{
		Sig:    sig,
		Time:   time,
		Extend: extend,
		Ext:    ext,
		Tel:    tel,
	}
}
