package qcloud

import "encoding/json"

// 实现sms.Sms各个接口
//
func (*qcloudsms) GetMethod() string {
    return "POST"
}

func (this *qcloudsms) GetUrl() string {
    return this.url
}

func (this *qcloudsms) GetData() []byte {
    bytes, err := json.Marshal(this.data)
    if err != nil {
        panic(err)
    }
    return bytes
}

func (this *qcloudsms) ParseResponse(rspn []byte) (bool, error) {
    this.rspn = new(QCloudResponse)
    err := json.Unmarshal(rspn, this.rspn)
    if err != nil {
        return false, err
    }
    return this.rspn.Result == 0, nil
}

func (this *qcloudsms) GetResponse() interface{} {
    return this.rspn
}
