/*
通过在RWLocker对象中引入writeProtectEndTime（写锁保护结束时间），来提高获取写锁的成功率。
当写锁获取失败时，就设置一个写锁保护结束时间，在这段时间内，只允许写锁进行获取，而读锁的获取请求会被拒绝。
通过重置写锁保护结束时间的时机，对写锁的优先级程度进行调整。有两个重置写锁保护结束时间的时机：
１、在成功获取到写锁时：此时重置，有利于下一个写锁需求者在当前写锁持有者处理逻辑时设置保护时间，从而当当前写锁持有者释放锁时，下一个写锁需求者可以立刻获得写锁；
２、在写锁解锁时：此时重置，给了读锁和写锁的需求者同样的机会进行锁的竞争机会；
综上：RWLocker可以提供３中级别的写锁优先级：
１、高级：在获取写锁失败时设置写锁保护结束时间；在获取写锁成功时重置。
２、中级：在获取写锁失败时设置写锁保护结束时间；在释放锁时重置。
３、无：不设置写锁保护时间。
*/
package syncUtil

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// 读写锁对象
type RWLocker struct {
	read                int
	write               int   // 使用int而不是bool值的原因，是为了与read保持类型的一致；
	writeProtectEndTime int64 // 写锁保护结束时间。如果当前时间小于该值，则会阻塞读锁请求；以便于提高写锁的优先级，避免连续的读锁导致写锁无法获得；
	prevStack           []byte
	mutex               sync.Mutex
}

// 尝试加写锁
// 返回值：加写锁是否成功
func (this *RWLocker) lock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败；并且设置写锁保护结束时间；以便于写锁可以优先竞争锁；
	if this.write == 1 || this.read > 0 {
		this.writeProtectEndTime = time.Now().UnixNano() + con_Write_Protect_Nanoseconds
		return false
	}

	// 否则，将写锁数量设置为１，并返回成功；并重置写锁保护结束时间；这样读锁和写锁都可以参与锁的竞争；
	this.write = 1
	this.writeProtectEndTime = time.Now().UnixNano()

	// 记录Stack信息
	if if_record_stack_info {
		this.prevStack = debug.Stack()
	}

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
	for i := 0; i < timeout; i++ {
		// 如果锁定成功，则返回成功
		if this.lock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠1ms，然后再重试
		time.Sleep(time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		if this.prevStack != nil && len(this.prevStack) > 0 {
			prevStack = string(this.prevStack)
		}
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
	// this.writeProtectEndTime = time.Now().UnixNano()
}

// 尝试加读锁
// 返回值：加读锁是否成功
func (this *RWLocker) rlock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，或者处于写锁保护时间段内，则返回失败
	if this.write == 1 || time.Now().UnixNano() < this.writeProtectEndTime {
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
	for i := 0; i < timeout; i++ {
		// 如果锁定成功，则返回成功
		if this.rlock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠1ms，然后再重试
		time.Sleep(time.Millisecond)
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
