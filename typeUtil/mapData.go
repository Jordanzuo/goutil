package typeUtil

import (
	"fmt"
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
