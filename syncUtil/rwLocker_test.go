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

func TestRWNewLocker5(t *testing.T) {
	count := 100
	rwLockerObj := NewRWLocker()
	ch := make(chan bool, 100)

	for i := 0; i < count; i++ {
		go func(num int, ch chan bool) {
			if num%2 == 0 {
				if successful, _, _ := rwLockerObj.Lock(100); successful {
					fmt.Println("I get write lock.")
					time.Sleep(2 * time.Millisecond)
					rwLockerObj.Unlock()
				} else {
					fmt.Println("Write lock timeout")
				}
			} else {
				if successful, _, _ := rwLockerObj.RLock(100); successful {
					fmt.Println("I get read lock.")
					time.Sleep(2 * time.Millisecond)
					rwLockerObj.RUnlock()
				} else {
					fmt.Println("Read lock timeout")
				}
			}
			ch <- true
		}(i, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
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
