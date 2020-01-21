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
// result:结果
// err:错误数据
func Byte(val interface{}) (result byte, err error) {
	return Uint8(val)
}

// 字节数据类型转换(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func ByteArray(valArray []interface{}) (result []uint8, err error) {
	return Uint8Array(valArray)
}

// 类型转换为int(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Int(val interface{}) (result int, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = int(val.(int))
	case uint:
		result = int(val.(uint))
	case int8:
		result = int(val.(int8))
	case uint8:
		result = int(val.(uint8))
	case int16:
		result = int(val.(int16))
	case uint16:
		result = int(val.(uint16))
	case int32:
		result = int(val.(int32))
	case uint32:
		result = int(val.(uint32))
	case int64:
		result = int(val.(int64))
	case uint64:
		result = int(val.(uint64))
	case float32:
		result = int(val.(float32))
	case float64:
		result = int(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = int(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为Int列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func IntArray(valArray []interface{}) (result []int, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]int, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Int(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为int8(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Int8(val interface{}) (result int8, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = int8(val.(int))
	case uint:
		result = int8(val.(uint))
	case int8:
		result = int8(val.(int8))
	case uint8:
		result = int8(val.(uint8))
	case int16:
		result = int8(val.(int16))
	case uint16:
		result = int8(val.(uint16))
	case int32:
		result = int8(val.(int32))
	case uint32:
		result = int8(val.(uint32))
	case int64:
		result = int8(val.(int64))
	case uint64:
		result = int8(val.(uint64))
	case float32:
		result = int8(val.(float32))
	case float64:
		result = int8(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = int8(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为int8列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Int8Array(valArray []interface{}) (result []int8, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]int8, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Int8(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为int16(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Int16(val interface{}) (result int16, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = int16(val.(int))
	case uint:
		result = int16(val.(uint))
	case int8:
		result = int16(val.(int8))
	case uint8:
		result = int16(val.(uint8))
	case int16:
		result = int16(val.(int16))
	case uint16:
		result = int16(val.(uint16))
	case int32:
		result = int16(val.(int32))
	case uint32:
		result = int16(val.(uint32))
	case int64:
		result = int16(val.(int64))
	case uint64:
		result = int16(val.(uint64))
	case float32:
		result = int16(val.(float32))
	case float64:
		result = int16(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = int16(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为int16列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Int16Array(valArray []interface{}) (result []int16, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]int16, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Int16(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为int32(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Int32(val interface{}) (result int32, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = int32(val.(int))
	case uint:
		result = int32(val.(uint))
	case int8:
		result = int32(val.(int8))
	case uint8:
		result = int32(val.(uint8))
	case int16:
		result = int32(val.(int16))
	case uint16:
		result = int32(val.(uint16))
	case int32:
		result = int32(val.(int32))
	case uint32:
		result = int32(val.(uint32))
	case int64:
		result = int32(val.(int64))
	case uint64:
		result = int32(val.(uint64))
	case float32:
		result = int32(val.(float32))
	case float64:
		result = int32(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = int32(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为int32列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Int32Array(valArray []interface{}) (result []int32, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]int32, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Int32(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为int64(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func Int64(val interface{}) (result int64, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = int64(val.(int))
	case uint:
		result = int64(val.(uint))
	case int8:
		result = int64(val.(int8))
	case uint8:
		result = int64(val.(uint8))
	case int16:
		result = int64(val.(int16))
	case uint16:
		result = int64(val.(uint16))
	case int32:
		result = int64(val.(int32))
	case uint32:
		result = int64(val.(uint32))
	case int64:
		result = int64(val.(int64))
	case uint64:
		result = int64(val.(uint64))
	case float32:
		result = int64(val.(float32))
	case float64:
		result = int64(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = int64(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为int64列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Int64Array(valArray []interface{}) (result []int64, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]int64, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Int64(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为uint(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Uint(val interface{}) (result uint, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = uint(val.(int))
	case uint:
		result = uint(val.(uint))
	case int8:
		result = uint(val.(int8))
	case uint8:
		result = uint(val.(uint8))
	case int16:
		result = uint(val.(int16))
	case uint16:
		result = uint(val.(uint16))
	case int32:
		result = uint(val.(int32))
	case uint32:
		result = uint(val.(uint32))
	case int64:
		result = uint(val.(int64))
	case uint64:
		result = uint(val.(uint64))
	case float32:
		result = uint(val.(float32))
	case float64:
		result = uint(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = uint(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为uint列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func UintArray(valArray []interface{}) (result []uint, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]uint, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Uint(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// uint8数据类型转换(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Uint8(val interface{}) (result uint8, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = uint8(val.(int))
	case uint:
		result = uint8(val.(uint))
	case int8:
		result = uint8(val.(int8))
	case uint8:
		result = uint8(val.(uint8))
	case int16:
		result = uint8(val.(int16))
	case uint16:
		result = uint8(val.(uint16))
	case int32:
		result = uint8(val.(int32))
	case uint32:
		result = uint8(val.(uint32))
	case int64:
		result = uint8(val.(int64))
	case uint64:
		result = uint8(val.(uint64))
	case float32:
		result = uint8(val.(float32))
	case float64:
		result = uint8(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = uint8(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// uint8数据类型转换(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Uint8Array(valArray []interface{}) (result []uint8, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]uint8, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Uint8(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为uint16(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Uint16(val interface{}) (result uint16, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = uint16(val.(int))
	case uint:
		result = uint16(val.(uint))
	case int8:
		result = uint16(val.(int8))
	case uint8:
		result = uint16(val.(uint8))
	case int16:
		result = uint16(val.(int16))
	case uint16:
		result = uint16(val.(uint16))
	case int32:
		result = uint16(val.(int32))
	case uint32:
		result = uint16(val.(uint32))
	case int64:
		result = uint16(val.(int64))
	case uint64:
		result = uint16(val.(uint64))
	case float32:
		result = uint16(val.(float32))
	case float64:
		result = uint16(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = uint16(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为uint16列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Uint16Array(valArray []interface{}) (result []uint16, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]uint16, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Uint16(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为uint32(转换过程不是类型安全的)
// val:待转换的值
// 返回值:
// result:结果
// err:错误数据
func Uint32(val interface{}) (result uint32, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = uint32(val.(int))
	case uint:
		result = uint32(val.(uint))
	case int8:
		result = uint32(val.(int8))
	case uint8:
		result = uint32(val.(uint8))
	case int16:
		result = uint32(val.(int16))
	case uint16:
		result = uint32(val.(uint16))
	case int32:
		result = uint32(val.(int32))
	case uint32:
		result = uint32(val.(uint32))
	case int64:
		result = uint32(val.(int64))
	case uint64:
		result = uint32(val.(uint64))
	case float32:
		result = uint32(val.(float32))
	case float64:
		result = uint32(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = uint32(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为uint32列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Uint32Array(valArray []interface{}) (result []uint32, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]uint32, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Uint32(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为uint64(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func Uint64(val interface{}) (result uint64, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = uint64(val.(int))
	case uint:
		result = uint64(val.(uint))
	case int8:
		result = uint64(val.(int8))
	case uint8:
		result = uint64(val.(uint8))
	case int16:
		result = uint64(val.(int16))
	case uint16:
		result = uint64(val.(uint16))
	case int32:
		result = uint64(val.(int32))
	case uint32:
		result = uint64(val.(uint32))
	case int64:
		result = uint64(val.(int64))
	case uint64:
		result = uint64(val.(uint64))
	case float32:
		result = uint64(val.(float32))
	case float64:
		result = uint64(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = uint64(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为uint64列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Uint64Array(valArray []interface{}) (result []uint64, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]uint64, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Uint64(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为float32(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func Float32(val interface{}) (result float32, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = float32(val.(int))
	case uint:
		result = float32(val.(uint))
	case int8:
		result = float32(val.(int8))
	case uint8:
		result = float32(val.(uint8))
	case int16:
		result = float32(val.(int16))
	case uint16:
		result = float32(val.(uint16))
	case int32:
		result = float32(val.(int32))
	case uint32:
		result = float32(val.(uint32))
	case int64:
		result = float32(val.(int64))
	case uint64:
		result = float32(val.(uint64))
	case float32:
		result = float32(val.(float32))
	case float64:
		result = float32(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = float32(tmp)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为float32列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Float32Array(valArray []interface{}) (result []float32, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]float32, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Float32(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为float64(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func Float64(val interface{}) (result float64, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = float64(val.(int))
	case uint:
		result = float64(val.(uint))
	case int8:
		result = float64(val.(int8))
	case uint8:
		result = float64(val.(uint8))
	case int16:
		result = float64(val.(int16))
	case uint16:
		result = float64(val.(uint16))
	case int32:
		result = float64(val.(int32))
	case uint32:
		result = float64(val.(uint32))
	case int64:
		result = float64(val.(int64))
	case uint64:
		result = float64(val.(uint64))
	case float32:
		result = float64(val.(float32))
	case float64:
		result = float64(val.(float64))
	case string:
		tmp, err1 := strconv.ParseFloat(val.(string), 64)
		if err1 != nil {
			err = fmt.Errorf("string convert error")
			return
		}

		result = tmp
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为Int列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func Float64Array(valArray []interface{}) (result []float64, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]float64, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Float64(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为bool(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func Bool(val interface{}) (result bool, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = (val.(int)) > 0
	case uint:
		result = (val.(uint)) > 0
	case int8:
		result = (val.(int8)) > 0
	case uint8:
		result = (val.(uint8)) > 0
	case int16:
		result = (val.(int16)) > 0
	case uint16:
		result = (val.(uint16)) > 0
	case int32:
		result = (val.(int32)) > 0
	case uint32:
		result = (val.(uint32)) > 0
	case int64:
		result = (val.(int64)) > 0
	case uint64:
		result = (val.(uint64)) > 0
	case float32:
		result = int(val.(float32)) > 0
	case float64:
		result = int(val.(float64)) > 0
	case bool:
		result = val.(bool)
	case string:
		tmp1, err1 := strconv.ParseBool(val.(string))
		if err1 != nil {
			// 先尝试转换成数值，再进行bool转换
			tmp2, err2 := strconv.ParseFloat(val.(string), 64)
			if err2 != nil {
				err = fmt.Errorf("string convert error")
				return
			}

			result = int(tmp2) > 0
			break
		}
		result = tmp1
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为Int列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func BoolArray(valArray []interface{}) (result []bool, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]bool, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := Bool(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换为字符串(转换过程不是类型安全的)
// 返回值:
// result:结果
// err:错误数据
func String(val interface{}) (result string, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case int:
		result = strconv.FormatInt(int64(val.(int)), 10)
	case uint:
		result = strconv.FormatUint(uint64(val.(uint)), 10)
	case int8:
		result = strconv.FormatInt(int64(val.(int8)), 10)
	case uint8:
		result = strconv.FormatUint(uint64(val.(uint8)), 10)
	case int16:
		result = strconv.FormatInt(int64(val.(int16)), 10)
	case uint16:
		result = strconv.FormatUint(uint64(val.(uint16)), 10)
	case int32:
		result = strconv.FormatInt(int64(val.(int32)), 10)
	case uint32:
		result = strconv.FormatUint(uint64(val.(uint32)), 10)
	case int64:
		result = strconv.FormatInt(int64(val.(int64)), 10)
	case uint64:
		result = strconv.FormatUint(uint64(val.(uint64)), 10)
	case float32:
		result = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case float64:
		result = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case string:
		result = val.(string)
	default:
		err = fmt.Errorf("val is not base type")
	}

	return
}

// 转换为Int列表(转换过程不是类型安全的)
// valArray:待转换的数据列表
// 返回值:
// result:结果
// err:错误数据
func StringArray(valArray []interface{}) (result []string, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]string, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := String(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 转换成时间格式
// val:待转换的数据,如果是字符串，则要求是格式:2006-01-02 15:04:05
// result:结果
// err:错误数据
func DateTime(val interface{}) (result time.Time, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case time.Time:
		result = val.(time.Time)
	case string:
		result, err = timeUtil.ToDateTime(val.(string))
	case int, int64, float32, float64:
		intVal, err1 := Int64(val)
		if err1 != nil {
			err = err1
			return
		}
		result = time.Unix(intVal, 0).Local()
	default:
		err = fmt.Errorf("unknown data type")
	}

	return
}

// 转换成时间格式
// valArray:待转换的数据,如果是字符串，则要求是格式:2006-01-02 15:04:05
// result:结果
// err:错误数据
func DateTimeArray(valArray []interface{}) (result []time.Time, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]time.Time, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := DateTime(item)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 转换成时间格式
// val:待转换的数据,如果是字符串，则使用format进行转换
// format:时间格式
// result:结果
// err:错误数据
func DateTimeByFormat(val interface{}, format string) (result time.Time, err error) {
	if val == nil {
		err = fmt.Errorf("val is nil")
		return
	}

	switch val.(type) {
	case time.Time:
		result = val.(time.Time)
	case string:
		result, err = time.ParseInLocation(val.(string), format, time.Local)
	case int, int64, float32, float64:
		intVal, err1 := Int64(val)
		if err1 != nil {
			err = err1
			return
		}
		result = time.Unix(intVal, 0).Local()
	default:
		err = fmt.Errorf("unknown data type")
	}

	return
}

// 转换成时间格式
// valArray:待转换的数据,如果是字符串，则使用format进行转换
// format:时间格式
// result:结果
// err:错误数据
func DateTimeArrayByFormat(valArray []interface{}, format string) (result []time.Time, err error) {
	if valArray == nil {
		err = fmt.Errorf("valArray is nil")
		return
	}

	result = make([]time.Time, 0, len(valArray))
	for _, item := range valArray {
		tmp, err1 := DateTimeByFormat(item, format)
		if err1 != nil {
			err = err1
			return
		}

		result = append(result, tmp)
	}

	return
}

// 类型转换（基础数据类型）
// val:原始值
// targetType:目标值类型
// 返回值:
// interface{}:结果
// error:错误信息
func Convert(val interface{}, targetType reflect.Kind) (result interface{}, err error) {
	switch targetType {
	case reflect.Int:
		result, err = Int(val)
	case reflect.Int8:
		result, err = Int8(val)
	case reflect.Int16:
		result, err = Int16(val)
	case reflect.Int32:
		result, err = Int32(val)
	case reflect.Int64:
		result, err = Int64(val)
	case reflect.Uint:
		result, err = Uint(val)
	case reflect.Uint8:
		result, err = Uint8(val)
	case reflect.Uint16:
		result, err = Uint16(val)
	case reflect.Uint32:
		result, err = Uint32(val)
	case reflect.Uint64:
		result, err = Uint64(val)
	case reflect.Float32:
		result, err = Float32(val)
	case reflect.Float64:
		result, err = Float64(val)
	case reflect.Bool:
		result, err = Bool(val)
	case reflect.String:
		result, err = String(val)
	default:
		err = fmt.Errorf("Unknown DataType:%s", targetType.String())
	}

	return
}
