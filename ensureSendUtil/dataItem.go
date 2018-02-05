package ensureSendUtil

import (
	"github.com/Jordan/goutil/zlibUtil"
)

type dataItem interface {
	// 返回原始数据
	String() string

	// 返回发送字节流
	Bytes() []byte

	// 设置发送次数
	SetCount(uint)

	// 返回发送次数
	Count() uint
}

/////////////////////////////////////////////////
// httpDataItem

type httpDataItem struct {
	// 数据
	data string

	// 发送次数
	count uint
}

func newHTTPData(_data string) dataItem {
	return &httpDataItem{
		data:  _data,
		count: 0,
	}
}

// 返回原始数据
func (this *httpDataItem) String() string {
	return this.data
}

// 返回原始数据用于发送
func (this *httpDataItem) Bytes() []byte {
	return []byte(this.data)
}

func (this *httpDataItem) SetCount(cnt uint) {
	this.count = cnt
}

func (this *httpDataItem) Count() uint {
	return this.count
}

/////////////////////////////////////////////////
// tcpDataItem

type tcpDataItem struct {
	// 原始数据
	origin string

	// 压缩后数据
	data []byte

	// 重试次数
	count uint
}

func newTCPDataItem(_data string) (dataItem, error) {
	compressed, err := zlibUtil.Compress([]byte(_data), 5)
	if err != nil {
		return nil, err
	}

	item := &tcpDataItem{
		origin: _data,
		data:   compressed,
		count:  0,
	}
	return item, nil
}

// 返回原始数据
func (this *tcpDataItem) String() string {
	return this.origin
}

// 返回压缩数据用于发送
func (this *tcpDataItem) Bytes() []byte {
	return this.data
}

func (this *tcpDataItem) SetCount(cnt uint) {
	this.count = cnt
}

func (this *tcpDataItem) Count() uint {
	return this.count
}
