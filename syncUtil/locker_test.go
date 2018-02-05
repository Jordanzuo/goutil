package syncUtil

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	count       = 10000000
	round       = 3
	resultCount = 0
)

func init() {
	chCount := 1
	ch := make(chan bool, chCount)

	obj := newT4()
	go benchmark(obj, "T4", count, ch)

	for i := 0; i < chCount; i++ {
		_ = <-ch
	}

	resultCount = int(obj.count)
}

func TestNewLocker(t *testing.T) {
	if round*count != resultCount {
		t.Errorf("Expected %d, but got %d.", count, resultCount)
	}
}

func benchmark(obj IIncrease, name string, count int, ch chan bool) {
	start := time.Now().Unix()

	var wg sync.WaitGroup
	wg.Add(round)

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			obj.Increase()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			obj.Increase()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			obj.Increase()
		}
	}()

	wg.Wait()

	end := time.Now().Unix()

	fmt.Printf("%s:%s use:%d\n", name, obj.ToString(), end-start)

	ch <- true
}

type IIncrease interface {
	Increase()
	ToString() string
}

type T4 struct {
	count int32
	lock  *Locker
}

func newT4() *T4 {
	return &T4{
		count: 0,
		lock:  NewLocker(),
	}
}
func (this *T4) Increase() {
	if this.lock.Lock(20) == false {
		fmt.Println("T4 Lock failed.")
		return
	}
	defer this.lock.Unlock()
	this.count++
}

func (this *T4) ToString() string {
	return fmt.Sprintf("%d", this.count)
}
