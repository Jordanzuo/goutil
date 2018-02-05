package ensureSendUtil

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/Jordan/goutil/debugUtil"
	"github.com/Jordan/goutil/zlibUtil"
)

// 保存接收的数据用于校验
var tcp_recv_msg = make([]byte, 0)

func init() {
	debugUtil.SetDebug(true)
}

// 创建socket服务器，保存收到的数据
func server(addr string) net.Listener {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			for {
				buff := make([]byte, 512)
				_, err := conn.Read(buff)
				if err != nil {
					break
				} else {
					decompressed, err := zlibUtil.Decompress(buff[4:])
					if err != nil {
						panic(err)
					} else {
						tcp_recv_msg = append(tcp_recv_msg, decompressed...)
					}
				}
			}
		}
	}()

	return listener
}

func Test_tcp(t *testing.T) {
	// 开启服务器
	l := server("127.0.0.1:9559")

	tcp, err := NewTCPSender("./test_tcp", "127.0.0.1:9559")
	if err != nil {
		t.Error(err)
	}

	// 发送消息
	tcp.Write("tcp-msg-1")
	time.Sleep(time.Millisecond * 50) // 等待协程发送数据

	// 关闭连接和服务器
	l.Close()
	(tcp.(*tcpSender)).conn.Close()

	// 发送消息，此数据会失败
	tcp.Write("tcp-msg-2")
	// time.Sleep(time.Millisecond * 50)

	// 保存数据
	tcp.Close()

	// 重启，检查是否重发tcp-msg-2
	l = server("127.0.0.1:9559")
	tcp, err = NewTCPSender("./test_tcp", "127.0.0.1:9559")
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 2)

	if string(tcp_recv_msg) != "tcp-msg-1tcp-msg-2" {
		t.Error("message error. got " + string(tcp_recv_msg))
	} else {
		fmt.Println("TCP OK")
	}

	tcp.Close()
	l.Close()
}
