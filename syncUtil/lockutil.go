package syncUtil

import (
	"sync"
)

type LockUtil struct {

	// 锁集合
	lockData map[string]sync.Locker

	// 锁对象
	lockObj sync.RWMutex
}

// 创建新的锁工具类
func NewLockUtil() *LockUtil {
	return &LockUtil{
		lockData: make(map[string]sync.Locker),
	}
}

// 获取锁对象
// lockName:锁名
// 返回值：
// sync.Locker:锁对象
func (this *LockUtil) GetLock(lockName string) sync.Locker {
	this.lockObj.RLock()

	lockItem, isExist := this.lockData[lockName]

	this.lockObj.RUnlock()

	if isExist == true {
		return lockItem
	}

	this.lockObj.Lock()
	defer this.lockObj.Unlock()

	lockItem, isExist = this.lockData[lockName]
	if isExist == false {
		lockItem = new(sync.RWMutex)
		this.lockData[lockName] = lockItem
	}

	return lockItem
}

// 释放锁对象
// lockName:锁名
func (this *LockUtil) ReleaseLock(lockName string) {
	this.lockObj.Lock()
	defer this.lockObj.Unlock()

	delete(this.lockData, lockName)
}
