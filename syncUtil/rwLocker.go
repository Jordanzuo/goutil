package syncUtil

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// 读写锁对象
type RWLocker struct {
	read      int
	write     int // 使用int而不是bool值的原因，是为了与read保持类型的一致；
	prevStack []byte
	mutex     sync.Mutex
}

// 尝试加写锁
// 返回值：加写锁是否成功
func (this *RWLocker) lock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 || this.read > 0 {
		return false
	}

	// 否则，将写锁数量设置为１，并返回成功
	this.write = 1

	// 记录Stack信息
	this.prevStack = debug.Stack()

	return true
}

// 写锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *RWLocker) Lock(timeout int) (successful bool, prevStack string, currStack string) {
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

// 写锁定(死等)
func (this *RWLocker) WaitLock() {
	successful, prevStack, currStack := this.Lock(0)
	if successful == false {
		fmt.Printf("RWLocker:WaitLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 释放写锁
func (this *RWLocker) Unlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.write = 0
}

// 尝试加读锁
// 返回值：加读锁是否成功
func (this *RWLocker) rlock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 {
		return false
	}

	// 否则，将读锁数量加１，并返回成功
	this.read += 1

	// 记录Stack信息
	this.prevStack = debug.Stack()

	return true
}

// 读锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *RWLocker) RLock(timeout int) (successful bool, prevStack string, currStack string) {
	timeout = getTimeout(timeout)

	// 遍历指定的次数（即指定的超时时间）
	// 读锁比写锁优先级更低，所以每次休眠2ms，所以尝试的次数就是时间/2
	for i := 0; i < timeout; i = i + con_RLock_Sleep_Millisecond {
		// 如果锁定成功，则返回成功
		if this.rlock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠2ms，然后再重试
		time.Sleep(con_RLock_Sleep_Millisecond * time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		prevStack = string(this.prevStack)
		currStack = string(debug.Stack())
	}

	return
}

// 读锁定(死等)
func (this *RWLocker) WaitRLock() {
	successful, prevStack, currStack := this.RLock(0)
	if successful == false {
		fmt.Printf("RWLocker:WaitRLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 释放读锁
func (this *RWLocker) RUnlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.read > 0 {
		this.read -= 1
	}
}

// 创建新的读写锁对象
func NewRWLocker() *RWLocker {
	return &RWLocker{}
}
