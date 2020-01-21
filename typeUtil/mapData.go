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
func (this MapData) Byte(key string) (value byte, err error) {
	return this.Uint8(key)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int(key string) (value int, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Int(val)
	return
}

// 类型转换为int8
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int8(key string) (value int8, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Int8(val)
	return
}

// 类型转换为int16
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int16(key string) (value int16, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Int16(val)
	return
}

// 类型转换为int32
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int32(key string) (value int32, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Int32(val)
	return
}

// 类型转换为int64
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Int64(key string) (value int64, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Int64(val)
	return
}

// 类型转换为uint
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint(key string) (value uint, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Uint(val)
	return
}

// 类型转换为uint8
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint8(key string) (value uint8, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Uint8(val)
	return
}

// 类型转换为uint16
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint16(key string) (value uint16, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Uint16(val)
	return
}

// 类型转换为uint32
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint32(key string) (value uint32, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Uint32(val)
	return
}

// 类型转换为uint64
// 返回值:
// int:结果
// error:错误数据
func (this MapData) Uint64(key string) (value uint64, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Uint64(val)
	return
}

// 类型转换为float32
// 返回值:
// float64:结果
// error:错误数据
func (this MapData) Float32(key string) (value float32, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Float32(val)
	return
}

// 类型转换为float64
// 返回值:
// float64:结果
// error:错误数据
func (this MapData) Float64(key string) (value float64, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Float64(val)
	return
}

// 类型转换为bool
// 返回值:
// bool:结果
// error:错误信息
func (this MapData) Bool(key string) (value bool, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = Bool(val)
	return
}

// 类型转换为字符串
// 返回值:
// string:结果
// error:错误信息
func (this MapData) String(key string) (value string, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = String(val)
	return
}

// 转换为时间格式，如果是字符串，则要求内容格式形如：2017-02-14 05:20:00
// 返回值:
// bool:结果
// error:错误信息
func (this MapData) DateTime(key string) (value time.Time, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value, err = DateTime(val)
	return
}

// 获取指定的值
// 返回值:
// interface{}:结果
// error:错误信息
func (this MapData) Interface(key string) (value interface{}, err error) {
	val, exist := this[key]
	if exist == false || val == nil {
		err = fmt.Errorf("Target key: [%s] doesn't exist", key)
		return
	}

	value = val
	return
}
