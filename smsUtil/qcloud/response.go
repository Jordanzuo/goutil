package qcloud

// 腾讯云平台的响应信息

// 响应信息
type QCloudResponse struct {
    Result  int     // 0表示成功(计费依据)，非0表示失败
    Errmsg  string  // result非0时的具体错误信息
    Ext     string  // 用户的session内容，腾讯server回包中会原样返回
    Sid     string  // 标识本次发送id，标识一次短信下发记录
    Fee     int     // 短信计费的条数
    Detail  []responsDetailItem  // 群发短信时才有
}

type responsDetailItem struct {
    Result      int
    Errmsg      string
    Mobile      string
    Nationcode  string
    Sid         string
    Fee         int
}
