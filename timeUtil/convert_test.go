package timeUtil

import (
	"testing"
	"time"
)

func TestConverToStandardFormat(t *testing.T) {
	str := "2018-10-10T10:10:10"
	expected := time.Date(2018, 10, 10, 10, 10, 10, 0, time.Local)

	got, err := ConverToStandardFormat(str)
	if err != nil {
		t.Errorf("发生错误，错误信息为：%s", err)
	}

	if got != expected {
		t.Errorf("转换不正确，期待：%s, 实际：%s", expected, got)
	}
}

func TestConvertToInt(t *testing.T) {
	date := time.Date(2018, 10, 10, 10, 10, 10, 0, time.Local)
	finalInt := ConvertToInt(date)
	expecteInt := 20181010

	if finalInt != expecteInt {
		t.Errorf("转换不正确，期待：%d, 实际：%d", expecteInt, finalInt)
	}
}

func TestSubDay(t *testing.T) {
	time1 := time.Now().AddDate(0, 0, 5)
	time2 := time.Now()
	expected := 5

	got := SubDay(time1, time2)
	if got != expected {
		t.Errorf("Expected %d, but now got %d.", expected, got)
	}
}

func TestParseTimeString(t *testing.T) {
	val := "12:13:14"
	expectedHour := 12
	expectedMinute := 13
	expectedSecond := 14

	err, hour, miniute, second := ParseTimeString(val)
	if err != nil {
		t.Error(err)
	}

	if expectedHour != hour || expectedMinute != miniute || expectedSecond != second {
		t.Fail()
	}
}
