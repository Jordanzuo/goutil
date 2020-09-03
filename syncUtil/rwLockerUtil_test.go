package syncUtil

import (
	"fmt"
	"testing"
)

var (
	rwLockerUtilObj = NewRWLockerUtil()
)

func RWLockerUtil_TestGetLock(t *testing.T) {
	count := 100
	for i := 1; i <= count; i++ {
		rwLockerUtilObj.GetLock(fmt.Sprintf("%d", i))
		if lockerCount := len(rwLockerUtilObj.lockerMap); lockerCount != i {
			t.Errorf("(GetLock)Expected %d locker, but now got: %d", count, lockerCount)
		}
	}

	for i := count; i > 0; i-- {
		rwLockerUtilObj.ReleaseLock(fmt.Sprintf("%d", i))
		if lockerCount := len(rwLockerUtilObj.lockerMap); lockerCount != i-1 {
			t.Errorf("(ReleaseLock)Expected %d locker, but now got: %d", count, lockerCount)
		}
	}
}
