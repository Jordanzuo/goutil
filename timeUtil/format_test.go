package timeUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Date(2015, 9, 11, 10, 10, 10, 0, time.Local)
	expectedString := "2015/09/11"
	result := Format(now, "yyyy/MM/dd")
	if result != expectedString {
		t.Errorf("Format Error, expected %s, Got %s", expectedString, result)
	}

	expectedString = "2015-09-11 %d:%s:%d"
	minutes := ""
	if now.Minute() >= 10 {
		minutes = fmt.Sprintf("%d", now.Minute())
	} else {
		minutes = fmt.Sprintf("0%d", now.Minute())
	}
	expectedString = fmt.Sprintf(expectedString, now.Hour(), minutes, now.Second())
	result = Format(now, "yyyy-MM-dd HH:mm:ss")
	if result != expectedString {
		t.Errorf("Format Error, expected %s, Got %s", expectedString, result)
	}
}
