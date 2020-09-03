package syncUtil

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// 写锁对象
type Locker struct {
	write     int // 使用int而不是bool值的原因，是为了与RWLocker中的read保持类型的一致；
	prevStack []byte
	mutex     sync.Mutex
}

// 内部锁
// 返回值：
// 加锁是否成功
func (this *Locker) lock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 {
		return false
	}

	// 否则，将写锁数量设置为１，并返回成功
	this.write = 1

	// 记录Stack信息
	this.prevStack = debug.Stack()

	return true
}

// 尝试加锁，如果在指定的时间内失败，则会返回失败；否则返回成功
// timeout:指定的毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *Locker) Lock(timeout int) (successful bool, prevStack string, currStack string) {
	timeout = getTimeout(timeout)

	// 遍历指定的次数（即指定的超时时间）
	for i := 0; i < timeout; i = i + con_Lock_Sleep_Millisecond {
		// 如果锁定成功，则返回成功
		if this.lock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠con_Lock_Sleep_Millisecond ms，然后再重试
		time.Sleep(con_Lock_Sleep_Millisecond * time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		prevStack = string(this.prevStack)
		currStack = string(debug.Stack())
	}

	return
}

// 锁定（死等方式）
func (this *Locker) WaitLock() {
	successful, prevStack, currStack := this.Lock(0)
	if successful == false {
		fmt.Printf("Locker.WaitLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 解锁
func (this *Locker) Unlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.write = 0
}

// 创建新的锁对象
func NewLocker() *Locker {
	return &Locker{}
}
