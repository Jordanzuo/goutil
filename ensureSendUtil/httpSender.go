package ensureSendUtil

import (
	"fmt"

	"github.com/Jordanzuo/goutil/webUtil"
)

// 实现 EnsureSender和sender接口
type httpSender struct {
	// 需要实现的接口
	EnsureSender

	// 包含sender接口部分实现
	*baseSender

	// 数据目录
	dataFolder string

	// 发送地址
	url string

	// 用于sendLoop和resendLoop发送退出信号
	closeSignal chan struct{}
}

// 创建一个http数据发送器
// 参数：
//      _dataFolder  数据存放目录
//      _url         发送地址
func NewHTTPSender(_dataFolder, _url string) (EnsureSender, error) {
	this := &httpSender{
		dataFolder:  _dataFolder,
		url:         _url,
		baseSender:  newBaseSender(),
		closeSignal: make(chan struct{}),
	}

	// 新开协程发送数据
	go sendLoop(this, this.closeSignal)

	// 定时重发
	go resendLoop(this, _dataFolder, this.closeSignal)

	return this, nil
}

// EnsureSender接口
// Write：写入数据
func (this *httpSender) Write(data string) error {
	item := newHTTPData(data)

	this.waitingDataChan <- item

	return nil
}

// EnsureSender接口
// Close：关闭
func (this *httpSender) Close() error {
	close(this.done)

	// 等待sendLoop和resendLoop退出
	<-this.closeSignal
	<-this.closeSignal

	// 保存数据
	_, e1 := saveData(this.Cache(), this.dataFolder)
	_, e2 := saveData(this.Data(), this.dataFolder)

	if e2 != nil {
		if e1 != nil {
			return fmt.Errorf("%s %s", e1, e2)
		}
		return e2
	} else {
		return e1
	}
}

// sender接口
// Send：发送数据
func (this *httpSender) Send(item dataItem) error {
	statusCode, _, err := webUtil.PostByteData2(this.url, item.Bytes(), nil, nil)
	if err != nil || statusCode != 200 {
		if err == nil {
			err = fmt.Errorf("StatusCode is not 200")
		}

		// 发送失败时发送次数+1
		item.SetCount(item.Count() + 1)
	}

	return err
}
