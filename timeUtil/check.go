package timeUtil

import (
	"time"
)

// 检查一个日期的时分秒 是否在指定的时间范围内
// checkTime:待检查的时间
// timeSpan1:时间范围1
// addSecond1:需要加上的时间偏差值
// timeSpan2:时间范围2
// addSecond2:需要加上的时间偏差值
func CheckIfInRange(checkTime time.Time, timeSpan1 string, addSecond1 int, timeSpan2 string, addSecond2 int) bool {
	var (
		hour1, minute1, second1 int
		hour2, minute2, second2 int
	)

	// 取出字符串的时分秒
	_, hour1, minute1, second1 = ParseTimeString(timeSpan1)
	_, hour2, minute2, second2 = ParseTimeString(timeSpan2)

	// 取出当前时间的时分秒
	checkTime = checkTime.Local()
	hour := checkTime.Hour()
	minute := checkTime.Minute()
	second := checkTime.Second()

	// 转换成时间值
	val := int64(time.Hour)*int64(hour) + int64(time.Minute)*int64(minute) + int64(time.Second)*int64(second)
	val1 := int64(time.Hour)*int64(hour1) + int64(time.Minute)*int64(minute1) + int64(time.Second)*int64((second1+addSecond1))
	val2 := int64(time.Hour)*int64(hour2) + int64(time.Minute)*int64(minute2) + int64(time.Second)*int64((second2+addSecond2))

	if val1 <= val && val <= val2 {
		return true
	}

	if val2 <= val && val <= val1 {
		return true
	}

	return false
}

// 检查一个日期的时分秒 是否在指定的时间范围内
// checkTime:待检查的时间
// timeSpan1:时间范围1
// timeSpan2:时间范围2
func CheckIfInRange2(checkTime time.Time, timeSpan1 string, timeSpan2 string) bool {
	return CheckIfInRange(checkTime, timeSpan1, 0, timeSpan2, 0)
}

// 检查两个日期是否在同一天
// time1:时间1
// time2:时间2
// 返回值:
// bool:true：在同一天 false:不在同一天
func CheckIfInSameDate(time1, time2 time.Time) bool {
	y1, m1, d1 := time1.Date()
	y2, m2, d2 := time2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}
