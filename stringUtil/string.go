package stringUtil

import (
	"regexp"
	"runtime"
	"sort"
	"strings"
	"unicode"
)

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

// 检查是否存在特殊符号
// 1. emoji字符
// 2. ascii控制字符
// 3. \ " '
// val:待检查的字符串
// 返回值:
// bool:true:有特殊字符 false:无特殊字符
func IfHaveSpecialChar(val string) bool {
	if len(val) <= 0 {
		return false
	}

	// 表情符号过滤
	// Wide UCS-4 build
	emojiReg, _ := regexp.Compile("[^\U00000000-\U0000FFFF]+")
	if emojiReg.Match([]byte(val)) {
		return true
	}

	// 排除控制字符和特殊字符
	for _, charItem := range val {
		// 排除控制字符
		if (charItem > 0 && charItem < 0x20) || charItem == 0x7F {
			return true
		}

		// 排除部分特殊字符：  \ " '
		switch charItem {
		case '\\':
			fallthrough
		case '"':
			fallthrough
		case '\'':
			return true
		}
	}

	return false
}

// 判断string数组是否内容唯一
func IsDistinct_string(list []string) (result bool) {
	if len(list) == 0 || len(list) == 1 {
		result = true
		return
	}

	sort.Strings(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			result = false
			return
		}
	}

	result = true
	return
}
