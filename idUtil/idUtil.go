/*
用于生成唯一的、递增的Id。生成的规则如下：
1、生成的Id包含一个固定前缀值
2、为了生成尽可能多的不重复数字，所以使用int64来表示一个数字，其中：
0 000000000000000 0000000000000000000000000000 00000000000000000000
第一部分：1位，固定为0
第二部分：共PrefixBitCount位，表示固定前缀值。范围为[0, math.Pow(2, PrefixBitCount))
第三部分：共TimeBitCount位，表示当前时间距离基础时间的秒数。范围为[0, math.Pow(2, TimeBitCount))，以2019-1-1 00:00:00为基准则可以持续到2025-07-01 00:00:00
第四部分：共SeedBitCount位，表示自增种子。范围为[0, math.Pow(2, SeedBitCount))
3、总体而言，此规则支持每秒生成math.Pow(2, SeedBitCount)个不同的数字，并且在math.Pow(2, TimeBitCount)/60/60/24/365年的时间范围内有效
*/

package idUtil

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Id生成器
type IdGenerator struct {
	prefixBitCount uint  // 前缀所占的位数
	minPrefix      int64 // 最小的前缀值
	maxPrefix      int64 // 最大的前缀值

	timeBitCount    uint  // 时间戳所占的位数
	baseTimeUnix    int64 // 基础时间
	maxTimeDuration int64 // 最大的时间范围

	seedBitCount uint  // 自增种子所占的位数
	currSeed     int64 // 当前种子值
	minSeed      int64 // 最小的种子值
	maxSeed      int64 // 最大的种子值

	mutex sync.Mutex // 锁对象
}

func (this *IdGenerator) getTimeStamp() int64 {
	return time.Now().Unix() - this.baseTimeUnix
}

func (this *IdGenerator) generateSeed() int64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.currSeed >= this.maxSeed {
		this.currSeed = this.minSeed
	} else {
		this.currSeed = this.currSeed + 1
	}

	return this.currSeed
}

// 生成新的Id
// prefix：Id的前缀值。取值范围必须可以用创建对象时指定的前缀值的位数来表示，否则会返回参数超出范围的错误
// 返回值：
// 新的Id
// 错误对象
func (this *IdGenerator) GenerateNewId(prefix int64) (int64, error) {
	if prefix < this.minPrefix || prefix > this.maxPrefix {
		return 0, fmt.Errorf("前缀值溢出，有效范围为【%d,%d】", this.minPrefix, this.maxPrefix)
	}

	stamp := this.getTimeStamp()
	seed := this.generateSeed()
	id := prefix<<(this.timeBitCount+this.seedBitCount) + stamp<<this.seedBitCount + seed

	return id, nil
}

// 创建新的Id生成器对象（为了保证Id的唯一，需要保证生成的对象全局唯一）
// prefixBitCount + timeBitCount + seedBitCount <= 63
// prefixBitCount:表示id前缀的位数
// timeBitCount:表示时间的位数
// seedBitCount:表示自增种子的位数
// 返回值：
// 新的Id生成器对象
// 错误对象
func New(prefixBitCount, timeBitCount, seedBitCount uint) (*IdGenerator, error) {
	// 之所以使用63位而不是64，是为了保证值为正数
	if prefixBitCount+timeBitCount+seedBitCount > 63 {
		return nil, fmt.Errorf("总位数%d超过63位，请调整所有值的合理范围。", prefixBitCount+timeBitCount+seedBitCount)
	}

	obj := &IdGenerator{
		prefixBitCount: prefixBitCount,
		timeBitCount:   timeBitCount,
		seedBitCount:   seedBitCount,
	}

	obj.minPrefix = 0
	obj.maxPrefix = int64(math.Pow(2, float64(prefixBitCount))) - 1
	obj.baseTimeUnix = time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local).Unix()
	obj.maxTimeDuration = (int64(math.Pow(2, float64(timeBitCount))) - 1) / 86400 / 365
	obj.currSeed = 0
	obj.minSeed = 0
	obj.maxSeed = int64(math.Pow(2, float64(seedBitCount))) - 1
	fmt.Printf("Prefix:[%d, %d], Time:%d Year, Seed:[%d, %d]\n", obj.minPrefix, obj.maxPrefix, obj.maxTimeDuration, obj.minSeed, obj.maxSeed)

	return obj, nil
}
