package netUtil

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// 获取远程Ip
func GetRemoteIp(conn net.Conn) string {
	items := strings.Split(conn.RemoteAddr().String(), ":")

	return items[0]
}

// 获取远程Ip和Port
func GetRemoteIpAndPort(conn net.Conn) (string, int) {
	items := strings.Split(conn.RemoteAddr().String(), ":")
	ip := items[0]
	port, _ := strconv.Atoi(items[1])

	return ip, port
}

// 获取远程地址（格式：Ip_Port）
func GetRemoteAddr(conn net.Conn) string {
	items := strings.Split(conn.RemoteAddr().String(), ":")

	return fmt.Sprintf("%s_%s", items[0], items[1])
}
