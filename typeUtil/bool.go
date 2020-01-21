package typeUtil

// bool值转换为int
// value:待转换的值
// 返回值：
// bool:转换结果
func BoolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

// int转换为bool值
// value:待转换的值
// 返回值：
// bool:转换结果
func IntToBool(value int) bool {
	if value > 0 {
		return true
	} else {
		return false
	}
}
