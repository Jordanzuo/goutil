package ensureSendUtil

import (
	"fmt"
)

/*
实现sender接口
*/

type baseSender struct {
	// 待发送的数据channel
	waitingDataChan chan dataItem

	// 失败数据缓存
	cachedDataChan chan dataItem

	// 用于停止协程
	done chan struct{}
}

func newBaseSender() *baseSender {
	return &baseSender{
		waitingDataChan: make(chan dataItem, 1024),
		cachedDataChan:  make(chan dataItem, 1024000),
		done:            make(chan struct{}),
	}
}

// Sender接口
// Send:
func (this *baseSender) Send() error {
	// baseSender不实现发送
	// 由tcpSender和httpSender实现发送
	return fmt.Errorf("baseSender dose not have Send Method")
}

// Sender接口
// Data: 返回待发送的数据channel
func (this *baseSender) Data() <-chan dataItem {
	return this.waitingDataChan
}

// Sender接口
// Cache：返回失败数据缓存channel
func (this *baseSender) Cache() chan dataItem {
	return this.cachedDataChan
}

// Sender接口
// Done：返回channel用于判断是否关闭
func (this *baseSender) Done() <-chan struct{} {
	return this.done
}
