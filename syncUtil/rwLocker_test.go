package syncUtil

import (
	"fmt"
	"testing"
)

func init() {
	chCount := 1
	ch := make(chan bool, chCount)

	obj := newT5()
	go benchmark(obj, "T5", count, ch)

	for i := 0; i < chCount; i++ {
		_ = <-ch
	}

	resultCount = int(obj.count)
}

func TestNewRWLocker(t *testing.T) {
	if round*count != resultCount {
		t.Errorf("Expected %d, but got %d.", count, resultCount)
	}
}

type T5 struct {
	count int32
	lock  *RWLocker
}

func newT5() *T5 {
	return &T5{
		count: 0,
		lock:  NewRWLocker(),
	}
}
func (this *T5) Increase() {
	if this.lock.Lock(20) == false {
		fmt.Println("T5 Lock failed.")
		return
	}
	defer this.lock.Unlock()
	this.count++
}

func (this *T5) ToString() string {
	return fmt.Sprintf("%d", this.count)
}
