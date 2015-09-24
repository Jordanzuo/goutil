package stringUtil

import (
	"runtime"
)

// 截取字符串
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

// 获取换行符
func GetNewLineString() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}