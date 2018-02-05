package timeUtil

import (
	"testing"
)

// 检查是否在范围内的常规测试
func TestCheckIfInRange(t *testing.T) {
	timeSpan1 := "13:00:00"
	timeSpan2 := "21:00:00"

	t1, err := ToDateTime("2017-09-04 15:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInRange(t1, timeSpan1, 0, timeSpan2, 0) == false {
		t.Fail()
		return
	}
}
