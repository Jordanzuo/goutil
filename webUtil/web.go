package webUtil

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"
)

// 获取请求的IP
// 返回值：
// 请求的IP
func GetRequestIP(r *http.Request) string {
	ipAndPort := strings.Split(r.RemoteAddr, ":")
	if len(ipAndPort) == 2 {
		return ipAndPort[0]
	} else {
		return ""
	}
}

func NewTransport() *http.Transport {
	return &http.Transport{}
}

// timeout 超时时间
func GetTimeoutTransport(transport *http.Transport, timeout int) *http.Transport {
	transport.Dial = func(netw, addr string) (net.Conn, error) {
		deadline := time.Now().Add(time.Duration(timeout) * time.Second)
		c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
		if err != nil {
			return nil, err
		}
		c.SetDeadline(deadline)
		return c, nil
	}

	return transport
}

// b 表示是否需要验证http证书
func GetTLSTransport(transport *http.Transport, b bool) *http.Transport {
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: b}

	return transport
}
