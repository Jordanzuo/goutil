package stringUtil

import (
	"fmt"
	"strconv"
)

// 首字母小写
func FirstCharToLower(str string) string {
	if len(str) < 1 {
		return ""
	}

	runeArray := []rune(str)
	if runeArray[0] >= 65 && runeArray[0] <= 90 {
		runeArray[0] += 32
	}
	return string(runeArray)
}

// 首字母大写
func FirstCharToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}

	runeArray := []rune(str)
	if runeArray[0] >= 97 && runeArray[0] <= 122 {
		runeArray[0] -= 32
	}
	return string(runeArray)
}

// 将形如1,2|3,4|5,6的字符串转化成map
// 返回值:
// map[string]string
// 错误对象
func StringToMap_String_String(str string, seps []string) (data map[string]string, err error) {
	strList := Split(str, seps)
	if len(strList) == 0 {
		err = fmt.Errorf("str is empty.")
		return
	}

	if len(strList)%2 != 0 {
		err = fmt.Errorf("str has odd items.")
		return
	}

	data = make(map[string]string, len(strList)/2)
	for i := 0; i < len(strList); i += 2 {
		data[strList[i]] = strList[i+1]
	}

	return
}

// 将形如1,2|3,4|5,6的字符串转化成map
// 返回值:
// map[string]int
// 错误对象
func StringToMap_String_Int(str string, seps []string) (data map[string]int, err error) {
	strList := Split(str, seps)
	if len(strList) == 0 {
		err = fmt.Errorf("str is empty.")
		return
	}

	if len(strList)%2 != 0 {
		err = fmt.Errorf("str has odd items.")
		return
	}

	data = make(map[string]int, len(strList)/2)
	for i := 0; i < len(strList); i += 2 {
		key := strList[i]
		value, err1 := strconv.Atoi(strList[i+1])
		if err1 != nil {
			err = fmt.Errorf("Type convertion failed. Value:%s, Error:%v", strList[i+1], err1)
			return
		}
		data[key] = value
	}

	return
}

// 将形如1,2|3,4|5,6的字符串转化成map
// 返回值:
// map[int]int
// 错误对象
func StringToMap_Int_Int(str string, seps []string) (data map[int]int, err error) {
	strList := Split(str, seps)
	if len(strList) == 0 {
		err = fmt.Errorf("str is empty.")
		return
	}

	if len(strList)%2 != 0 {
		err = fmt.Errorf("str has odd items.")
		return
	}

	data = make(map[int]int, len(strList)/2)
	for i := 0; i < len(strList); i += 2 {
		key, err1 := strconv.Atoi(strList[i])
		if err1 != nil {
			err = fmt.Errorf("Type convertion failed. Value:%s, Error:%v", strList[i], err1)
			return
		}
		value, err2 := strconv.Atoi(strList[i+1])
		if err2 != nil {
			err = fmt.Errorf("Type convertion failed. Value:%s, Error:%v", strList[i+1], err2)
			return
		}
		data[key] = value
	}

	return
}

// 将形如1,2|3,4|5,6的字符串转化成map
// 返回值:
// map[int32]int32
// 错误对象
func StringToMap_Int32_Int32(str string, seps []string) (data map[int32]int32, err error) {
	strList := Split(str, seps)
	if len(strList) == 0 {
		err = fmt.Errorf("str is empty.")
		return
	}

	if len(strList)%2 != 0 {
		err = fmt.Errorf("str has odd items.")
		return
	}

	data = make(map[int32]int32, len(strList)/2)
	for i := 0; i < len(strList); i += 2 {
		key, err1 := strconv.Atoi(strList[i])
		if err1 != nil {
			err = fmt.Errorf("Type convertion failed. Value:%s, Error:%v", strList[i], err1)
			return
		}
		value, err2 := strconv.Atoi(strList[i+1])
		if err2 != nil {
			err = fmt.Errorf("Type convertion failed. Value:%s, Error:%v", strList[i+1], err2)
			return
		}
		data[int32(key)] = int32(value)
	}

	return
}
