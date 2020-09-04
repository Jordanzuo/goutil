package stringUtil

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	"github.com/Jordanzuo/goutil/securityUtil"
)

// 获取新的GUID字符串
// 返回值：
// 新的GUID字符串
func GetNewGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return securityUtil.Md5String(base64.URLEncoding.EncodeToString(b), true)
}

// 生成空的GUID字符串
// 返回值：
// 空的GUID字符串
func GetEmptyGUID() string {
	return "00000000-0000-0000-0000-000000000000"
}

// 判断GUID是否为空
// guid：GUID
// 返回值：
// 是否为空
func IsGUIDEmpty(guid string) bool {
	if guid == "" || guid == "00000000-0000-0000-0000-000000000000" {
		return true
	}

	return false
}

// 获取新的GUID字符串
// 返回值：
// 新的GUID字符串
func GetNewUUID() string {
	str := GetNewGUID()
	var builder strings.Builder
	builder.WriteString(Substring(str, 0, 8))
	builder.WriteString("-")
	builder.WriteString(Substring(str, 8, 4))
	builder.WriteString("-")
	builder.WriteString(Substring(str, 12, 4))
	builder.WriteString("-")
	builder.WriteString(Substring(str, 16, 4))
	builder.WriteString("-")
	builder.WriteString(Substring(str, 20, 12))

	return strings.ToLower(builder.String())
}
