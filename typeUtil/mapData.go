package typeUtil

import (
	"fmt"
	"time"
)

// KeyValue数据集合
type MapData map[string]interface{}

// 创建新的MapData
// mapData:原有的map数据
// 返回
// 新的Map对象
func NewMapData(mapData map[string]interface{}) MapData {
	return MapData(mapData)
}

// 类型转换为byte
// 返回值:
// byte:结果
// error:错误数据
func (this MapData) Byte(key string) (byte, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Byte(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int32(key string) (int32, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Int32(val)
}

// 类型转换为uint32
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint32(key string) (uint32, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Uint32(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int(key string) (int, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Int(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint(key string) (uint, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Uint(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int64(key string) (int64, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Int64(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint64(key string) (uint64, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Uint64(val)
}

// 类型转换为int
// 返回值:
// float64:结果
// error:错误数据
func (this MapData) Float64(key string) (float64, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, fmt.Errorf("no exist target key")
	}

	return Float64(val)
}

// 类型转换为bool
// 返回值:
// bool:结果
// error:错误信息
func (this MapData) Bool(key string) (bool, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return false, fmt.Errorf("no exist target key")
	}

	return Bool(val)
}

// 类型转换为字符串
// 返回值:
// string:结果
// error:错误信息
func (this MapData) String(key string) (string, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return "", fmt.Errorf("no exist target key")
	}

	return String(val)
}

// 转换为时间格式，如果是字符串，则要求内容格式形如：2017-02-14 05:20:00
// 返回值:
// bool:结果
// error:错误信息
func (this MapData) DateTime(key string) (time.Time, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return time.Time{}, fmt.Errorf("no exist target key")
	}

	return DateTime(val)
}

// 获取指定的值
// 返回值:
// interface{}:结果
// error:错误信息
func (this MapData) Interface(key string) (interface{}, error) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return nil, fmt.Errorf("no exist target key")
	}

	return val, nil
}
