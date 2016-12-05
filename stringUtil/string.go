package stringUtil

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Jordanzuo/goutil/securityUtil"
	"io"
	"runtime"
	"strconv"
	"strings"
)

// 截取字符串
// start：开始位置
// length：截取长度
// 返回值：
// 截取后的字符串
func Substring(str string, start, length int) string {
	// 先将字符串转化为[]rune格式（由于rune是字符串的基本单位）
	runeString := []rune(str)
	runeLength := len(runeString)
	end := 0

	// 计算起始位置
	if start > runeLength {
		start = runeLength
	}

	// 计算终止位置
	end = start + length
	if end > runeLength {
		end = runeLength
	}

	if start > end {
		start, end = end, start
	}

	return string(runeString[start:end])
}

// 根据不同平台获取换行符
// 返回值：
// 换行符
func GetNewLineString() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

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

// 将字符串切割为[]int
// str:输入字符串
// 返回值:
// []int
// error
func SplitToIntSlice(s, sep string) ([]int, error) {
	// 先按照分隔符进行切割
	strSlice := strings.Split(s, sep)

	// 定义int slice
	intSlice := make([]int, 0, len(strSlice))
	for _, value := range strSlice {
		// 去除空格
		if value = strings.TrimSpace(value); value == "" {
			continue
		}

		if value_int, err := strconv.Atoi(value); err != nil {
			return nil, err
		} else {
			intSlice = append(intSlice, value_int)
		}
	}

	return intSlice, nil
}
