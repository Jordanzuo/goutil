package timeUtil

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Jordanzuo/goutil/stringUtil"
)

// format time like java, such as: yyyy-MM-dd HH:mm:ss
// t：时间
// format：格式化字符串
// 返回值：
// 格式化后的字符串
func Format(t time.Time, format string) string {
	//year
	if strings.ContainsAny(format, "y") {
		year := strconv.Itoa(t.Year())

		if strings.Count(format, "yy") == 1 && strings.Count(format, "y") == 2 {
			format = strings.Replace(format, "yy", year[2:], 1)
		} else if strings.Count(format, "yyyy") == 1 && strings.Count(format, "y") == 4 {
			format = strings.Replace(format, "yyyy", year, 1)
		} else {
			panic("format year error! please 'yyyy' or 'yy'")
		}
	}

	//month
	if strings.ContainsAny(format, "M") {
		var month string

		if int(t.Month()) < 10 {
			month = "0" + strconv.Itoa(int(t.Month()))
		} else {
			month = strconv.Itoa(int(t.Month()))
		}

		if strings.Count(format, "MM") == 1 && strings.Count(format, "M") == 2 {
			format = strings.Replace(format, "MM", month, 1)
		} else {
			panic("format month error! please 'MM'")
		}
	}

	//day
	if strings.ContainsAny(format, "d") {
		var day string

		if t.Day() < 10 {
			day = "0" + strconv.Itoa(t.Day())
		} else {
			day = strconv.Itoa(t.Day())
		}

		if strings.Count(format, "dd") == 1 && strings.Count(format, "d") == 2 {
			format = strings.Replace(format, "dd", day, 1)
		} else {
			panic("format day error! please 'dd'")
		}
	}

	//hour
	if strings.ContainsAny(format, "H") {
		var hour string

		if t.Hour() < 10 {
			hour = "0" + strconv.Itoa(t.Hour())
		} else {
			hour = strconv.Itoa(t.Hour())
		}

		if strings.Count(format, "HH") == 1 && strings.Count(format, "H") == 2 {
			format = strings.Replace(format, "HH", hour, 1)
		} else {
			panic("format hour error! please 'HH'")
		}
	}

	//minute
	if strings.ContainsAny(format, "m") {
		var minute string

		if t.Minute() < 10 {
			minute = "0" + strconv.Itoa(t.Minute())
		} else {
			minute = strconv.Itoa(t.Minute())
		}
		if strings.Count(format, "mm") == 1 && strings.Count(format, "m") == 2 {
			format = strings.Replace(format, "mm", minute, 1)
		} else {
			panic("format minute error! please 'mm'")
		}
	}

	//second
	if strings.ContainsAny(format, "s") {
		var second string

		if t.Second() < 10 {
			second = "0" + strconv.Itoa(t.Second())
		} else {
			second = strconv.Itoa(t.Second())
		}

		if strings.Count(format, "ss") == 1 && strings.Count(format, "s") == 2 {
			format = strings.Replace(format, "ss", second, 1)
		} else {
			panic("format second error! please 'ss'")
		}
	}

	return format
}

// 转换成日期字符串
// timeVal：待转换的时间
// 返回值：
// string:格式形如：2016-10-10
/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday
*/
func ToDateString(timeVal time.Time) string {
	return timeVal.Local().Format("2006-01-02")
}

// 忽略时区，转换成日期字符串
// timeVal：待转换的时间
// 返回值：
// string:格式形如：2016-10-10
/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday
*/
func ToDateString2(timeVal time.Time) string {
	return timeVal.Format("2006-01-02")
}

// 以本地时区为准，转换成时间字符串
// timeVal：待转换的时间
// 返回值：
// string:格式形如：2016-10-10 10:10:10
/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday
*/
func ToDateTimeString(timeVal time.Time) string {
	return ToDateTimeStringEx(timeVal, false)
}

func ToDateTimeStringEx(timeVal time.Time, flagT bool) string {
	if flagT {
		val := timeVal.Local().Format("2006-01-02 15:04:05")
		return strings.Replace(val, " ", "T", -1)
	}

	return timeVal.Local().Format("2006-01-02 15:04:05")
}

// 忽略时区，转换成时间字符串
// timeVal：待转换的时间
// 返回值：
// string:格式形如：2016-10-10 10:10:10
/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday
*/
func ToDateTimeString2(timeVal time.Time) string {
	return ToDateTimeStringEx2(timeVal, false)
}

// 日期和时间中间带T方式
func ToDateTimeStringEx2(timeVal time.Time, flagT bool) string {
	if flagT {
		val := timeVal.Format("2006-01-02 15:04:05")
		return strings.Replace(val, " ", "T", -1)
	}

	return timeVal.Format("2006-01-02 15:04:05")
}

// 转换成日期格式
func ToDateTime(timeVal string) (time.Time, error) {
	if stringUtil.IsEmpty(timeVal) {
		return time.Time{}, fmt.Errorf("timeval is empty")
	}

	return time.ParseInLocation("2006-01-02 15:04:05", timeVal, time.Local)
}

// 以指定时区，转换成日期格式
func ToDateTime2(timeVal string, location *time.Location) (time.Time, error) {
	if stringUtil.IsEmpty(timeVal) {
		return time.Time{}, fmt.Errorf("timeval is empty")
	}

	return time.ParseInLocation("2006-01-02 15:04:05", timeVal, location)
}

// 转换成时间格式
func ToDate(timeVal string) (time.Time, error) {
	if stringUtil.IsEmpty(timeVal) {
		return time.Time{}, fmt.Errorf("timeval is empty")
	}

	return time.ParseInLocation("2006-01-02", timeVal, time.Local)
}

// 转换成时间格式
func ToDate2(timeVal string, location *time.Location) (time.Time, error) {
	if stringUtil.IsEmpty(timeVal) {
		return time.Time{}, fmt.Errorf("timeval is empty")
	}

	return time.ParseInLocation("2006-01-02", timeVal, location)
}

// 转换成yyyyMMddHHmmssms的格式
func ToInt64(timeVal time.Time) int64 {
	year := timeVal.Year()
	month := int(timeVal.Month())
	day := timeVal.Day()
	hour := timeVal.Hour()
	minute := timeVal.Minute()
	second := timeVal.Second()

	return int64(int64(year)*1e10) + int64(month*1e8) + int64(day*1e6) + int64(hour*1e4) + int64(minute*1e2) + int64(second)
}
