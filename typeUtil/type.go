package typeUtil

// bool值转换为int
// val:待转换的值
// 返回值：
// bool:转换结果
func BoolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

// int转换为bool值
// val:待转换的值
// 返回值：
// bool:转换结果
func IntToBool(val int) bool {
	if val > 0 {
		return true
	} else {
		return false
	}
}
