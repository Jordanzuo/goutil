package mathUtil

import (
	"errors"
	"math/rand"
	"time"
)

// 获得Rand对象
func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 获取指定区间的随机数[lower, upper)
// lower:区间下限
// upper:区间上限
// 返回值：随机数
func GetRandRangeInt(lower, upper int) int {
	return lower + getRand().Intn(upper-lower)
}

// 获取随机数[0, n)
// n:范围上限
// 返回值：随机数
func GetRandInt(n int) int {
	return getRand().Intn(n)
}

// 获取随机数[0, n)
// n:范围上限
// 返回值：随机数
func GetRandInt32(n int32) int32 {
	return getRand().Int31n(n)
}

// 获取随机数[0, n)
// n:范围上限
// 返回值：随机数
func GetRandInt64(n int64) int64 {
	return getRand().Int63n(n)
}

// 获取随机数[0, 1)
// 返回值：随机数
func GetRandFloat32() float32 {
	return getRand().Float32()
}

// 获取随机数[0, 1)
// 返回值：随机数
func GetRandFloat64() float64 {
	return getRand().Float64()
}

// 获取随机数列表（1~10000，超过10000会抛出异常）
// minValue:获取随机数的区间下限值
// maxValue:获取随机数的区间上限值
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func GetRandNumList(minValue, maxValue, count int, ifAllowDuplicate bool) ([]int, error) {
	if minValue > maxValue {
		return nil, errors.New("minValue can't be bigger than maxValue.")
	}

	if !ifAllowDuplicate && (maxValue-minValue+1) < count {
		return nil, errors.New("随机的数量超过区间的元素数量")
	}

	if (maxValue - minValue + 1) > 10000 {
		return nil, errors.New("随机数的区间不能大于10000")
	}

	// 定义返回值
	resultList := make([]int, 0, count)
	var err error

	// 如果允许重复，则直接随机；否则调用GetRandList来随机
	if ifAllowDuplicate {
		for {
			if len(resultList) < count {
				resultList = append(resultList, GetRandRangeInt(minValue, maxValue+1))
			} else {
				break
			}
		}
	} else {
		inSlice := make([]int, maxValue-minValue+1, maxValue-minValue+1)
		for index := 0; index < len(inSlice); index++ {
			inSlice[index] = minValue + index
		}

		resultList, err = GetRandIntList(inSlice, count, ifAllowDuplicate)
	}

	return resultList, err
}

// 获取随机的int列表
// source:源列表
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func GetRandIntList(source []int, count int, ifAllowDuplicate bool) ([]int, error) {
	// 在不允许重复的情况下，需要产生的随机数不能超过范围限制
	if ifAllowDuplicate == false && len(source) < count {
		return nil, errors.New("随机的数量超过列表的元素数量")
	}

	// 使用源列表的数据量来初始化一个仅存放索引值的数组
	indexList := make([]int, count, count)
	for index := 0; index < count; index++ {
		indexList[index] = index
	}

	// 定义返回值
	resultList := make([]int, 0, count)

	// 遍历列表并获取随机对象(通过不断缩小随机的范围来实现)
	maxIndex := len(indexList) - 1
	for {
		if len(resultList) < count {
			// 获取随机索引(由于Next方法不取上限值，所以需要maxIndex+1)
			randIndex := GetRandRangeInt(0, maxIndex+1)

			// 将数据添加到列表，并增加findCount
			resultList = append(resultList, source[indexList[randIndex]])

			// 如果不允许重复，则需要特殊处理
			if !ifAllowDuplicate {
				// C#版本
				// // 并将该位置的数据设置为当前遍历的最大值
				// indexList[randIndex] = indexList[maxIndex]

				// Go版本
				// 将该位置的数据和最大位置的数据进行交换
				indexList[randIndex], indexList[maxIndex] = indexList[maxIndex], indexList[randIndex]

				// 将随机的范围缩小
				maxIndex = maxIndex - 1
			}
		} else {
			break
		}
	}

	return resultList, nil
}

// 获取随机的interface{}列表
// source:源列表
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func GetRandInterfaceList(source []interface{}, count int, ifAllowDuplicate bool) ([]interface{}, error) {
	// 在不允许重复的情况下，需要产生的随机数不能超过范围限制
	if ifAllowDuplicate == false && len(source) < count {
		return nil, errors.New("随机的数量超过列表的元素数量")
	}

	// 使用源列表的数据量来初始化一个仅存放索引值的数组
	indexList := make([]int, count, count)
	for index := 0; index < count; index++ {
		indexList[index] = index
	}

	// 定义返回值
	resultList := make([]interface{}, 0, count)

	// 遍历列表并获取随机对象(通过不断缩小随机的范围来实现)
	maxIndex := len(indexList) - 1
	for {
		if len(resultList) < count {
			// 获取随机索引(由于Next方法不取上限值，所以需要maxIndex+1)
			randIndex := GetRandRangeInt(0, maxIndex+1)

			// 将数据添加到列表，并增加findCount
			resultList = append(resultList, source[indexList[randIndex]])

			// 如果不允许重复，则需要特殊处理
			if !ifAllowDuplicate {
				// C#版本
				// // 并将该位置的数据设置为当前遍历的最大值
				// indexList[randIndex] = indexList[maxIndex]

				// Go版本
				// 将该位置的数据和最大位置的数据进行交换
				indexList[randIndex], indexList[maxIndex] = indexList[maxIndex], indexList[randIndex]

				// 将随机的范围缩小
				maxIndex = maxIndex - 1
			}
		} else {
			break
		}
	}

	return resultList, nil
}
