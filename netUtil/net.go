package netUtil

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// 远程地址解析器
type RemoteAddrParser struct {
	// 主机地址（IP）
	Host string

	// 端口
	Port int
}

// 解析远程地址
func (this *RemoteAddrParser) parseRemoteAddr(remoteAddr string) {
	/*
		http中调用JoinHostPort来给RemoteAddr赋值；它的规则如下：
		JoinHostPort combines host and port into a network address of the
		form "host:port" or, if host contains a colon or a percent sign,
		"[host]:port".
		net包中是类似的

		所以现在要将RemoteAddr解析成host和port，则需要找到最后一个:，前面的部分则是host；
		如果host包含[]，则需要去除
	*/

	// 找到分隔host、port的:
	index := strings.LastIndex(remoteAddr, ":")
	if index == -1 {
		return
	}

	// 取出host部分
	this.Host = remoteAddr[:index]
	this.Port, _ = strconv.Atoi(remoteAddr[index+1:])

	// 处理host中可能的[]
	if strings.Index(this.Host, "[") == -1 {
		return
	}
	this.Host = this.Host[1:]

	if strings.Index(this.Host, "]") == -1 {
		return
	}
	this.Host = this.Host[:len(this.Host)-1]

	return
}

// 直接读取IP地址
func GetHttpAddr(request *http.Request) *RemoteAddrParser {
	this := &RemoteAddrParser{}
	this.parseRemoteAddr(request.RemoteAddr)
	return this
}

// 优先获取header中的代理地址，如果未设置代理地址，则使用request地址
func GetHttpAddr2(request *http.Request) *RemoteAddrParser {
	remoteAddr := request.Header.Get("HTTP_X_FORWARDED_FOR")
	if len(remoteAddr) > 0 {
		if len(remoteAddr) >= 10 {
			strArray := strings.Split(remoteAddr, ",")
			if len(strArray) > 0 {
				remoteAddr = strArray[0]
			}
		}

		if len(remoteAddr) > 0 && len(remoteAddr) <= 15 {
			this := &RemoteAddrParser{
				Host: remoteAddr,
			}
			return this
		}
	}

	remoteAddr = request.Header.Get("X-Real-IP")
	if len(remoteAddr) > 0 {
		this := &RemoteAddrParser{
			Host: remoteAddr,
		}
		return this
	}

	remoteAddr = request.Header.Get("X-Forwarded-For")
	if len(remoteAddr) > 0 {
		this := &RemoteAddrParser{
			Host: remoteAddr,
		}
		return this
	}

	this := &RemoteAddrParser{}
	this.parseRemoteAddr(request.RemoteAddr)
	return this
}

func GetWebSocketAddr(conn *websocket.Conn) *RemoteAddrParser {
	this := &RemoteAddrParser{}
	this.parseRemoteAddr(conn.RemoteAddr().String())
	return this
}

func GetConnAddr(conn net.Conn) *RemoteAddrParser {
	this := &RemoteAddrParser{}
	this.parseRemoteAddr(conn.RemoteAddr().String())
	return this
}
