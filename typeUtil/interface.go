package typeutil

import (
	"strconv"
)

// 类型转换为int
// 返回值:
// int:结果
// bool:是否转换成功
func Int(val interface{}) (int, bool) {
	if val == nil {
		return 0, false
	}

	switch val.(type) {
	case int:
		return val.(int), true
	case int32:
		return int(val.(int32)), true
	case int64:
		return int(val.(int64)), true
	case int8:
		return int(val.(int8)), true
	case int16:
		return int(val.(int16)), true
	case float32:
		return int(val.(float32)), true
	case float64:
		return int(val.(float64)), true
	case string:
		result, errMsg := strconv.Atoi(val.(string))
		if errMsg != nil {
			return 0, false
		}

		return result, true
	}

	return 0, false
}

// 类型转换为float64
// 返回值:
// float64:结果
// bool:是否转换成功
func Float64(val interface{}) (float64, bool) {

	if val == nil {
		return 0, false
	}

	switch val.(type) {
	case int:
		return float64(val.(int)), true
	case int32:
		return float64(val.(int32)), true
	case int64:
		return float64(val.(int64)), true
	case int8:
		return float64(val.(int8)), true
	case int16:
		return float64(val.(int16)), true
	case float32:
		return float64(val.(float32)), true
	case float64:
		return float64(val.(float64)), true
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, false
		}

		return result, true
	}

	return 0, false
}

// 类型转换为bool
// 返回值:
// bool:结果
// bool:是否转换成功
func Bool(val interface{}) (bool, bool) {
	if val == nil {
		return false, false
	}

	switch val.(type) {
	case int:
		return int(val.(int)) > 0, true
	case int32:
		return int32(val.(int32)) > 0, true
	case int64:
		return int64(val.(int64)) > 0, true
	case int8:
		return int8(val.(int8)) > 0, true
	case int16:
		return int8(val.(int16)) > 0, true
	case float32:
		return int(val.(float32)) > 0, true
	case float64:
		return int(val.(float64)) > 0, true
	case string:
		result, errMsg := strconv.ParseBool(val.(string))
		if errMsg != nil {
			return false, false
		}

		return result, true
	}

	return false, false
}

// 类型转换为字符串
// 返回值:
// string:结果
// bool:是否转换成功
func String(val interface{}) (string, bool) {
	if val == nil {
		return "", false
	}

	switch val.(type) {
	case int:
		return string(val.(int)), true
	case int32:
		return string(val.(int32)), true
	case int64:
		return string(val.(int64)), true
	case int8:
		return string(val.(int8)), true
	case int16:
		return string(val.(int16)), true
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'F', 5, 32), true
	case float64:
		return strconv.FormatFloat(val.(float64), 'F', 5, 64), true
	case string:
		return val.(string), true
	}

	return "", false
}
