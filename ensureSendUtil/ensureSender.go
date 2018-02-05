package ensureSendUtil

type EnsureSender interface {
	// use Write to send data
	Write(string) error

	// stop sender
	Close() error
}

// resend和dataSaver通过此接口调用tcpSender与httpSender
type sender interface {
	// 发送数据
	Send(dataItem) error

	// 返回待发送的数据channel
	Data() <-chan dataItem

	// 返回失败数据缓存channel
	Cache() chan dataItem

	// 用于判断是否关闭
	Done() <-chan struct{}
}
