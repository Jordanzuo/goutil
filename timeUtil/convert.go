package timeUtil

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 将字符串转换为标准的时间格式
// str：输入，格式为：2015-10-25T17:07:30
// 返回值：
// 标准时间格式对象
func ConverToStandardFormat(str string) time.Time {
	str = strings.Replace(str, "T", ":", -1)
	str = strings.Replace(str, "-", ":", -1)
	slice := strings.Split(str, ":")
	intSlice := make([]int, len(slice))
	for index, item := range slice {
		if intItem, err := strconv.Atoi(item); err != nil {
			panic(errors.New(fmt.Sprintf("输入字符串的格式错误:%s", str)))
		} else {
			intSlice[index] = intItem
		}
	}

	// 检查数量
	if len(intSlice) != 6 {
		panic(errors.New(fmt.Sprintf("输入字符串的格式错误:%s", str)))
	}

	return time.Date(intSlice[0], time.Month(intSlice[1]), intSlice[2], intSlice[3], intSlice[4], intSlice[5], 0, time.Local)
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
