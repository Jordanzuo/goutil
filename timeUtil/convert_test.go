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
