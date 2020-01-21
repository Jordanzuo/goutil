package webUtil

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sort"
	"time"

	"github.com/Jordanzuo/goutil/netUtil"
)

// 获取请求的IP(obsolete, 直接调用netUtil.GetHttpAddr)
// 返回值：
// 请求的IP
func GetRequestIP(r *http.Request) string {
	/*
		http中调用JoinHostPort来给RemoteAddr赋值；它的规则如下：
		JoinHostPort combines host and port into a network address of the
		form "host:port" or, if host contains a colon or a percent sign,
		"[host]:port".

		所以现在要将RemoteAddr解析成host和port，则需要找到最后一个:，前面的部分则是host；
		如果host包含[]，则需要去除
	*/

	return netUtil.GetHttpAddr(r).Host
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

// 组装请求字符串
// requestMap:请求数据集合
func AssembleRequestParam(requestMap map[string]string) (dataStr string) {
	count := 0
	for k, v := range requestMap {
		if count == len(requestMap)-1 {
			dataStr += fmt.Sprintf("%s=%s", k, v)
		} else {
			dataStr += fmt.Sprintf("%s=%s&", k, v)
		}
		count += 1
	}

	return
}

// 按照参数名称升序组装请求字符串
// requestMap:请求数据集合
// asc:是否升序
func AssembleRequestParamSort(requestMap map[string]string, asc bool) (dataStr string) {
	finalKeys := make([]string, 0, len(requestMap))

	// 按key进行排序
	sortKeys := make([]string, 0, len(requestMap))
	for k := range requestMap {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys)

	// 排序是升序还是降序
	if asc {
		finalKeys = sortKeys
	} else {
		for i := len(sortKeys) - 1; i >= 0; i-- {
			finalKeys = append(finalKeys, sortKeys[i])
		}
	}

	for index, key := range finalKeys {
		if index == len(finalKeys)-1 {
			dataStr += fmt.Sprintf("%s=%s", key, requestMap[key])
		} else {
			dataStr += fmt.Sprintf("%s=%s&", key, requestMap[key])
		}
	}

	return
}
