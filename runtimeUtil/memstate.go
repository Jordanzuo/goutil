package runtimeUtil

import "runtime"

// 获取当前正在使用的内存大小，单位：字节数
// 具体参见文档：
// 1. http://blog.csdn.net/webxscan/article/details/72857292
// 2. https://studygolang.com/static/pkgdoc/pkg/runtime.htm#MemStats
// 返回值:
// int64:正在使用的内存大小
func GetMemSize() uint64 {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	return memStat.Alloc
}
