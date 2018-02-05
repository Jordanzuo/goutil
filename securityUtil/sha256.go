package securityUtil

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// 对字符串进行SHA256加密，并且可以选择返回大、小写
// s:输入字符串
// ifUpper:输出是否大写
// 返回值：sha256加密后的十六进制字符串
func Sha256String(s string, ifUpper bool) string {
	if len(s) == 0 {
		panic(errors.New("input string can't be empty"))
	}

	buf := Sha256Bytes([]byte(s))

	if ifUpper {
		return fmt.Sprintf("%X", buf)
	} else {
		return fmt.Sprintf("%x", buf)
	}
}

// 对字符数组进行SHA256加密
// b:输入字符数组
// 返回值：sha256加密后的原数据
func Sha256Bytes(b []byte) []byte {
	if len(b) == 0 {
		panic(errors.New("input []byte can't be empty"))
	}

	sha1Instance := sha256.New()
	sha1Instance.Write(b)
	result := sha1Instance.Sum(nil)

	return result
}
