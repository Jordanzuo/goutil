package syncUtil

import (
	"fmt"
	"testing"
)

var (
	lockerUtilObj = NewLockerUtil()
)

func LockerUtil_TestGetLock(t *testing.T) {
	count := 100
	for i := 1; i <= count; i++ {
		lockerUtilObj.GetLock(fmt.Sprintf("%d", i))
		if lockerCount := len(lockerUtilObj.lockerMap); lockerCount != i {
			t.Errorf("(GetLock)Expected %d locker, but now got: %d", count, lockerCount)
		}
	}

	for i := count; i > 0; i-- {
		lockerUtilObj.ReleaseLock(fmt.Sprintf("%d", i))
		if lockerCount := len(lockerUtilObj.lockerMap); lockerCount != i-1 {
			t.Errorf("(ReleaseLock)Expected %d locker, but now got: %d", count, lockerCount)
		}
	}
}
