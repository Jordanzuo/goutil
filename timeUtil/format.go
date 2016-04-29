package timeUtil

import (
	"strconv"
	"strings"
	"time"
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
