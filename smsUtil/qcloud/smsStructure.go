package qcloud

// 电话结构
type telField struct {
    Nationcode  string `json:"nationcode"`
    Mobile      string `json:"mobile"`
}

// 公有字段
type commonField struct {
    // 签名
    Sig     string      `json:"sig"`
    // Unix时间戳
    Time    int64       `json:"time"`
    // 通道扩展码，可选字段，默认没有开通(需要填空)。
    Extend  string      `json:"extend"`
    // 用户的session内容，腾讯server回包中会原样返回
    Ext     string      `json:"ext"`

    // 接收方
    // 单发短信时填mobile
    // 群发时填[]mobile
    Tel     interface{} `json:"tel"`
}

// 直接发送短信字段
type msgSmsField struct {
    //0:普通短信;1:营销短信（强调：要按需填值，不然会影响到业务的正常使用）
    Type    interface{}     `json:"type,omitempty"`
    // 短信内容，需要匹配审核通过的模板内容
    Msg     string          `json:"msg,omitempty"`
}

// 发送模板短信字段
type tmplSmsField struct {
    // 签名 (前缀)
    Sign    string      `json:"sign,omitempty"`
    // 模板id
    Tpl_id  int         `json:"tpl_id,omitempty"`
    // 模板参数
    Params  []string    `json:"params,omitempty"`
}

// 发送短信请求结构
type smsData struct {
    // 公共字段
    *commonField

    // 短信字段
    // 根据类型，对其中一个结构的字段进行赋值
    *msgSmsField
    *tmplSmsField
}

// qcloud sms
type qcloudsms struct {
    url   string
    data  *smsData
    rspn  *QCloudResponse

    rnd   int
    appid string
}
