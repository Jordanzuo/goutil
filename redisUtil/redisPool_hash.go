/*
未实现的哈希表方法：
HSCAN
*/
package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

/*
HSET hash field value
可用版本： >= 2.0.0
时间复杂度： O(1)
将哈希表 hash 中域 field 的值设置为 value 。

如果给定的哈希表并不存在， 那么一个新的哈希表将被创建并执行 HSET 操作。

如果域 field 已经存在于哈希表中， 那么它的旧值将被新值 value 覆盖。

返回值
当 HSET 命令在哈希表中新创建 field 域并成功为它设置值时， 命令返回 1 ； 如果域 field 已经存在于哈希表， 并且 HSET 命令成功使用新值覆盖了它的旧值， 那么命令返回 0 。
*/
func (this *RedisPool) HSet(key, field string, value interface{}) (result int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	result, err = redis.Int(conn.Do("HSET", key, field, value))
	return
}

/*
HSETNX hash field value
可用版本： >= 2.0.0
时间复杂度： O(1)
当且仅当域 field 尚未存在于哈希表的情况下， 将它的值设置为 value 。

如果给定域已经存在于哈希表当中， 那么命令将放弃执行设置操作。

如果哈希表 hash 不存在， 那么一个新的哈希表将被创建并执行 HSETNX 命令。

返回值
HSETNX 命令在设置成功时返回 1 ， 在给定域已经存在而放弃执行设置操作时返回 0 。
*/
func (this *RedisPool) HSetNX(key, field string, value interface{}) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	result, err = redis.Int(conn.Do("HSETNX", key, field, value))
	if err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
HGET hash field
可用版本： >= 2.0.0
时间复杂度： O(1)
返回哈希表中给定域的值。

返回值
HGET 命令在默认情况下返回给定域的值。

如果给定域不存在于哈希表中， 又或者给定的哈希表并不存在， 那么命令返回 nil 。
*/
func (this *RedisPool) HGet(key, field string) (value interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	value, err = conn.Do("HGET", key, field)
	if err != nil {
		return
	}
	if value == nil {
		return
	}

	exist = true
	return
}

/*
HEXISTS hash field
可用版本： >= 2.0.0
时间复杂度： O(1)
检查给定域 field 是否存在于哈希表 hash 当中。

返回值
HEXISTS 命令在给定域存在时返回 1 ， 在给定域不存在时返回 0 。
*/
func (this *RedisPool) HExists(key, field string) (exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	result, err = redis.Int(conn.Do("HEXISTS", key, field))
	if err != nil {
		return
	}

	if result == 1 {
		exist = true
	}

	return
}

/*
HDEL
HDEL key field [field …]

删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。

Note

在Redis2.4以下的版本里， HDEL 每次只能删除单个域，如果你需要在一个原子时间内删除多个域，请将命令包含在 MULTI / EXEC 块内。

可用版本：
>= 2.0.0

时间复杂度:
O(N)， N 为要删除的域的数量。

返回值:
被成功移除的域的数量，不包括被忽略的域。
*/
func (this *RedisPool) HDel(key string, field ...string) (succeedCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	succeedCount, err = redis.Int(conn.Do("HDEL", redis.Args{}.Add(key).AddFlat(field)...))
	return
}

/*
HLEN
HLEN key

返回哈希表 key 中域的数量。

时间复杂度：
O(1)

返回值：
哈希表中域的数量。
当 key 不存在时，返回 0 。
*/
func (this *RedisPool) HLen(key string) (count int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	count, err = redis.Int(conn.Do("HLEN", key))
	return
}

/*
HSTRLEN
HSTRLEN key field

返回哈希表 key 中， 与给定域 field 相关联的值的字符串长度（string length）。

如果给定的键或者域不存在， 那么命令返回 0 。

可用版本：
>= 3.2.0

时间复杂度：
O(1)

返回值：
一个整数。
*/
func (this *RedisPool) HStrlen(key, field string) (length int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	length, err = redis.Int(conn.Do("HSTRLEN", key, field))
	return
}

/*
HINCRBY
HINCRBY key field increment

为哈希表 key 中的域 field 的值加上增量 increment 。

增量也可以为负数，相当于对给定域进行减法操作。

如果 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令。

如果域 field 不存在，那么在执行命令前，域的值被初始化为 0 。

对一个储存字符串值的域 field 执行 HINCRBY 命令将造成一个错误。

本操作的值被限制在 64 位(bit)有符号数字表示之内。

可用版本：
>= 2.0.0

时间复杂度：
O(1)

返回值：
执行 HINCRBY 命令之后，哈希表 key 中域 field 的值。
*/
func (this *RedisPool) HIncrBy(key, field string, increment int64) (newValue int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Int64(conn.Do("HINCRBY", key, field, increment))
	return
}

/*
HINCRBYFLOAT
HINCRBYFLOAT key field increment

为哈希表 key 中的域 field 加上浮点数增量 increment 。

如果哈希表中没有域 field ，那么 HINCRBYFLOAT 会先将域 field 的值设为 0 ，然后再执行加法操作。

如果键 key 不存在，那么 HINCRBYFLOAT 会先创建一个哈希表，再创建域 field ，最后再执行加法操作。

当以下任意一个条件发生时，返回一个错误：

域 field 的值不是字符串类型(因为 redis 中的数字和浮点数都以字符串的形式保存，所以它们都属于字符串类型）

域 field 当前的值或给定的增量 increment 不能解释(parse)为双精度浮点数(double precision floating point number)

HINCRBYFLOAT 命令的详细功能和 INCRBYFLOAT key increment 命令类似，请查看 INCRBYFLOAT key increment 命令获取更多相关信息。

可用版本：
>= 2.6.0

时间复杂度：
O(1)

返回值：
执行加法操作之后 field 域的值。
*/
func (this *RedisPool) HIncrByFloat(key, field string, increment float64) (newValue float64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	newValue, err = redis.Float64(conn.Do("HINCRBYFLOAT", key, field, increment))
	return
}

/*
HMSET
HMSET key field value [field value …]

同时将多个 field-value (域-值)对设置到哈希表 key 中。

此命令会覆盖哈希表中已存在的域。

如果 key 不存在，一个空哈希表被创建并执行 HMSET 操作。

可用版本：
>= 2.0.0

时间复杂度：
O(N)， N 为 field-value 对的数量。

返回值：
如果命令执行成功，返回 OK 。
当 key 不是哈希表(hash)类型时，返回一个错误。
*/
func (this *RedisPool) HMSet(key string, value interface{}) (err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_, err = conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	return
}

/*
HMGET
HMGET key field [field …]

返回哈希表 key 中，一个或多个给定域的值。

如果给定的域不存在于哈希表，那么返回一个 nil 值。

因为不存在的 key 被当作一个空哈希表来处理，所以对一个不存在的 key 进行 HMGET 操作将返回一个只带有 nil 值的表。

可用版本：
>= 2.0.0

时间复杂度：
O(N)， N 为给定域的数量。

返回值：
一个包含多个给定域的关联值的表，表值的排列顺序和给定域参数的请求顺序一样。
*/
func (this *RedisPool) HMGet(key string, fields ...string) (reply interface{}, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...)
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
HKEYS
HKEYS key

返回哈希表 key 中的所有域。

可用版本：
>= 2.0.0

时间复杂度：
O(N)， N 为哈希表的大小。

返回值：
一个包含哈希表中所有域的表。
当 key 不存在时，返回一个空表。
*/
func (this *RedisPool) HKeys(key string) (fieldList []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	fieldList, err = redis.Strings(conn.Do("HKEYS", key))
	return
}

/*
VALS
HVALS key

返回哈希表 key 中所有域的值。

可用版本：
>= 2.0.0

时间复杂度：
O(N)， N 为哈希表的大小。

返回值：
一个包含哈希表中所有值的表。
当 key 不存在时，返回一个空表。
*/
func (this *RedisPool) HVals(key string) (reply interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do("HVALS", key)
	return
}

/*
HGETALL
HGETALL key

返回哈希表 key 中，所有的域和值。

在返回值里，紧跟每个域名(field name)之后是域的值(value)，所以返回值的长度是哈希表大小的两倍。

可用版本：
>= 2.0.0

时间复杂度：
O(N)， N 为哈希表的大小。

返回值：
以列表形式返回哈希表的域和域的值。
若 key 不存在，返回空列表。
*/
func (this *RedisPool) HGetAll(key string) (reply interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do("HGETALL", key)
	return
}

func (this *RedisPool) HGetAll_Struct(key string, obj interface{}) (exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return
	}
	if len(values) == 0 {
		return
	}

	err = redis.ScanStruct(values, obj)
	if err != nil {
		return
	}

	exist = true
	return
}
