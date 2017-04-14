package dataSender

// 发送器
// 用于发送Requester

type Sender interface {
    Send(req Requester) ([]byte, error)
}
