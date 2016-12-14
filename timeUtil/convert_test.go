package timeUtil

import (
	"testing"
	"time"
)

func TestConverToStandardFormat(t *testing.T) {
	str := "2018-10-10T10:10:10"
	finalDate, err := ConverToStandardFormat(str)
	if err != nil {
		t.Errorf("发生错误，错误信息为：%s", err)
	}
	expectDate := time.Date(2018, 10, 10, 10, 10, 10, 0, time.Local)

	if finalDate != expectDate {
		t.Errorf("转换不正确，期待：%s, 实际：%s", expectDate, finalDate)
	}
}

func TestConvertToInt(t *testing.T) {
	now := time.Now()
	finalInt := ConvertToInt(now)
	expecteInt := 20161020

	if finalInt != expecteInt {
		t.Errorf("转换不正确，期待：%d, 实际：%d", expecteInt, finalInt)
	}
}
