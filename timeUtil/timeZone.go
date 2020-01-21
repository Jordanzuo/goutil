package timeUtil

import (
	"time"
)

// 把时区转换到UTC时区，但时间值不变(去掉时区的影响)
func GetUTCTime(timeVal time.Time) time.Time {
	_, offset := timeVal.Zone()
	timeVal = timeVal.Add(time.Duration(offset) * time.Second)
	timeVal = timeVal.In(time.UTC)

	return timeVal
}

// 把时间转换成本地时区，但时间值不变(去掉时区的影响)
func GetLocalTime(timeVal time.Time) time.Time {
	// 获取本地时区的时间偏移
	tmpVal := time.Now()
	_, localOffset := tmpVal.Zone()

	// 获取指定时间值的时区偏移
	_, timeOffset := timeVal.Zone()
	timeVal = timeVal.Add(time.Duration(timeOffset-localOffset) * time.Second)
	timeVal = timeVal.In(time.Local)

	return timeVal
}

// 增加本地时区的值到指定时间上
func AddLocalTimeZone(timeVal int64) int64 {
	// 获取本地时区的时间偏移
	tmpVal := time.Now()
	_, localOffset := tmpVal.Zone()

	return timeVal + int64(localOffset)
}

// 增加本地时区的值到指定时间上
func AddLocalTimeZone2(timeVal time.Time) time.Time {
	// 获取本地时区的时间偏移
	tmpVal := time.Now()
	_, localOffset := tmpVal.Zone()

	return timeVal.Add(time.Duration(localOffset) * time.Second)
}

// 减去本地时区到指定时间上
func SubLocalTimeZone(timeVal int64) int64 {
	// 获取本地时区的时间偏移
	tmpVal := time.Now()
	_, localOffset := tmpVal.Zone()

	return timeVal + -1*int64(localOffset)
}

// 减去本地时区到指定时间上
func SubLocalTimeZone2(timeVal time.Time) time.Time {
	// 获取本地时区的时间偏移
	tmpVal := time.Now()
	_, localOffset := tmpVal.Zone()

	return timeVal.Add(-1 * time.Duration(localOffset) * time.Second)
}
