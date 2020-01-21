package qcloud

import (
	"encoding/json"
	"fmt"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/smsUtil/sender/httpSender"
)

// qcloud sms
type qcloudsms struct {
	url   string
	data  *smsData
	rnd   int
	appid string

	rspn *QCloudResponse
}

func newSms(url string, data *smsData, rnd int, appid string) *qcloudsms {
	return &qcloudsms{
		url:   url,
		data:  data,
		rnd:   rnd,
		appid: appid,
	}
}

// 实现sms.Sms、httpSender.Requester、httpSender.Responser接口

// httpSender.Requester接口
// 返回请求方式
func (*qcloudsms) GetMethod() string {
	return "POST"
}

// httpSender.Requester接口
// 返回请求url
func (this *qcloudsms) GetUrl() string {
	// url 额外参数
	url := fmt.Sprintf("%s?sdkappid=%s&random=%d", this.url, this.appid, this.rnd)

	return url
}

// httpSender.Requester接口
// 返回发送数据
func (this *qcloudsms) GetData() []byte {
	bytes, err := json.Marshal(this.data)
	if err != nil {
		debugUtil.Println("failed to json marshal sms data")
		return []byte("")
	}
	return bytes
}

// 解析返回数据，判断是否发送成功
func (this *qcloudsms) ParseReponse(rspn []byte) (bool, error) {
	this.rspn = new(QCloudResponse)
	err := json.Unmarshal(rspn, this.rspn)
	if err != nil {
		return false, err
	}

	if this.rspn.Result == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf(this.rspn.Errmsg)
	}
}

// sms.Sms接口
func (this *qcloudsms) Send() (bool, error) {
	sender := httpSender.New()
	rspn, err := sender.Send(this)
	if err != nil {
		return false, err
	}

	return this.ParseReponse(rspn)
}

func (this *qcloudsms) GetResponse() interface{} {
	return this.rspn
}
