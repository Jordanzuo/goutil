package typeUtil

import (
	"fmt"
	"strconv"
)

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Int(val interface{}) (int, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return val.(int), nil
	case int32:
		return int(val.(int32)), nil
	case int64:
		return int(val.(int64)), nil
	case int8:
		return int(val.(int8)), nil
	case int16:
		return int(val.(int16)), nil
	case float32:
		return int(val.(float32)), nil
	case float64:
		return int(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return int(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func IntArray(val []interface{}) ([]int, error) {
	array := make([]int, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Int(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Int64(val interface{}) (int64, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return int64(val.(int)), nil
	case int32:
		return int64(val.(int32)), nil
	case int64:
		return val.(int64), nil
	case int8:
		return int64(val.(int8)), nil
	case int16:
		return int64(val.(int16)), nil
	case float32:
		return int64(val.(float32)), nil
	case float64:
		return int64(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return int64(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func Int64Array(val []interface{}) ([]int64, error) {
	array := make([]int64, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Int64(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为float64
// 返回值:
// float64:结果
// error:错误数据
func Float64(val interface{}) (float64, error) {

	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return float64(val.(int)), nil
	case int32:
		return float64(val.(int32)), nil
	case int64:
		return float64(val.(int64)), nil
	case int8:
		return float64(val.(int8)), nil
	case int16:
		return float64(val.(int16)), nil
	case float32:
		return float64(val.(float32)), nil
	case float64:
		return float64(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return result, nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func Float64Array(val []interface{}) ([]float64, error) {
	array := make([]float64, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Float64(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为bool
// 返回值:
// bool:结果
// error:错误数据
func Bool(val interface{}) (bool, error) {
	if val == nil {
		return false, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return int(val.(int)) > 0, nil
	case int32:
		return int32(val.(int32)) > 0, nil
	case int64:
		return int64(val.(int64)) > 0, nil
	case int8:
		return int8(val.(int8)) > 0, nil
	case int16:
		return int8(val.(int16)) > 0, nil
	case float32:
		return int(val.(float32)) > 0, nil
	case float64:
		return int(val.(float64)) > 0, nil
	case string:
		result, errMsg := strconv.ParseBool(val.(string))
		if errMsg != nil {
			return false, fmt.Errorf("string convert error")
		}

		return result, nil
	}

	return false, fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []bool:结果
// error:错误数据
func BoolArray(val []interface{}) ([]bool, error) {
	array := make([]bool, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Bool(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为字符串
// 返回值:
// string:结果
// error:错误数据
func String(val interface{}) (string, error) {
	if val == nil {
		return "", fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return string(val.(int)), nil
	case int32:
		return string(val.(int32)), nil
	case int64:
		return string(val.(int64)), nil
	case int8:
		return string(val.(int8)), nil
	case int16:
		return string(val.(int16)), nil
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'F', 5, 32), nil
	case float64:
		return strconv.FormatFloat(val.(float64), 'F', 5, 64), nil
	case string:
		return val.(string), nil
	}

	return "", fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []string:结果
// error:错误数据
func StringArray(val []interface{}) ([]string, error) {
	array := make([]string, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := String(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, fmt.Errorf("val is not base type")
}
