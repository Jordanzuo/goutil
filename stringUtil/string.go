package stringUtil

import (
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/Jordanzuo/goutil/mathUtil"
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

// 使用多分隔符来进行分割
// eg:1,2;3|4||5,6;7|8||9
// 返回值:
// []string
func Split(s string, seps []string) []string {
	retList := make([]string, 0, 32)

	// 如果seps为nil，则使用默认值
	if seps == nil {
		seps = []string{",", ";", ":", "|", "||"}
	}

	// 根据所有的分隔符来一点一点地切割字符串，直到不可切割为止
	for {
		startIndex := len(s) - 1
		endIndex := 0
		exists := false

		// 遍历，找到第一个分割的位置
		for _, sep := range seps {
			index := strings.Index(s, sep)

			// 如果找到有匹配项，则寻找最小的pos，如果有多个相同的pos，则使用长度最长的分隔符
			if index > -1 {
				exists = true

				// 说明有多个有效的分隔符，如|和||
				if index < startIndex {
					startIndex = index
					endIndex = startIndex + len(sep) - 1
				} else if index == startIndex {
					if startIndex+len(sep)-1 > endIndex {
						endIndex = startIndex + len(sep) - 1
					}
				}
			}
		}

		// 如果没有找到匹配的pos，则分割过程结束
		if !exists {
			retList = append(retList, s)
			break
		}

		// 切割字符串
		sub := s[:startIndex]
		if sub != "" {
			retList = append(retList, sub)
		}
		s = s[endIndex+1:]
	}

	return retList
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
// s:输入字符串
// 返回值:
// []int
// error
func SplitToInt32Slice(s, sep string) ([]int32, error) {
	// 先获得int slice
	count := 0
	intSlice, err := SplitToIntSlice(s, sep)
	if err != nil {
		return nil, err
	} else {
		count = len(intSlice)
	}

	// 定义int32 slice
	int32Slice := make([]int32, 0, count)
	for _, item := range intSlice {
		int32Slice = append(int32Slice, int32(item))
	}

	return int32Slice, nil
}

// 将字符串切割为IntRegion列表
// s:输入字符串，形如：1-200,201-400,401-1000
// outerSep:外部分隔符
// innerSep:内部分隔符
// 返回值：
// IntRegion列表
// 错误对象
func SplitToIntRegion(s, outerSep, innerSep string) (intRegionList []*mathUtil.IntRegion, err error) {
	if s == "" {
		err = fmt.Errorf("Input is empty")
		return
	}

	outerRegionList := make([]string, 0, 4)
	outerRegionList = strings.Split(s, outerSep)
	if len(outerRegionList) == 0 {
		err = fmt.Errorf("%s:Format invalid. Such as:1-100,101-200", s)
		return
	}

	for _, item := range outerRegionList {
		innerRegionList := make([]string, 0, 2)
		innerRegionList = strings.Split(item, innerSep)
		if len(innerRegionList) != 2 {
			err = fmt.Errorf("%s:Format invalid. Such as:1-100", item)
			return
		}

		var lower, upper int
		lower, err = strconv.Atoi(innerRegionList[0])
		if err != nil {
			return
		}
		upper, err = strconv.Atoi(innerRegionList[1])
		if err != nil {
			return
		}
		if lower > upper {
			err = fmt.Errorf("lower:%d should less than upper:%d", lower, upper)
			return
		}

		intRegionList = append(intRegionList, mathUtil.NewIntRegion(lower, upper))
	}

	return
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
