package timeUtil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 将字符串转换为标准的时间格式
// str：输入，格式为：2015-10-25T17:07:30
// 返回值：
// 标准时间格式对象
func ConverToStandardFormat(str string) (time.Time, error) {
	newStr := strings.Replace(str, "T", ":", -1)
	newStr = strings.Replace(newStr, "-", ":", -1)
	newStr = strings.Replace(newStr, ".", ":", -1)
	slice := strings.Split(newStr, ":")
	// 只取前6位（表示年-月-日 时:分:秒）
	slice = slice[:6]

	intSlice := make([]int, len(slice))
	for index, item := range slice {
		if intItem, err := strconv.Atoi(item); err != nil {
			return time.Now(), fmt.Errorf("输入字符串的格式错误:%s,转换后的格式为:%s", str, newStr)
		} else {
			intSlice[index] = intItem
		}
	}

	return time.Date(intSlice[0], time.Month(intSlice[1]), intSlice[2], intSlice[3], intSlice[4], intSlice[5], 0, time.Local), nil
}

// 将时间转换为int类型（20160120，共8位）
// t：时间
// 返回值：
// int类型的数字
func ConvertToInt(t time.Time) int {
	year := int(t.Year())
	month := int(t.Month())
	day := int(t.Day())

	return year*10e3 + month*10e1 + day
}
