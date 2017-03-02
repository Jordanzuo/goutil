package dbUtil

import (
	"fmt"
	"time"

	"github.com/Jordanzuo/goutil/typeUtil"
)

// 类型转换为byte
// 返回值:
// byte:结果
// error:错误数据
func Byte(row *DataRow, key string) (byte, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Byte(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Int32(row *DataRow, key string) (int32, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Int32(val)
}

// 类型转换为uint32
// 返回值:
// int:结果
// error:错误数据
func Uint32(row *DataRow, key string) (uint32, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Uint32(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Int(row *DataRow, key string) (int, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Int(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Uint(row *DataRow, key string) (uint, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Uint(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Int64(row *DataRow, key string) (int64, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Int64(val)
}

// 类型转换为int
// 返回值:
// int:结果
// error:错误数据
func Uint64(row *DataRow, key string) (uint64, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Uint64(val)
}

// 类型转换为int
// 返回值:
// float64:结果
// error:错误数据
func Float64(row *DataRow, key string) (float64, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return 0, errMsg
	}

	if val == nil {
		return 0, fmt.Errorf("value is nil")
	}

	return typeUtil.Float64(val)
}

// 类型转换为bool
// 返回值:
// bool:结果
// error:错误信息
func Bool(row *DataRow, key string) (bool, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return false, errMsg
	}

	if val == nil {
		return false, fmt.Errorf("value is nil")
	}

	return typeUtil.Bool(val)
}

// 类型转换为字符串
// 返回值:
// string:结果
// error:错误信息
func String(row *DataRow, key string) (string, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return "", errMsg
	}

	if val == nil {
		return "", fmt.Errorf("value is nil")
	}

	return typeUtil.String(val)
}

// 转换为时间格式，如果是字符串，则要求内容格式形如：2017-02-14 05:20:00
// 返回值:
// bool:结果
// error:错误信息
func DateTime(row *DataRow, key string) (time.Time, error) {
	val, errMsg := row.CellValueByName(key)
	if errMsg != nil {
		return time.Time{}, errMsg
	}

	if val == nil {
		return time.Time{}, fmt.Errorf("value is nil")
	}

	return typeUtil.DateTime(val)
}
