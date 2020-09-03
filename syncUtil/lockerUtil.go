package syncUtil

import (
	"sync"
)

// 锁工具类
type LockerUtil struct {
	// 锁集合
	lockerMap map[string]*Locker

	// 锁对象
	rwMutex sync.RWMutex
}

// 创建新的锁工具类
func NewLockerUtil() *LockerUtil {
	return &LockerUtil{
		lockerMap: make(map[string]*Locker, 8),
	}
}

// 获取锁对象
// lockName:锁名
// 返回值：
// *Locker:锁对象
func (this *LockerUtil) GetLock(lockName string) *Locker {
	var lockerObj *Locker
	var exists bool

	func() {
		this.rwMutex.RLock()
		defer this.rwMutex.RUnlock()
		lockerObj, exists = this.lockerMap[lockName]
	}()
	if exists {
		return lockerObj
	}

	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()

	lockerObj, exists = this.lockerMap[lockName]
	if !exists {
		lockerObj = NewLocker()
		this.lockerMap[lockName] = lockerObj
	}

	return lockerObj
}

// 释放锁对象
// lockName:锁名
func (this *LockerUtil) ReleaseLock(lockName string) {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	delete(this.lockerMap, lockName)
}
