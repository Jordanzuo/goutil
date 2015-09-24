package mathUtil

import (
	"fmt"
)

func GetSizeDesc(size int64) string {
	str := ""
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
