package qcloud

type response struct {
	Result int    `json: "result"`
	ErrMsg string `json:"errmsg"`
	CallId string `json:"callid"`
	Ext    string `json:"ext"`
}
