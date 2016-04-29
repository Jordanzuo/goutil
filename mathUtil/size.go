package mathUtil

import (
	"fmt"
)

// 获取字节大小的描述信息
// size：字节大小
// 返回值：
// 描述信息
func GetSizeDesc(size int64) string {
	str := ""

	// 判断输入是否超过int64的范围
	if size < 0 || size > (1<<63-1) {
		return str
	}

	switch {
	case size >= 1024*1024*1024*1024*1024*1024:
		str = fmt.Sprintf("%.2fEB", float64(size)/1024/1024/1024/1024/1024/1024)
	case size >= 1024*1024*1024*1024*1024:
		str = fmt.Sprintf("%.2fPB", float64(size)/1024/1024/1024/1024/1024)
	case size >= 1024*1024*1024*1024:
		str = fmt.Sprintf("%.2fTB", float64(size)/1024/1024/1024/1024)
	case size >= 1024*1024*1024:
		str = fmt.Sprintf("%.2fGB", float64(size)/1024/1024/1024)
	case size >= 1024*1024:
		str = fmt.Sprintf("%dMB", size/1024/1024)
	case size >= 1024:
		str = fmt.Sprintf("%dKB", size/1024)
	default:
		str = fmt.Sprintf("%dB", size)
	}

	return str
}
