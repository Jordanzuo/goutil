package timeUtil

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	timeVal := time.Date(2018, 4, 25, 9, 36, 1, 0, time.Local)
	timeStr1 := ToDateTimeString2(timeVal)

	utcTime := GetUTCTime(timeVal)
	timeStr2 := ToDateTimeString2(utcTime)

	if timeStr1 != timeStr2 {
		t.Errorf("获取UTC时间出错，两个时间不对等")
	}

	utcTime2 := GetUTCTime(utcTime)
	timeStr3 := ToDateTimeString2(utcTime2)
	if timeStr1 != timeStr3 {
		t.Errorf("两次的UTC时间不对等")
	}

	utcTime4 := GetLocalTime(utcTime)
	timeStr4 := ToDateTimeString2(utcTime4)
	if timeStr4 != timeStr1 {
		t.Errorf("local变更了时间 time1:%v time4:%v", timeStr1, timeStr4)
	}

	utcTime5 := GetLocalTime(utcTime)
	timeStr5 := ToDateTimeString2(utcTime5)
	if timeStr4 != timeStr5 {
		t.Errorf("两次的local时间不对等")
	}
}
