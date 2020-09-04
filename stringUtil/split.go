package stringUtil

import (
	"fmt"
	"strconv"
	"strings"

	"public.com/goutil/mathUtil"
)

// 使用多分隔符来进行分割(默认分隔符为：",", ";", ":", "|", "||")
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

// 将字符串切割为[]float64
// s:输入字符串
// 返回值:
// []float64
// error
func SplitToFloat64Slice(s, sep string) ([]float64, error) {
	// 先按照分隔符进行切割
	strSlice := strings.Split(s, sep)

	// 定义float64 slice
	floatSlice := make([]float64, 0, len(strSlice))
	for _, value := range strSlice {
		// 去除空格
		if value = strings.TrimSpace(value); value == "" {
			continue
		}

		if value_float, err := strconv.ParseFloat(value, 64); err != nil {
			return nil, err
		} else {
			floatSlice = append(floatSlice, value_float)
		}
	}

	return floatSlice, nil
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
