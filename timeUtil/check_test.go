package timeUtil

import (
	"testing"
)

// 检查是否在范围内的常规测试
func TestCheckIfInRange(t *testing.T) {
	// Check with correct data
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

	if CheckIfInRange(t1, timeSpan2, 0, timeSpan1, 0) == false {
		t.Fail()
		return
	}

	if CheckIfInRange(t1, timeSpan1, 0, timeSpan2, 5) == false {
		t.Fail()
		return
	}

	if CheckIfInRange(t1, timeSpan2, 0, timeSpan1, 5) == false {
		t.Fail()
		return
	}

	// Check with incorrect data
	t1, err = ToDateTime("2017-09-04 22:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInRange(t1, timeSpan1, 0, timeSpan2, 0) {
		t.Fail()
		return
	}

	if CheckIfInRange(t1, timeSpan2, 0, timeSpan1, 0) {
		t.Fail()
		return
	}
}

func TestCheckIfInRange2(t *testing.T) {
	timeSpan1 := "13:00:00"
	timeSpan2 := "21:00:00"

	t1, err := ToDateTime("2017-09-04 15:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInRange2(t1, timeSpan1, timeSpan2) == false {
		t.Fail()
		return
	}

	if CheckIfInRange2(t1, timeSpan2, timeSpan1) == false {
		t.Fail()
		return
	}

	if CheckIfInRange2(t1, timeSpan1, timeSpan2) == false {
		t.Fail()
		return
	}

	if CheckIfInRange2(t1, timeSpan2, timeSpan1) == false {
		t.Fail()
		return
	}

	t1, err = ToDateTime("2017-09-04 22:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInRange2(t1, timeSpan1, timeSpan2) {
		t.Fail()
		return
	}

	if CheckIfInRange2(t1, timeSpan2, timeSpan1) {
		t.Fail()
		return
	}
}

func TestCheckIfInSameDate(t *testing.T) {
	// Check with the same date time
	t1, err := ToDateTime("2017-09-04 00:00:00")
	if err != nil {
		t.Error(err)
		return
	}
	t2, err := ToDateTime("2017-09-04 15:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInSameDate(t1, t2) == false {
		t.Fail()
		return
	}

	// Check with different date time
	t2, err = ToDateTime("2019-12-26 15:00:00")
	if err != nil {
		t.Error(err)
		return
	}

	if CheckIfInSameDate(t1, t2) {
		t.Fail()
		return
	}
}
