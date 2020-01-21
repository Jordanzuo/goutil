package mathUtil

import (
	"errors"
	"math/rand"
	"time"
)

type Rand struct {
	*rand.Rand
}

// 获得Rand对象(如果是循环生成多个数据，则只需要调用本方法一次，而不能调用多次，否则会得到相同的随机值)
func GetRand() *Rand {
	return &Rand{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// 获取指定区间的随机数[lower, upper)
// lower:区间下限
// upper:区间上限
// 返回值：随机数
func (this *Rand) GetRandRangeInt(lower, upper int) int {
	return lower + this.Intn(upper-lower)
}

// 获取随机数[0, n)
// randObj:随机对象
// n:范围上限
// 返回值：随机数
func (this *Rand) GetRandInt(n int) int {
	return this.Intn(n)
}

// 获取随机数[0, n)
// n:范围上限
// 返回值：随机数
func (this *Rand) GetRandInt32(n int32) int32 {
	return this.Int31n(n)
}

// 获取随机数[0, n)
// randObj:随机对象
// n:范围上限
// 返回值：随机数
func (this *Rand) GetRandInt64(n int64) int64 {
	return this.Int63n(n)
}

// 获取随机数[0, 1)
// randObj:随机对象
// 返回值：随机数
func (this *Rand) GetRandFloat32() float32 {
	return this.Float32()
}

// 获取随机数[0, 1)
// 返回值：随机数
func (this *Rand) GetRandFloat64() float64 {
	return this.Float64()
}

// 获取随机数列表（1~10000，超过10000会抛出异常）
// minValue:获取随机数的区间下限值
// maxValue:获取随机数的区间上限值
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func (this *Rand) GetRandNumList(minValue, maxValue, count int, ifAllowDuplicate bool) ([]int, error) {
	if minValue > maxValue {
		return nil, errors.New("minValue can't be bigger than maxValue.")
	}

	if !ifAllowDuplicate && (maxValue-minValue+1) < count {
		return nil, errors.New("随机的数量超过区间的元素数量")
	}

	if (maxValue - minValue + 1) > 10000 {
		return nil, errors.New("随机数的区间不能大于10000")
	}

	// 定义原始数据
	sourceCount := maxValue - minValue + 1
	source := make([]int, sourceCount, sourceCount)
	for index := 0; index < sourceCount; index++ {
		source[index] = minValue + index
	}

	// 定义返回值
	resultList := make([]int, 0, count)

	// 获取随机的索引列表
	randIndextList := this.getRandIndexList(len(source), count, ifAllowDuplicate)
	for _, index := range randIndextList {
		// 判断是否已经取到足够数量的数据？
		if count <= 0 {
			break
		}

		resultList = append(resultList, source[index])
		count -= 1
	}

	return resultList, nil
}

// 获取随机的int列表
// source:源列表
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func (this *Rand) GetRandIntList(source []int, count int, ifAllowDuplicate bool) ([]int, error) {
	// 在不允许重复的情况下，需要产生的随机数不能超过范围限制
	if ifAllowDuplicate == false && len(source) < count {
		return nil, errors.New("随机的数量超过列表的元素数量")
	}

	// 定义返回值
	resultList := make([]int, 0, count)

	// 获取随机的索引列表
	randIndextList := this.getRandIndexList(len(source), count, ifAllowDuplicate)
	for _, index := range randIndextList {
		// 判断是否已经取到足够数量的数据？
		if count <= 0 {
			break
		}

		resultList = append(resultList, source[index])
		count -= 1
	}

	return resultList, nil
}

// 获取随机的interface{}列表
// source:源列表
// count:随机数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机数列表
func (this *Rand) GetRandInterfaceList(source []interface{}, count int, ifAllowDuplicate bool) ([]interface{}, error) {
	// 在不允许重复的情况下，需要产生的随机数不能超过范围限制
	if ifAllowDuplicate == false && len(source) < count {
		return nil, errors.New("随机的数量超过列表的元素数量")
	}

	// 定义返回值
	resultList := make([]interface{}, 0, count)

	// 获取随机的索引列表
	randIndextList := this.getRandIndexList(len(source), count, ifAllowDuplicate)
	for _, index := range randIndextList {
		// 判断是否已经取到足够数量的数据？
		if count <= 0 {
			break
		}

		resultList = append(resultList, source[index])
		count -= 1
	}

	return resultList, nil
}

// 获取随机的索引列表
// count:数量
// ifAllowDuplicate:是否允许重复
// 返回值
// 随机索引值列表
func (this *Rand) getRandIndexList(maxNum, count int, ifAllowDuplicate bool) []int {
	// 定义返回值
	randIndextList := make([]int, 0, count)

	// 使用源列表的数据量来初始化一个仅存放索引值的数组
	indexList := make([]int, maxNum, maxNum)
	for index := 0; index < maxNum; index++ {
		indexList[index] = index
	}

	// 遍历列表并获取随机对象(通过不断缩小随机的范围来实现)
	maxIndex := len(indexList) - 1
	for {
		if len(randIndextList) < count {
			// 获取随机索引(由于Next方法不取上限值，所以需要maxIndex+1)
			randIndex := this.Intn(maxIndex + 1)

			// 将数据添加到列表，并增加findCount
			randIndextList = append(randIndextList, indexList[randIndex])

			// 如果不允许重复，则需要特殊处理
			if !ifAllowDuplicate {
				// 将该位置的数据和最大位置的数据进行交换
				indexList[randIndex], indexList[maxIndex] = indexList[maxIndex], indexList[randIndex]

				// 将随机的范围缩小
				maxIndex -= 1
			}
		} else {
			break
		}
	}

	return randIndextList
}

// 获取带权重的随机数据
// source:源数据
// 返回值
// 数据项
// 错误对象
func (this *Rand) GetRandWeight(source []IWeight) (result IWeight, err error) {
	if source == nil || len(source) == 0 {
		err = errors.New("待随机的列表为空")
		return
	}

	// 计算出总的数据量，并随机一个[0, total)的值
	total := 0
	for _, item := range source {
		total += item.GetWeight()
	}

	randNum := this.GetRandInt(total)

	// 根据随机出来的值，判断位于哪个区间
	total = 0
	for _, item := range source {
		total += item.GetWeight()
		if randNum < total {
			result = item
			return
		}
	}

	err = errors.New("未找到有效的数据")
	return
}
