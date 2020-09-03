package syncUtil

import (
	"sync"
)

// 读写锁工具类
type RWLockerUtil struct {
	// 锁集合
	lockerMap map[string]*RWLocker

	// 锁对象
	rwMutex sync.RWMutex
}

// 创建新的锁工具类
func NewRWLockerUtil() *RWLockerUtil {
	return &RWLockerUtil{
		lockerMap: make(map[string]*RWLocker, 8),
	}
}

// 获取锁对象
// lockName:锁名
// 返回值：
// RWLocker:读写锁对象
func (this *RWLockerUtil) GetLock(lockName string) *RWLocker {
	var rwLockerObj *RWLocker
	var exists bool

	func() {
		this.rwMutex.RLock()
		defer this.rwMutex.RUnlock()
		rwLockerObj, exists = this.lockerMap[lockName]
	}()
	if exists {
		return rwLockerObj
	}

	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()

	rwLockerObj, exists = this.lockerMap[lockName]
	if exists == false {
		rwLockerObj = NewRWLocker()
		this.lockerMap[lockName] = rwLockerObj
	}

	return rwLockerObj
}

// 释放读写锁对象
// lockName:锁名
func (this *RWLockerUtil) ReleaseLock(lockName string) {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	delete(this.lockerMap, lockName)
}
