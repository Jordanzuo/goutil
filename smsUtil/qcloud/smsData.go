package qcloud

// 发送短信请求结构
type smsData struct {
	// 公共字段
	*commonField

	// 短信字段
	// 根据类型，对其中一个结构的字段进行赋值
	*msgSmsField
	*tmplSmsField
}

func newSmsData(common *commonField, msg *msgSmsField, tmpl *tmplSmsField) *smsData {
	return &smsData{common, msg, tmpl}
}
