package timeUtil

import (
	"testing"
	"time"
)

func TestConverToStandardFormat(t *testing.T) {
	str := "2015-10-10T10:10:10"
	finalDate := ConverToStandardFormat(str)
	expectDate := time.Date(2015, 10, 10, 10, 10, 10, 0, time.Local)

	if finalDate != expectDate {
		t.Errorf("转换不正确，期待：%s, 实际：%s", expectDate, finalDate)
	}
}

func TestConvertToInt(t *testing.T) {
	now := time.Now()
	finalInt := ConvertToInt(now)
	expecteInt := 20160120

	if finalInt != expecteInt {
		t.Errorf("转换不正确，期待：%d, 实际：%d", expecteInt, finalInt)
	}
}
