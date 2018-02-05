package typeUtil

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Jordanzuo/goutil/timeUtil"
)

// 字节数据类型转换(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// byte:结果
// error:错误数据
func Byte(val interface{}) (byte, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return val.(byte), nil
	case int:
		return byte(val.(int)), nil
	case int32:
		return byte(val.(int32)), nil
	case uint32:
		return byte(val.(uint32)), nil
	case int64:
		return byte(val.(int64)), nil
	case uint64:
		return byte(val.(uint64)), nil
	case int8:
		return byte(val.(int8)), nil
	case int16:
		return byte(val.(int16)), nil
	case uint16:
		return byte(val.(uint16)), nil
	case float32:
		return byte(val.(float32)), nil
	case float64:
		return byte(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return byte(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 类型转换为int32(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// int:结果
// error:错误数据
func Int32(val interface{}) (int32, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return int32(val.(byte)), nil
	case int:
		return int32(val.(int)), nil
	case uint:
		return int32(val.(uint)), nil
	case int32:
		return val.(int32), nil
	case uint32:
		return int32(val.(uint32)), nil
	case int64:
		return int32(val.(int64)), nil
	case uint64:
		return int32(val.(uint64)), nil
	case int8:
		return int32(val.(int8)), nil
	case int16:
		return int32(val.(int16)), nil
	case uint16:
		return int32(val.(uint16)), nil
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

// 转换为int32列表(转换过程不是类型安全的)
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

// 类型转换为uint32(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// int:结果
// error:错误数据
func Uint32(val interface{}) (uint32, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return uint32(val.(byte)), nil
	case int:
		return uint32(val.(int)), nil
	case uint:
		return uint32(val.(uint)), nil
	case int32:
		return uint32(val.(int32)), nil
	case uint32:
		return uint32(val.(uint32)), nil
	case int64:
		return uint32(val.(int64)), nil
	case uint64:
		return uint32(val.(uint64)), nil
	case int8:
		return uint32(val.(int8)), nil
	case int16:
		return uint32(val.(int16)), nil
	case uint16:
		return uint32(val.(uint16)), nil
	case float32:
		return uint32(val.(float32)), nil
	case float64:
		return uint32(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return uint32(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为uint32列表(转换过程不是类型安全的)
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func Uint32Array(val []interface{}) ([]uint32, error) {
	array := make([]uint32, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Uint32(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为int(转换过程不是类型安全的)
// 返回值:
// int:结果
// error:错误数据
func Int(val interface{}) (int, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return int(val.(byte)), nil
	case int:
		return val.(int), nil
	case uint:
		return int(val.(uint)), nil
	case int32:
		return int(val.(int32)), nil
	case uint32:
		return int(val.(uint32)), nil
	case int64:
		return int(val.(int64)), nil
	case uint64:
		return int(val.(uint64)), nil
	case int8:
		return int(val.(int8)), nil
	case int16:
		return int(val.(int16)), nil
	case uint16:
		return int(val.(uint16)), nil
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

// 转换为Int列表(转换过程不是类型安全的)
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

// 类型转换为uint(转换过程不是类型安全的)
// 返回值:
// int:结果
// error:错误数据
func Uint(val interface{}) (uint, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return uint(val.(byte)), nil
	case int:
		return uint(val.(int)), nil
	case uint:
		return uint(val.(uint)), nil
	case int32:
		return uint(val.(int32)), nil
	case uint32:
		return uint(val.(uint32)), nil
	case int64:
		return uint(val.(int64)), nil
	case uint64:
		return uint(val.(uint64)), nil
	case int8:
		return uint(val.(int8)), nil
	case int16:
		return uint(val.(int16)), nil
	case uint16:
		return uint(val.(uint16)), nil
	case float32:
		return uint(val.(float32)), nil
	case float64:
		return uint(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return uint(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为uint列表(转换过程不是类型安全的)
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func UintArray(val []interface{}) ([]uint, error) {
	array := make([]uint, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Uint(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为int64(转换过程不是类型安全的)
// 返回值:
// int:结果
// error:错误数据
func Int64(val interface{}) (int64, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return int64(val.(byte)), nil
	case int:
		return int64(val.(int)), nil
	case uint:
		return int64(val.(uint)), nil
	case int32:
		return int64(val.(int32)), nil
	case uint32:
		return int64(val.(uint32)), nil
	case int64:
		return val.(int64), nil
	case uint64:
		return int64(val.(uint64)), nil
	case int8:
		return int64(val.(int8)), nil
	case int16:
		return int64(val.(int16)), nil
	case uint16:
		return int64(val.(uint16)), nil
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

// 转换为int64列表(转换过程不是类型安全的)
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

// 类型转换为uint64(转换过程不是类型安全的)
// 返回值:
// int:结果
// error:错误数据
func Uint64(val interface{}) (uint64, error) {
	if val == nil {
		return 0, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case byte:
		return uint64(val.(byte)), nil
	case int:
		return uint64(val.(int)), nil
	case uint:
		return uint64(val.(uint)), nil
	case int32:
		return uint64(val.(int32)), nil
	case uint32:
		return uint64(val.(uint32)), nil
	case int64:
		return uint64(val.(int64)), nil
	case uint64:
		return uint64(val.(uint64)), nil
	case int8:
		return uint64(val.(int8)), nil
	case int16:
		return uint64(val.(int16)), nil
	case uint16:
		return uint64(val.(uint16)), nil
	case float32:
		return uint64(val.(float32)), nil
	case float64:
		return uint64(val.(float64)), nil
	case string:
		result, errMsg := strconv.ParseFloat(val.(string), 64)
		if errMsg != nil {
			return 0, fmt.Errorf("string convert error")
		}

		return uint64(result), nil
	}

	return 0, fmt.Errorf("val is not base type")
}

// 转换为uint64列表(转换过程不是类型安全的)
// val:待转换的数据列表
// 返回值:
// []int:结果
// error:错误数据
func Uint64Array(val []interface{}) ([]uint64, error) {
	array := make([]uint64, 0, len(val))
	if val == nil {
		return array, fmt.Errorf("val is nil")
	}

	// 转换成数组
	for _, item := range val {
		tmpResult, errMsg := Uint64(item)
		if errMsg != nil {
			return nil, errMsg
		}

		array = append(array, tmpResult)
	}

	return array, nil
}

// 类型转换为float64(转换过程不是类型安全的)
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
	case uint:
		return float64(val.(uint)), nil
	case int32:
		return float64(val.(int32)), nil
	case uint32:
		return float64(val.(uint32)), nil
	case int64:
		return float64(val.(int64)), nil
	case uint64:
		return float64(val.(uint64)), nil
	case int8:
		return float64(val.(int8)), nil
	case uint8:
		return float64(val.(uint8)), nil
	case int16:
		return float64(val.(int16)), nil
	case uint16:
		return float64(val.(uint16)), nil
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

// 转换为Int列表(转换过程不是类型安全的)
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

// 类型转换为bool(转换过程不是类型安全的)
// 返回值:
// bool:结果
// error:错误数据
func Bool(val interface{}) (bool, error) {
	if val == nil {
		return false, fmt.Errorf("val is nil")
	}

	switch val.(type) {
	case int:
		return (val.(int)) > 0, nil
	case uint:
		return (val.(uint)) > 0, nil
	case int32:
		return (val.(int32)) > 0, nil
	case uint32:
		return (val.(uint32)) > 0, nil
	case int64:
		return (val.(int64)) > 0, nil
	case uint64:
		return (val.(uint64)) > 0, nil
	case int8:
		return (val.(int8)) > 0, nil
	case uint8:
		return (val.(uint8)) > 0, nil
	case int16:
		return (val.(int16)) > 0, nil
	case uint16:
		return (val.(uint16)) > 0, nil
	case float32:
		return int(val.(float32)) > 0, nil
	case float64:
		return int(val.(float64)) > 0, nil
	case bool:
		return val.(bool), nil
	case string:
		result, errMsg := strconv.ParseBool(val.(string))
		if errMsg != nil {
			// 先尝试转换成数值值，再进行bool转换
			var tmpVal float64
			tmpVal, errMsg = strconv.ParseFloat(val.(string), 64)
			if errMsg != nil {
				return false, fmt.Errorf("string convert error")
			}

			result = int(tmpVal) > 0
		}

		return result, nil
	}

	return false, fmt.Errorf("val is not base type")
}

// 转换为Int列表(转换过程不是类型安全的)
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

// 类型转换为字符串(转换过程不是类型安全的)
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
	case uint:
		return string(val.(uint)), nil
	case int32:
		return string(val.(int32)), nil
	case uint32:
		return string(val.(uint32)), nil
	case int64:
		return string(val.(int64)), nil
	case uint64:
		return string(val.(uint64)), nil
	case int8:
		return string(val.(int8)), nil
	case uint8:
		return string(val.(uint8)), nil
	case int16:
		return string(val.(int16)), nil
	case uint16:
		return string(val.(uint16)), nil
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'F', 5, 32), nil
	case float64:
		return strconv.FormatFloat(val.(float64), 'F', 5, 64), nil
	case string:
		return val.(string), nil
	}

	return "", fmt.Errorf("val is not base type")
}

// 转换为Int列表(转换过程不是类型安全的)
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

// 类型转换（基础数据类型）
// val:原始值
// targetType:目标值类型
// 返回值:
// interface{}:结果
// error:错误信息
func Convert(val interface{}, targetType reflect.Kind) (interface{}, error) {
	switch targetType {
	case reflect.Int:
		return Int(val)
	case reflect.Int8:
		{
			val, err := Int(val)
			return int8(val), err
		}
	case reflect.Int16:
		{
			val, err := Int(val)
			return int16(val), err
		}
	case reflect.Int32:
		return Int32(val)
	case reflect.Int64:
		return Int64(val)
	case reflect.Uint:
		return Uint(val)
	case reflect.Uint8:
		{
			val, err := Uint(val)
			return uint8(val), err
		}
	case reflect.Uint16:
		{
			val, err := Uint(val)
			return uint16(val), err
		}
	case reflect.Uint32:
		return Uint32(val)
	case reflect.Uint64:
		return Uint64(val)
	case reflect.Float32:
		{
			val, err := Float64(val)
			return float32(val), err
		}
	case reflect.Float64:
		return Float64(val)
	case reflect.Bool:
		return Bool(val)
	case reflect.String:
		return String(val)
	}

	return nil, fmt.Errorf("Unknown DataType:%s", targetType.String())
}
