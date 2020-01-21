/*
未实现的列表方法：
BLPOP、BRPOP、BRPOPLPUSH
*/
package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

/*
LPUSH key value [value …]
可用版本： >= 1.0.0
时间复杂度： O(1)
将一个或多个值 value 插入到列表 key 的表头

如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表头： 比如说，对空列表 mylist 执行命令 LPUSH mylist a b c ，列表的值将是 c b a ，这等同于原子性地执行 LPUSH mylist a 、 LPUSH mylist b 和 LPUSH mylist c 三个命令。

如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。

当 key 存在但不是列表类型时，返回一个错误。

Note

在Redis 2.4版本以前的 LPUSH 命令，都只接受单个 value 值。

返回值
执行 LPUSH 命令后，列表的长度。
*/
func (this *RedisPool) LPush(key string, values ...interface{}) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(values)...))
	return
}

/*
LPUSHX key value
可用版本： >= 2.2.0
时间复杂度： O(1)
将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。

和 LPUSH key value [value …] 命令相反，当 key 不存在时， LPUSHX 命令什么也不做。

返回值
LPUSHX 命令执行之后，表的长度。
*/
func (this *RedisPool) LPushX(key string, values ...interface{}) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("LPUSHX", redis.Args{}.Add(key).AddFlat(values)...))
	return
}

/*
RPUSH key value [value …]
可用版本： >= 1.0.0
时间复杂度： O(1)
将一个或多个值 value 插入到列表 key 的表尾(最右边)。

如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表尾：比如对一个空列表 mylist 执行 RPUSH mylist a b c ，得出的结果列表为 a b c ，等同于执行命令 RPUSH mylist a 、 RPUSH mylist b 、 RPUSH mylist c 。

如果 key 不存在，一个空列表会被创建并执行 RPUSH 操作。

当 key 存在但不是列表类型时，返回一个错误。

Note

在 Redis 2.4 版本以前的 RPUSH 命令，都只接受单个 value 值。

返回值
执行 RPUSH 操作后，表的长度。
*/
func (this *RedisPool) RPush(key string, values ...interface{}) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(values)...))
	return
}

/*
RPUSHX key value
可用版本： >= 2.2.0
时间复杂度： O(1)
将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表。

和 RPUSH key value [value …] 命令相反，当 key 不存在时， RPUSHX 命令什么也不做。

返回值
RPUSHX 命令执行之后，表的长度。
*/
func (this *RedisPool) RPushX(key string, values ...interface{}) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("RPUSHX", redis.Args{}.Add(key).AddFlat(values)...))
	return
}

/*
LPOP key
可用版本： >= 1.0.0
时间复杂度： O(1)
移除并返回列表 key 的头元素。

返回值
列表的头元素。 当 key 不存在时，返回 nil 。
*/
func (this *RedisPool) LPop(key string) (item interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = conn.Do("LPOP", key)
	if err != nil {
		return
	}
	if item == nil {
		return
	}

	exist = true

	return
}

/*
RPOP key
可用版本： >= 1.0.0
时间复杂度： O(1)
移除并返回列表 key 的尾元素。

返回值
列表的尾元素。 当 key 不存在时，返回 nil 。
*/
func (this *RedisPool) RPop(key string) (item interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = conn.Do("RPOP", key)
	if err != nil {
		return
	}
	if item == nil {
		return
	}

	exist = true

	return
}

/*
RPOPLPUSH source destination
可用版本： >= 1.2.0
时间复杂度： O(1)
命令 RPOPLPUSH 在一个原子时间内，执行以下两个动作：

将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。

将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。

举个例子，你有两个列表 source 和 destination ， source 列表有元素 a, b, c ， destination 列表有元素 x, y, z ，执行 RPOPLPUSH source destination 之后， source 列表包含元素 a, b ， destination 列表包含元素 c, x, y, z ，并且元素 c 会被返回给客户端。

如果 source 不存在，值 nil 被返回，并且不执行其他动作。

如果 source 和 destination 相同，则列表中的表尾元素被移动到表头，并返回该元素，可以把这种特殊情况视作列表的旋转(rotation)操作。

返回值
被弹出的元素。
*/
func (this *RedisPool) RPopLPush(source, destination string) (item interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = conn.Do("RPOPLPUSH", source, destination)
	if err != nil {
		return
	}

	return
}

/*
LREM key count value
可用版本： >= 1.0.0
时间复杂度： O(N)， N 为列表的长度。
根据参数 count 的值，移除列表中与参数 value 相等的元素。

count 的值可以是以下几种：

count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。

count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。

count = 0 : 移除表中所有与 value 相等的值。

返回值
被移除元素的数量。 因为不存在的 key 被视作空表(empty list)，所以当 key 不存在时， LREM 命令总是返回 0 。
*/
// 错误对象
func (this *RedisPool) LRem(key string, count int, value string) (removeCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	removeCount, err = redis.Int(conn.Do("LREM", key, count, value))
	return
}

/*
LLEN key
可用版本： >= 1.0.0
时间复杂度： O(1)
返回列表 key 的长度。

如果 key 不存在，则 key 被解释为一个空列表，返回 0 .

如果 key 不是列表类型，返回一个错误。

返回值
列表 key 的长度。
*/
func (this *RedisPool) LLen(key string) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("LLEN", key))
	return
}

/*
LINDEX key index
可用版本： >= 1.0.0
时间复杂度：O(N)， N 为到达下标 index 过程中经过的元素数量。因此，对列表的头元素和尾元素执行 LINDEX 命令，复杂度为O(1)。
返回列表 key 中，下标为 index 的元素。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

如果 key 不是列表类型，返回一个错误。

返回值
列表中下标为 index 的元素。 如果 index 参数的值不在列表的区间范围内(out of range)，返回 nil 。
*/
func (this *RedisPool) LIndex(key string, index int) (item interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = conn.Do("LINDEX", key, index)
	if err != nil {
		return
	}
	if item == nil {
		return
	}

	exist = true
	return
}

/*
LINSERT key BEFORE|AFTER pivot value
可用版本： >= 2.2.0
时间复杂度: O(N)， N 为寻找 pivot 过程中经过的元素数量。
将值 value 插入到列表 key 当中，位于值 pivot 之前或之后。

当 pivot 不存在于列表 key 时，不执行任何操作。

当 key 不存在时， key 被视为空列表，不执行任何操作。

如果 key 不是列表类型，返回一个错误。

返回值
如果命令执行成功，返回插入操作完成之后，列表的长度。 如果没有找到 pivot ，返回 -1 。 如果 key 不存在或为空列表，返回 0 。
*/
func (this *RedisPool) LInsert(key, beforeOrAfter string, pivot, value interface{}) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("LINSERT", key, beforeOrAfter, pivot, value))
	return
}

/*
LSET key index value
可用版本： >= 1.0.0
时间复杂度：对头元素或尾元素进行 LSET 操作，复杂度为 O(1)。其他情况下，为 O(N)， N 为列表的长度。
将列表 key 下标为 index 的元素的值设置为 value 。

当 index 参数超出范围，或对一个空列表( key 不存在)进行 LSET 时，返回一个错误。

关于列表下标的更多信息，请参考 LINDEX key index 命令。

返回值
操作成功返回 ok ，否则返回错误信息。
*/
func (this *RedisPool) LSet(key string, index int, value interface{}) (err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_, err = conn.Do("LSet", key, index, value)
	return
}

/*
LRANGE key start stop
可用版本： >= 1.0.0
时间复杂度: O(S+N)， S 为偏移量 start ， N 为指定区间内元素的数量。
返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

注意LRANGE命令和编程语言区间函数的区别
假如你有一个包含一百个元素的列表，对该列表执行 LRANGE list 0 10 ，结果是一个包含11个元素的列表，这表明 stop 下标也在 LRANGE 命令的取值范围之内(闭区间)，这和某些语言的区间函数可能不一致，比如Ruby的 Range.new 、 Array#slice 和Python的 range() 函数。

超出范围的下标
超出范围的下标值不会引起错误。

如果 start 下标比列表的最大下标 end ( LLEN list 减去 1 )还要大，那么 LRANGE 返回一个空列表。

如果 stop 下标比 end 下标还要大，Redis将 stop 的值设置为 end 。

返回值
一个列表，包含指定区间内的元素。
*/
func (this *RedisPool) LRange(key string, start, stop int) (reply interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do("LRANGE", key, start, stop)
	return
}

/*
LTRIM key start stop
可用版本： >= 1.0.0
时间复杂度: O(N)， N 为被移除的元素的数量。
对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。

举个例子，执行命令 LTRIM list 0 2 ，表示只保留列表 list 的前三个元素，其余元素全部删除。

下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。

你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。

当 key 不是列表类型时，返回一个错误。

LTRIM 命令通常和 LPUSH key value [value …] 命令或 RPUSH key value [value …] 命令配合使用，举个例子：

LPUSH log newest_log
LTRIM log 0 99
这个例子模拟了一个日志程序，每次将最新日志 newest_log 放到 log 列表中，并且只保留最新的 100 项。注意当这样使用 LTRIM 命令时，时间复杂度是O(1)，因为平均情况下，每次只有一个元素被移除。

注意LTRIM命令和编程语言区间函数的区别
假如你有一个包含一百个元素的列表 list ，对该列表执行 LTRIM list 0 10 ，结果是一个包含11个元素的列表，这表明 stop 下标也在 LTRIM 命令的取值范围之内(闭区间)，这和某些语言的区间函数可能不一致，比如Ruby的 Range.new 、 Array#slice 和Python的 range() 函数。

超出范围的下标
超出范围的下标值不会引起错误。

如果 start 下标比列表的最大下标 end ( LLEN list 减去 1 )还要大，或者 start > stop ， LTRIM 返回一个空列表(因为 LTRIM 已经将整个列表清空)。

如果 stop 下标比 end 下标还要大，Redis将 stop 的值设置为 end 。

返回值
命令执行成功时，返回 ok 。
*/
func (this *RedisPool) LTrim(key string, start, stop int) (err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_, err = conn.Do("LTRIM", key, start, stop)
	return
}
