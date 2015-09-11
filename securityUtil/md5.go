package securityUtil

import (
	"crypto/md5"
	"errors"
	"fmt"
)

// 对字符串进行MD5加密，并且可以选择返回大、小写
// s:输入字符串
// ifUpper:输出是否大写
// 返回值：md5加密后的字符串
func Md5String(s string, ifUpper bool) string {
	if len(s) == 0 {
		panic(errors.New("input string can't be empty"))
	}

	return Md5Bytes([]byte(s), ifUpper)
}

// 对字符数组进行MD5加密，并且可以选择返回大、小写
// b:输入字符数组
// ifUpper:输出是否大写
// 返回值：md5加密后的字符串
func Md5Bytes(b []byte, ifUpper bool) string {
	if len(b) == 0 {
		panic(errors.New("input []byte can't be empty"))
	}

	md5Instance := md5.New()
	md5Instance.Write(b)
	result := md5Instance.Sum([]byte(""))
	if ifUpper {
		return fmt.Sprintf("%X", result)
	} else {
		return fmt.Sprintf("%x", result)
	}
}
