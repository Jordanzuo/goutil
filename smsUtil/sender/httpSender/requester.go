package httpSender

// 请求
type Requester interface {
	GetMethod() string
	GetUrl() string
	GetData() []byte
}
