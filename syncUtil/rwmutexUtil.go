package syncUtil

import (
	"sync"
)

type RWMutexUtil struct {
	// 锁集合
	lockData map[string]*RWLocker

	// 锁对象
	lockObj sync.RWMutex
}

// 创建新的锁工具类
func NewRWMutexUtil() *RWMutexUtil {
	return &RWMutexUtil{
		lockData: make(map[string]*RWLocker),
	}
}

// 获取锁对象
// lockName:锁名
// 返回值：
// sync.RWMutex:锁对象
func (this *RWMutexUtil) GetLock(lockName string) *RWLocker {
	var lockItem *RWLocker
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
		lockItem = NewRWLocker()
		this.lockData[lockName] = lockItem
	}

	return lockItem
}
