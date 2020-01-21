package httpSender

import (
	"errors"
	"fmt"

	// "github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/webUtil"
)

// 实现Sender相关接口

type httpClient struct{}

// 发送Requester
func (this *httpClient) Send(req Requester) (rspn []byte, err error) {
	if req.GetMethod() == "POST" {
		rspn, err = this.post(req)
		return
	} else {
		err = errors.New(fmt.Sprintf("request: unsupport method (%s)", req.GetMethod()))
		return
	}
}

// 发送 post 请求
func (*httpClient) post(req Requester) ([]byte, error) {
	url := req.GetUrl()
	bytes := req.GetData()

	// debugUtil.Printf("httpClient-POST %s\r\n%v\n", url, string(bytes))

	rspn, err := webUtil.PostByteData(url, bytes, nil)
	if err != nil {
		return nil, err
	}

	// body := []byte("{\"result\":0,\"errmsg\":\"OK\",\"ext\":\"hello-world\",\"sid\":\"6:59106203271444582828\",\"fee\":1}")

	// debugUtil.Printf("httpClient-Res: %s", string(rspn))

	return rspn, nil
}

func New() *httpClient {
	return new(httpClient)
}
