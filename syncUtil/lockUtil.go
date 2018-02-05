package syncUtil

import (
	"sync"
)

type LockUtil struct {
	// 锁集合
	lockData map[string]*Locker

	// 锁对象
	lockObj sync.RWMutex
}

// 创建新的锁工具类
func NewLockUtil() *LockUtil {
	return &LockUtil{
		lockData: make(map[string]*Locker),
	}
}

// 获取锁对象
// lockName:锁名
// 返回值：
// *Locker:锁对象
func (this *LockUtil) GetLock(lockName string) *Locker {
	var lockItem *Locker
	var isExist bool

	func() {
		this.lockObj.RLock()
		defer this.lockObj.RUnlock()
		lockItem, isExist = this.lockData[lockName]
	}()

	if isExist == true {
		return lockItem
	}

	this.lockObj.Lock()
	defer this.lockObj.Unlock()

	lockItem, isExist = this.lockData[lockName]
	if isExist == false {
		lockItem = NewLocker()
		this.lockData[lockName] = lockItem
	}

	return lockItem
}
