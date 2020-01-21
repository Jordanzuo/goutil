package sms

type Sms interface {
	// 发送
	Send() (bool, error)

	// 用于获取准确的返回数据
	GetResponse() interface{}
}
