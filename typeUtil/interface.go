package typeUtil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Jordanzuo/goutil/timeUtil"
)

// 类型转换为int32
// 返回值:
// int:结果
// error:错误数据
func Int32(val interface{}) (int32, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return int32(val.(int)), nil
	case int32:
		return val.(int32), nil
	case int64:
		return int32(val.(int64)), nil
	case int8:
		return int32(val.(int8)), nil
	case int16:
		return int32(val.(int16)), nil
	case float32:
		return int32(val.(float32)), nil
	case float64:
		return int32(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return int32(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为Int列表
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func Int32Array(val []interface{}) ([]int32, error) {
	array := make([]int32, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Int32(item)
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

// 转换成时间格式
// val:待转换的数据,如果是字符串，则要求是格式:2006-01-02 15:04:05
// *time.Time:转换结果
// error:转换的错误信息
func DateTime(val interface{}) (time.Time, error) {
	if val == nil {
		return time.Time{}, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case time.Time:
		return val.(time.Time), nil
	case string:
		return timeUtil.ToDateTime(val.(string))
	case int, int64, float32, float64:
		intVal, _ := Int64(val)
		return time.Unix(intVal, 0).Local(), nil
	default:
		return time.Time{}, fmt.Errorf("unknown data type")
	}
}

// 转换成时间格式
// val:待转换的数据,如果是字符串，则要求是格式:2006-01-02 15:04:05
// *time.Time:转换结果
// error:转换的错误信息
func DateTimeArray(val []interface{}) ([]time.Time, error) {
	array := make([]time.Time, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := DateTime(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, fmt.Errorf("val is not base type")
}

// 转换成时间格式
// val:待转换的数据,如果是字符串，则使用format进行转换
// format:时间格式
// *time.Time:转换结果
// error:转换的错误信息
func DateTimeByFormat(val interface{}, format string) (time.Time, error) {
	if val == nil {
		return time.Time{}, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case time.Time:
		return val.(time.Time), nil
	case string:
		return time.ParseInLocation(val.(string), format, time.Local)
	case int, int64, float32, float64:
		intVal, _ := Int64(val)
		return time.Unix(intVal, 0).Local(), nil
	default:
		return time.Time{}, fmt.Errorf("unknown data type")
	}
}

// 转换成时间格式
// val:待转换的数据,如果是字符串，则使用format进行转换
// format:时间格式
// *time.Time:转换结果
// error:转换的错误信息
func DateTimeArrayByFormat(val []interface{}, format string) ([]time.Time, error) {
	array := make([]time.Time, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := DateTimeByFormat(item, format)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, fmt.Errorf("val is not base type")
}
