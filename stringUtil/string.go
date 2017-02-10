package stringUtil

import (
	"runtime"
	"strconv"
	"strings"
	"unicode"
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

// 将字符串切割为[]int32
// str:输入字符串
// 返回值:
// []int
// error
func SplitToInt32Slice(s, sep string) ([]int32, error) {
	// 先获得int slice
	count := 0
	if intSlice, err := SplitToIntSlice(s, sep); err != nil {
		return nil, err
	} else {
		count = len(intSlice)
	}

	// 定义int32 slice
	int32Slice := make([]int32, 0, count)
	for _, item := range int32Slice {
		int32Slice = append(int32Slice, int32(item))
	}

	return int32Slice, nil
}

// 检查一个字符串是否是空字符串
// content:上下文字符串
// 返回值：
// bool:true：空字符串 false：非空字符串
func IsEmpty(content string) bool {
	if len(content) <= 0 {
		return true
	}

	return strings.IndexFunc(content, func(item rune) bool {
		return unicode.IsSpace(item) == false
	}) < 0
}
