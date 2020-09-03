package syncUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestRWNewLocker1(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 1

	lockerObj := NewRWLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go rwLockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestRWNewLocker2(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 100

	lockerObj := NewRWLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go rwLockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestRWNewLocker3(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 10000

	lockerObj := NewRWLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go rwLockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestRWNewLocker4(t *testing.T) {
	lockerObj := NewRWLocker()
	if successful, _, _ := lockerObj.RLock(100); successful == false {
		t.Errorf("It should be successful to get a read lock, but now it fails.")
		return
	}
	if successful, _, _ := lockerObj.RLock(100); successful == false {
		t.Errorf("It should be successful to get a read lock, but now it fails.")
		return
	}
	if successful, _, _ := lockerObj.RLock(100); successful == false {
		t.Errorf("It should be successful to get a read lock, but now it fails.")
		return
	}
	lockerObj.RUnlock()
	lockerObj.RUnlock()
	lockerObj.RUnlock()

	if successful, _, _ := lockerObj.Lock(100); successful == false {
		t.Errorf("It should be successful to get a write lock, but now it fails.")
		return
	}
	if successful, _, _ := lockerObj.Lock(100); successful {
		t.Errorf("It should be failed to get a write lock, but now it succeeds.")
		return
	}
	if successful, _, _ := lockerObj.RLock(100); successful {
		t.Errorf("It should be failed to get a read lock, but now it succeeds.")
		return
	}

	lockerObj.Unlock()
}

func rwLockerTest(lockerObj *RWLocker, succeedCount *int, count int, ch chan bool) {
	if success, _, _ := lockerObj.Lock(10000); !success {
		fmt.Printf("[%v]获取锁超时\n", time.Now())
		return
	}
	defer lockerObj.Unlock()

	for i := 0; i < count; i++ {
		*succeedCount += 1
	}

	ch <- true
}
