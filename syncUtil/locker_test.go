package syncUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLocker1(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 1

	lockerObj := NewLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go lockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestNewLocker2(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 100

	lockerObj := NewLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go lockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestNewLocker3(t *testing.T) {
	count := 1000000
	succeedCount := 0
	expected := 1000000
	goroutineCount := 10000

	lockerObj := NewLocker()
	ch := make(chan bool, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go lockerTest(lockerObj, &succeedCount, count/goroutineCount, ch)
	}

	for i := 0; i < goroutineCount; i++ {
		<-ch
	}

	if succeedCount != expected {
		t.Errorf("Expected %d, but got %d", expected, succeedCount)
	}
}

func TestNewLocker4(t *testing.T) {
	lockerObj := NewLocker()
	if successful, _, _ := lockerObj.Lock(100); successful == false {
		t.Errorf("Lock should be successful, but now it fails.")
	}

	if successful, _, _ := lockerObj.Lock(100); successful {
		t.Errorf("Lock should be failed, but now it succeeds.")
	}
}

func lockerTest(lockerObj *Locker, succeedCount *int, count int, ch chan bool) {
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
