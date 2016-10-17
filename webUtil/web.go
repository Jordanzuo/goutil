package webUtil

import (
	"net/http"
	"strings"
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
