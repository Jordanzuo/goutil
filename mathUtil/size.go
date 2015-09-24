package mathUtil

import (
	"fmt"
)

func GetSizeDesc(size int64) string {
	str := ""
	switch {
	case size >= 1024*1024*1024*1024*1024*1024:
		str = fmt.Sprintf("%dEB", size/1024/1024/1024/1024/1024/1024)
	case size >= 1024*1024*1024*1024*1024:
		str = fmt.Sprintf("%dPB", size/1024/1024/1024/1024/1024)
	case size >= 1024*1024*1024*1024:
		str = fmt.Sprintf("%dTB", size/1024/1024/1024/1024)
	case size >= 1024*1024*1024:
		str = fmt.Sprintf("%dGB", size/1024/1024/1024)
	case size >= 1024*1024:
		str = fmt.Sprintf("%dMB", size/1024/1024)
	case size >= 1024:
		str = fmt.Sprintf("%dKB", size/1024)
	default:
		str = fmt.Sprintf("%dB", size)
	}

	return str
}
