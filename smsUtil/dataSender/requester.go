package dataSender

// 请求

type Requester interface {
    GetMethod() string
    GetUrl()    string
    GetData()   []byte

    // 解析数据
    // 返回:
    //      bool    是否发送成功
    //      error   错误
    ParseResponse([]byte) (bool, error)
}
