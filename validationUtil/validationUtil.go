package validationUtil

import (
	"errors"

	"github.com/Jordanzuo/goutil/stringUtil"
)

const (
	MaxInt32 int = 0x7fffffff
	MinInt32 int = -0x7fffffff
)

// 检查数值范围
// errList:错误列表
// val:待检查的值
// min:最小值
// max:最大值
// msg:错误提示
// 返回值:
// bool:是否在范围内
func CheckIntRange(errList *([]error), val int, min int, max int, msg string) bool {
	if val < min || val > max {
		if errList != nil {
			*errList = append(*errList, errors.New(msg))
		}

		return true
	}

	return false
}

// 检查数值范围
// errList:错误列表
// val:待检查的值
// min:最小值
// max:最大值
// msg:错误提示
// 返回值:
// bool:是否在范围内
func CheckFloatRange(errList *([]error), val float64, min float64, max float64, msg string) bool {
	if val < min || val > max {
		if errList != nil {
			*errList = append(*errList, errors.New(msg))
		}

		return true
	}

	return false
}

// 检查字符串是否为空
// errList:错误列表
// val:待检查的值
// msg:错误提示
// 返回值:
// bool:是否在范围内
func Require(errList *([]error), val string, msg string) bool {
	if stringUtil.IsEmpty(val) {
		if errList != nil {
			*errList = append(*errList, errors.New(msg))
		}

		return true
	}

	return false
}

// 检查数据是否存在重复项(此函数效率较低，如有效率要求请不要调用此函数)
func IsDistinct(count int, isEqualFunc func(i, j int) bool) bool {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i == j {
				continue
			}

			if isEqualFunc(i, j) {
				return false
			}
		}
	}

	return true
}
