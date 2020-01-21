package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

/*
SET key value [EX seconds] [PX milliseconds] [NX|XX]
可用版本： >= 1.0.0
时间复杂度： O(1)
将字符串值 value 关联到 key 。

如果 key 已经持有其他值， SET 就覆写旧值， 无视类型。

当 SET 命令对一个带有生存时间（TTL）的键进行设置之后， 该键原有的 TTL 将被清除。

可选参数
从 Redis 2.6.12 版本开始， SET 命令的行为可以通过一系列参数来修改：

EX seconds ： 将键的过期时间设置为 seconds 秒。 执行 SET key value EX seconds 的效果等同于执行 SETEX key seconds value 。

PX milliseconds ： 将键的过期时间设置为 milliseconds 毫秒。 执行 SET key value PX milliseconds 的效果等同于执行 PSETEX key milliseconds value 。

NX ： 只在键不存在时， 才对键进行设置操作。 执行 SET key value NX 的效果等同于执行 SETNX key value 。

XX ： 只在键已经存在时， 才对键进行设置操作。

Note

因为 SET 命令可以通过参数来实现 SETNX 、 SETEX 以及 PSETEX 命令的效果， 所以 Redis 将来的版本可能会移除并废弃 SETNX 、 SETEX 和 PSETEX 这三个命令。

返回值
在 Redis 2.6.12 版本以前， SET 命令总是返回 OK 。

从 Redis 2.6.12 版本开始， SET 命令只在设置操作成功完成时才返回 OK ； 如果命令使用了 NX 或者 XX 选项， 但是因为条件没达到而造成设置操作未执行， 那么命令将返回空批量回复（NULL Bulk Reply）。
*/
/*
expireType: "EX"|"PX"|""(参照上面的说明)
expireTime: seconds|milliseconds|无(根据expireType的不同而不同，参照上面的说明)
setType: "NX"|"XX"|""(参照上面的说明)
*/
/*
返回值：
successful: 是否成功
err: 错误对象
*/
func (this *RedisPool) Set(key string, value interface{}, expireType string, expireTime int, setType string) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := redis.Args{}.Add(key).Add(value)
	if expireType != "" {
		args = args.Add(expireType).Add(expireTime)
	}
	if setType != "" {
		args = args.Add(setType)
	}

	reply, err := conn.Do("SET", args...)
	if err != nil {
		return
	}
	if reply == nil {
		return
	}

	successful = true
	return
}

/*
GET key
可用版本： >= 1.0.0
时间复杂度： O(1)
返回与键 key 相关联的字符串值。

返回值
如果键 key 不存在， 那么返回特殊值 nil ； 否则， 返回键 key 的值。

如果键 key 的值并非字符串类型， 那么返回一个错误， 因为 GET 命令只能用于字符串值。
*/
func (this *RedisPool) Get(key string) (reply interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do("GET", key)
	if err != nil {
		return
	}
	if reply == nil {
		return
	}

	exist = true
	return
}

/*
GETSET key value
可用版本： >= 1.0.0
时间复杂度： O(1)
将键 key 的值设为 value ， 并返回键 key 在被设置之前的旧值。

返回值
返回给定键 key 的旧值。

如果键 key 没有旧值， 也即是说， 键 key 在被设置之前并不存在， 那么命令返回 nil 。

当键 key 存在但不是字符串类型时， 命令返回一个错误。
*/
func (this *RedisPool) GetSet(key, value interface{}) (oldValue interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	oldValue, err = conn.Do("GetSet", key, value)
	if err != nil {
		return
	}
	if oldValue == nil {
		return
	}

	exist = true
	return
}

/*
STRLEN key
可用版本： >= 2.2.0
复杂度： O(1)
返回键 key 储存的字符串值的长度。

返回值
STRLEN 命令返回字符串值的长度。

当键 key 不存在时， 命令返回 0 。

当 key 储存的不是字符串值时， 返回一个错误。
*/
func (this *RedisPool) StrLen(key string) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("STRLEN", key))
	return
}

/*
APPEND key value
可用版本： >= 2.0.0
时间复杂度： 平摊O(1)
如果键 key 已经存在并且它的值是一个字符串， APPEND 命令将把 value 追加到键 key 现有值的末尾。

如果 key 不存在， APPEND 就简单地将键 key 的值设为 value ， 就像执行 SET key value 一样。

返回值
追加 value 之后， 键 key 的值的长度。
*/
func (this *RedisPool) Append(key, value string) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("APPEND", key, value))
	return
}

/*
SETRANGE key offset value
可用版本： >= 2.2.0
时间复杂度：对于长度较短的字符串，命令的平摊复杂度O(1)；对于长度较大的字符串，命令的复杂度为 O(M) ，其中 M 为 value 的长度。
从偏移量 offset 开始， 用 value 参数覆写(overwrite)键 key 储存的字符串值。

不存在的键 key 当作空白字符串处理。

SETRANGE 命令会确保字符串足够长以便将 value 设置到指定的偏移量上， 如果键 key 原来储存的字符串长度比偏移量小(比如字符串只有 5 个字符长，但你设置的 offset 是 10 )， 那么原字符和偏移量之间的空白将用零字节(zerobytes, "\x00" )进行填充。

因为 Redis 字符串的大小被限制在 512 兆(megabytes)以内， 所以用户能够使用的最大偏移量为 2^29-1(536870911) ， 如果你需要使用比这更大的空间， 请使用多个 key 。

Warning

当生成一个很长的字符串时， Redis 需要分配内存空间， 该操作有时候可能会造成服务器阻塞(block)。 在2010年出产的Macbook Pro上， 设置偏移量为 536870911(512MB 内存分配)将耗费约 300 毫秒， 设置偏移量为 134217728(128MB 内存分配)将耗费约 80 毫秒， 设置偏移量 33554432(32MB 内存分配)将耗费约 30 毫秒， 设置偏移量为 8388608(8MB 内存分配)将耗费约 8 毫秒。

返回值
SETRANGE 命令会返回被修改之后， 字符串值的长度。
*/
func (this *RedisPool) SetRange(key string, offset int, value string) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("SETRANGE", key, offset, value))
	return
}

/*
GETRANGE key start end
可用版本： >= 2.4.0
时间复杂度： O(N)，其中 N 为被返回的字符串的长度。
返回键 key 储存的字符串值的指定部分， 字符串的截取范围由 start 和 end 两个偏移量决定 (包括 start 和 end 在内)。

负数偏移量表示从字符串的末尾开始计数， -1 表示最后一个字符， -2 表示倒数第二个字符， 以此类推。

GETRANGE 通过保证子字符串的值域(range)不超过实际字符串的值域来处理超出范围的值域请求。

Note

GETRANGE 命令在 Redis 2.0 之前的版本里面被称为 SUBSTR 命令。

返回值
GETRANGE 命令会返回字符串值的指定部分。
*/
func (this *RedisPool) GetRange(key string, start, end int) (value string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	value, err = redis.String(conn.Do("GETRANGE", key, start, end))
	return
}

/*
INCR key
可用版本： >= 1.0.0
时间复杂度： O(1)
为键 key 储存的数字值加上一。

如果键 key 不存在， 那么它的值会先被初始化为 0 ， 然后再执行 INCR 命令。

如果键 key 储存的值不能被解释为数字， 那么 INCR 命令将返回一个错误。

本操作的值限制在 64 位(bit)有符号数字表示之内。

Note

INCR 命令是一个针对字符串的操作。 因为 Redis 并没有专用的整数类型， 所以键 key 储存的值在执行 INCR 命令时会被解释为十进制 64 位有符号整数。

返回值
INCR 命令会返回键 key 在执行加一操作之后的值。
*/
func (this *RedisPool) Incr(key string) (newValue int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Int64(conn.Do("INCR", key))
	return
}

/*
INCRBY key increment
可用版本： >= 1.0.0
时间复杂度： O(1)
为键 key 储存的数字值加上增量 increment 。

如果键 key 不存在， 那么键 key 的值会先被初始化为 0 ， 然后再执行 INCRBY 命令。

如果键 key 储存的值不能被解释为数字， 那么 INCRBY 命令将返回一个错误。

本操作的值限制在 64 位(bit)有符号数字表示之内。

关于递增(increment) / 递减(decrement)操作的更多信息， 请参见 INCR 命令的文档。

返回值
在加上增量 increment 之后， 键 key 当前的值。
*/
func (this *RedisPool) IncrBy(key string, increment int64) (newValue int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Int64(conn.Do("INCRBY", key, increment))
	return
}

/*
INCRBYFLOAT key increment
可用版本： >= 2.6.0
时间复杂度： O(1)
为键 key 储存的值加上浮点数增量 increment 。

如果键 key 不存在， 那么 INCRBYFLOAT 会先将键 key 的值设为 0 ， 然后再执行加法操作。

如果命令执行成功， 那么键 key 的值会被更新为执行加法计算之后的新值， 并且新值会以字符串的形式返回给调用者。

无论是键 key 的值还是增量 increment ， 都可以使用像 2.0e7 、 3e5 、 90e-2 那样的指数符号(exponential notation)来表示， 但是， 执行 INCRBYFLOAT 命令之后的值总是以同样的形式储存， 也即是， 它们总是由一个数字， 一个（可选的）小数点和一个任意长度的小数部分组成（比如 3.14 、 69.768 ，诸如此类)， 小数部分尾随的 0 会被移除， 如果可能的话， 命令还会将浮点数转换为整数（比如 3.0 会被保存成 3 ）。

此外， 无论加法计算所得的浮点数的实际精度有多长， INCRBYFLOAT 命令的计算结果最多只保留小数点的后十七位。

当以下任意一个条件发生时， 命令返回一个错误：

键 key 的值不是字符串类型(因为 Redis 中的数字和浮点数都以字符串的形式保存，所以它们都属于字符串类型）；

键 key 当前的值或者给定的增量 increment 不能被解释(parse)为双精度浮点数。

返回值
在加上增量 increment 之后， 键 key 的值。
*/
func (this *RedisPool) IncrByFloat(key string, increment float64) (newValue float64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Float64(conn.Do("INCRBYFLOAT", key, increment))
	return
}

/*
DECR key
可用版本： >= 1.0.0
时间复杂度： O(1)
为键 key 储存的数字值减去一。

如果键 key 不存在， 那么键 key 的值会先被初始化为 0 ， 然后再执行 DECR 操作。

如果键 key 储存的值不能被解释为数字， 那么 DECR 命令将返回一个错误。

本操作的值限制在 64 位(bit)有符号数字表示之内。

关于递增(increment) / 递减(decrement)操作的更多信息， 请参见 INCR 命令的文档。

返回值
DECR 命令会返回键 key 在执行减一操作之后的值。
*/
func (this *RedisPool) Decr(key string) (newValue int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Int64(conn.Do("DECR", key))
	return
}

/*
DECRBY key decrement
可用版本： >= 1.0.0
时间复杂度： O(1)
将键 key 储存的整数值减去减量 decrement 。

如果键 key 不存在， 那么键 key 的值会先被初始化为 0 ， 然后再执行 DECRBY 命令。

如果键 key 储存的值不能被解释为数字， 那么 DECRBY 命令将返回一个错误。

本操作的值限制在 64 位(bit)有符号数字表示之内。

关于更多递增(increment) / 递减(decrement)操作的更多信息， 请参见 INCR 命令的文档。

返回值
DECRBY 命令会返回键在执行减法操作之后的值。
*/
func (this *RedisPool) DecrBy(key string, decrement int64) (newValue int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Int64(conn.Do("DECRBY", key, decrement))
	return
}

/*
MSET key value [key value …]
可用版本： >= 1.0.1
时间复杂度： O(N)，其中 N 为被设置的键数量。
同时为多个键设置值。

如果某个给定键已经存在， 那么 MSET 将使用新值去覆盖旧值， 如果这不是你所希望的效果， 请考虑使用 MSETNX 命令， 这个命令只会在所有给定键都不存在的情况下进行设置。

MSET 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。

返回值
MSET 命令总是返回 OK 。
*/
func (this *RedisPool) MSet(key_value_map map[string]interface{}) (err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_, err = conn.Do("MSET", redis.Args{}.AddFlat(key_value_map)...)
	return
}

/*
MSETNX key value [key value …]
可用版本： >= 1.0.1
时间复杂度： O(N)， 其中 N 为被设置的键数量。
当且仅当所有给定键都不存在时， 为所有给定键设置值。

即使只有一个给定键已经存在， MSETNX 命令也会拒绝执行对所有键的设置操作。

MSETNX 是一个原子性(atomic)操作， 所有给定键要么就全部都被设置， 要么就全部都不设置， 不可能出现第三种状态。

返回值
当所有给定键都设置成功时， 命令返回 1 ； 如果因为某个给定键已经存在而导致设置未能成功执行， 那么命令返回 0 。
*/
func (this *RedisPool) MSetNX(key_value_map map[string]interface{}) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var count int
	count, err = redis.Int(conn.Do("MSETNX", redis.Args{}.AddFlat(key_value_map)...))
	if err != nil {
		return
	}
	if count == 1 {
		successful = true
	}

	return
}

/*
MGET key [key …]
可用版本： >= 1.0.0
时间复杂度： O(N) ，其中 N 为给定键的数量。
返回给定的一个或多个字符串键的值。

如果给定的字符串键里面， 有某个键不存在， 那么这个键的值将以特殊值 nil 表示。

返回值
MGET 命令将返回一个列表， 列表中包含了所有给定键的值。
*/
func (this *RedisPool) MGet(keyList []string) (reply interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	// valueList, err = redis.Values(conn.Do("MGET", redis.Args{}.AddFlat(keyList)...))
	reply, err = conn.Do("MGET", redis.Args{}.AddFlat(keyList)...)
	return
}
