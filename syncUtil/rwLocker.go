package syncUtil

import (
	"sync/atomic"
	"time"
)

// 读写锁对象
type RWLocker struct {
	read  *int32
	write *int32
}

// 写锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
func (this *RWLocker) Lock(timeout int) (success bool) {
	// 写锁优先级更高，所以每次休眠的时间更短，且可以预增加
	interval := 2 * time.Millisecond
	occupied := false
	leftTimeout := timeout

	defer func() {
		// 如果锁失败，且已经预占用了写锁，则将其释放
		if !success && occupied {
			atomic.CompareAndSwapInt32(this.write, 1, 0)
		}
	}()

	for {
		// 如果不是死等，则需要计算超时
		if timeout > 0 {
			// 由于是先扣除时间，所以判断timeout时使用timeout<0，而不是timeout<=0
			leftTimeout--
			if leftTimeout < 0 {
				return
			}
		}

		// 如果写锁没有被占用，则预占用；否则等待下次判断
		if *this.write == 0 {
			// 如果预占用失败，表示被另一个写请求占用；则先休眠，稍后再判断
			if atomic.CompareAndSwapInt32(this.write, 0, 1) == false {
				time.Sleep(interval)
				continue
			} else {
				occupied = true
			}
		} else {
			time.Sleep(interval)
			continue
		}

		// 判断当前的读锁数量是否为0，如果为0表示此次锁定成功，否则需要等待
		if *this.read == 0 {
			success = true
			return
		} else {
			time.Sleep(interval)
		}
	}

	return
}

// 写锁定(死等)
func (this *RWLocker) WaitLock() {
	this.Lock(-1)
}

// 解写锁
func (this *RWLocker) Unlock() {
	atomic.CompareAndSwapInt32(this.write, 1, 0)
}

// 读锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
func (this *RWLocker) RLock(timeout int) (success bool) {
	interval := 3 * time.Millisecond
	leftTimeout := timeout

	for {
		if timeout > 0 {
			// 由于是先扣除时间，所以判断timeout时使用timeout<0，而不是timeout<=0
			leftTimeout--
			if leftTimeout < 0 {
				return
			}
		}

		// 如果已经有写锁，则等待
		if *this.write == 1 {
			time.Sleep(interval)
			continue
		}

		// 如果没有写锁，则将读+1
		atomic.AddInt32(this.read, 1)

		// 再次判断是否有写锁，如果有，则将读-1；并重新进行循环判断
		if *this.write == 1 {
			atomic.AddInt32(this.read, -1)
			time.Sleep(interval)
			continue
		}

		success = true
		return
	}

	return
}

// 读锁定(死等)
func (this *RWLocker) WaitRLock() {
	this.RLock(-1)
}

// 解读锁
func (this *RWLocker) RUnlock() {
	atomic.AddInt32(this.read, -1)
}

// 创建新的读写锁对象
func NewRWLocker() *RWLocker {
	read, write := int32(0), int32(0)
	return &RWLocker{
		read:  &read,
		write: &write,
	}
}
