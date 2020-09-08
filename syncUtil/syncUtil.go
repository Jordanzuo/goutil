/*
由于Go不提供超时锁，所以自己实现了支持超时机制的互斥锁Locker和读写锁RWLocker。
为了方便供第三方程序使用，提供了根据Key获取超时互斥锁和超时读写锁的复合对象LockerUtil和RWLockerUtil。
为了在出现锁超时时方便查找问题，会记录上次成功获得锁时的堆栈信息；并且在本次获取锁失败时，同时返回上次成功时的堆栈信息和本次的堆栈信息。
*/
package syncUtil

const (
	// 默认超时的毫秒数(1小时)
	con_Default_Timeout_Milliseconds = 60 * 60 * 1000

	// 写锁保护时间（纳秒）
	con_Write_Protect_Nanoseconds = 5 * 1000 * 1000
)

// 获取超时时间
func getTimeout(timeout int) int {
	if timeout > 0 {
		return timeout
	} else {
		return con_Default_Timeout_Milliseconds
	}
}
