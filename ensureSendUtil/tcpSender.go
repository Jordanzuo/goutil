package ensureSendUtil

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Jordan/Framework/goroutineMgr"
	"github.com/Jordan/goutil/intAndBytesUtil"
)

var (
	errConnectEmpty = fmt.Errorf("scoket reconnecting...")
	byterOrder      = binary.LittleEndian
)

// 实现 EnsureSender和sender接口
type tcpSender struct {
	// 需要实现的接口
	EnsureSender

	// 包含sender接口部分实现
	*baseSender

	// 数据目录
	dataFolder string

	// 服务器地址
	address string

	// 连接
	conn net.Conn

	// 用于重连时互斥
	mutex sync.Mutex

	// 用于sendLoop和resendLoop发送退出信号
	closeSignal chan struct{}
}

// 创建一个tcp数据发送器
// 参数：
// 		_dataFolder  数据存放目录
// 		_address     连接地址
func NewTCPSender(_dataFolder, _address string) (EnsureSender, error) {
	// 连接服务器
	conn, err := net.DialTimeout("tcp", _address, 5*time.Second)
	if err != nil {
		return nil, err
	}

	this := &tcpSender{
		dataFolder:  _dataFolder,
		baseSender:  newBaseSender(),
		address:     _address,
		conn:        conn,
		closeSignal: make(chan struct{}),
	}

	// 新开协程发送数据
	go sendLoop(this, this.closeSignal)

	// 定时重发
	go resendLoop(this, _dataFolder, this.closeSignal)

	// 发送心跳包
	go this.heartBeat()

	return this, nil
}

// 每隔15秒发送心跳包
func (this *tcpSender) heartBeat() {
	name := "ensureSendUtil.tcpSender.heartBeat"
	goroutineMgr.MonitorZero(name)
	defer goroutineMgr.ReleaseMonitor(name)

	tick := time.Tick(time.Second * 15)

	for {
		select {
		case <-this.Done():
			return
		case <-tick:
			this.sendBytes([]byte{})
		}
	}
}

// EnsureSender接口
// Write：写入数据
func (this *tcpSender) Write(data string) error {
	item, err := newTCPDataItem(data)
	if err != nil {
		return err
	}

	this.waitingDataChan <- item

	return nil
}

// EnsureSender接口
// Close：关闭
func (this *tcpSender) Close() error {
	close(this.done)

	// 关闭socket连接
	conn := this.conn
	if conn != nil {
		conn.Close()
	}

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

// Sender接口
// Send：发送dataItem
func (this *tcpSender) Send(item dataItem) error {
	err := this.sendBytes(item.Bytes())
	if err != nil && err != errConnectEmpty {
		// 发送失败时发送次数+1
		item.SetCount(item.Count() + 1)
	}

	return err
}

// 发送字节数据
// 发送格式：[lenght+data]
func (this *tcpSender) sendBytes(data []byte) error {
	conn := this.conn
	if conn == nil {
		return errConnectEmpty
	}

	// 将长度转化为字节数组
	header := intAndBytesUtil.Int32ToBytes(int32(len(data)), byterOrder)

	if len(data) > 0 {
		data = append(header, data...)
	} else {
		data = header
	}

	_, err := conn.Write(data)
	if err != nil {
		this.mutex.Lock()
		// 发送失败
		// 检查失败的conn是否this.conn（避免多个线程失败后均调用reconnect）
		// 是则关闭并重连
		if conn == this.conn {
			this.conn.Close()
			this.conn = nil
			this.mutex.Unlock()

			// 重连
			go this.reconnect()
		} else {
			this.mutex.Unlock()
		}
	}

	return err
}

// 重连服务器
func (this *tcpSender) reconnect() error {
	// lock-it
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for {
		// 检查是否已经重连
		if this.conn != nil {
			return nil
		}

		conn, err := net.DialTimeout("tcp", this.address, 5*time.Second)
		if err != nil {
			// 连接失败，5秒后重试
			<-time.After(time.Second * 5)
			continue
		}

		this.conn = conn
	}
}
