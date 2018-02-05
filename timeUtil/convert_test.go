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
	expecteInt := 20170712

	if finalInt != expecteInt {
		t.Errorf("转换不正确，期待：%d, 实际：%d", expecteInt, finalInt)
	}
}

func TestSubDay(t *testing.T) {
	time1 := time.Now().AddDate(0, 0, 5)
	time2 := time.Now()
	expected := 5

	if actual := SubDay(time1, time2); actual != expected {
		t.Errorf("Expected %d, but now got %d.", expected, actual)
	}
}

func TestParseTimeString(t *testing.T) {
	val := "12:13:14"
	err, hour, miniute, second := ParseTimeString(val)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Hour:%v Miniute:%v Second:%v", hour, miniute, second)
}
