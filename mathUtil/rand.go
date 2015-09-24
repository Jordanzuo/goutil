package mathUtil

import (
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
