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
// 错误对象
func ConverToStandardFormat(str string) (result time.Time, err error) {
	newStr := strings.Replace(str, "T", ":", -1)
	newStr = strings.Replace(newStr, "-", ":", -1)
	newStr = strings.Replace(newStr, ".", ":", -1)

	slice := strings.Split(newStr, ":")
	slice = slice[:6] // 只取前6位（表示年-月-日 时:分:秒）

	intSlice := make([]int, len(slice))
	for index, item := range slice {
		if intItem, err1 := strconv.Atoi(item); err1 != nil {
			err = fmt.Errorf("输入字符串的格式错误:%s,转换后的格式为:%s", str, newStr)
			return
		} else {
			intSlice[index] = intItem
		}
	}

	result = time.Date(intSlice[0], time.Month(intSlice[1]), intSlice[2], intSlice[3], intSlice[4], intSlice[5], 0, time.Local)
	return
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

// 计算两个时间的日期差值
func SubDay(time1, time2 time.Time) int {
	// 当前时间距离00:00:00的秒数
	awayFromZero := func(val time.Time) int64 {
		hour := val.Hour()
		minute := val.Minute()
		second := val.Second()
		return int64(hour*3600 + minute*60 + second)
	}

	// 每天对应的秒数
	var eachDaySecond int64 = 24 * 3600

	// 计算出两个时间对应的00:00:00时的时间戳
	unix1 := time1.Unix() - awayFromZero(time1)
	unix2 := time2.Unix() - awayFromZero(time2)

	if unix1 < unix2 {
		return int((unix2 - unix1) / eachDaySecond)
	} else {
		return int((unix1 - unix2) / eachDaySecond)
	}
}

// 解析时间字符串，要求时间格式形式为：12:59:59 这种形式
// timeStr:时间字符串
// 返回值:
// err:错误信息
// hour:小时值
// minute:分钟值
// second:秒数
func ParseTimeString(timeStr string) (err error, hour int, minute int, second int) {
	timeSlice := strings.Split(timeStr, ":")
	if len(timeSlice) != 3 {
		err = fmt.Errorf("时间字符串格式不正确：%v", timeStr)
		return
	}

	hour, _ = strconv.Atoi(timeSlice[0])
	minute, _ = strconv.Atoi(timeSlice[1])
	second, _ = strconv.Atoi(timeSlice[2])

	return
}

// 获取时间的日期值
// timeVal:时间值
// 返回值:
// time.Time:日期值
func GetDate(timeVal time.Time) time.Time {
	year, month, day := timeVal.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, timeVal.Location())
}
