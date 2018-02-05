package syncUtil

import (
	"sync/atomic"
	"time"
)

// 锁对象
type Locker struct {
	value *int32
}

// 尝试加锁，如果在指定的时间内失败，则会返回失败；否则返回成功
// timeout:指定的毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
func (this *Locker) Lock(timeout int) bool {
	leftTimeout := timeout
	success := false
	for !success {
		success = atomic.CompareAndSwapInt32(this.value, 0, 1)
		if success {
			break
		}

		if timeout > 0 {
			leftTimeout--
			if leftTimeout == 0 {
				break
			}
		}

		time.Sleep(time.Millisecond)
	}

	return success
}

// 锁定（死等方式）
func (this *Locker) WaitLock(){
	this.Lock(-1)
}

// 解锁
func (this *Locker) Unlock() {
	atomic.CompareAndSwapInt32(this.value, 1, 0)
}

// 创建新的锁对象
func NewLocker() *Locker {
	i := int32(0)
	return &Locker{
		value: &i,
	}
}
