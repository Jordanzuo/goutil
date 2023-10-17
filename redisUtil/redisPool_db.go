/*
未实现的哈希表方法：
MOVE、SCAN、SORT、FLUSHDB、FLUSHALL、SELECT、SWAPDB
*/
package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

/*
// Do sends a command to the server and returns the received reply.
// This function will use the timeout which was set when the connection is created
*/
func (this *RedisPool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err = conn.Do(commandName, args...)
	return
}

/*
EXISTS key
可用版本： >= 1.0.0
时间复杂度： O(1)
检查给定 key 是否存在。

返回值
若 key 存在，返回 1 ，否则返回 0 。
*/
func (this *RedisPool) Exists(key string) (exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	result, err = redis.Int(conn.Do("EXISTS", key))
	if err != nil {
		return
	}

	if result == 1 {
		exist = true
	}

	return
}

/*
TYPE key
可用版本： >= 1.0.0
时间复杂度： O(1)
返回 key 所储存的值的类型。

返回值
none (key不存在)

string (字符串)

list (列表)

set (集合)

zset (有序集)

hash (哈希表)

stream （流）
*/
func (this *RedisPool) Type(key string) (_type string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_type, err = redis.String(conn.Do("TYPE", key))
	return
}

/*
RENAME key newkey

将 key 改名为 newkey 。

当 key 和 newkey 相同，或者 key 不存在时，返回一个错误。

当 newkey 已经存在时， RENAME 命令将覆盖旧值。

可用版本：
>= 1.0.0
时间复杂度：
O(1)
返回值：
改名成功时提示 OK ，失败时候返回一个错误。
*/
func (this *RedisPool) Rename(key, newkey string) (err error) {
	conn := this.GetConnection()
	defer conn.Close()

	_, err = conn.Do("RENAME", key, newkey)
	return
}

/*
RENAMENX key newkey

当且仅当 newkey 不存在时，将 key 改名为 newkey 。

当 key 不存在时，返回一个错误。

可用版本：
>= 1.0.0
时间复杂度：
O(1)
返回值：
修改成功时，返回 1 。
如果 newkey 已经存在，返回 0 。
*/
func (this *RedisPool) RenameNX(key, newkey string) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	result, err = redis.Int(conn.Do("RENAMENX", key, newkey))
	if err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
DEL key [key ...]

删除给定的一个或多个 key 。

不存在的 key 会被忽略。

可用版本：
>= 1.0.0
时间复杂度：
O(N)， N 为被删除的 key 的数量。
删除单个字符串类型的 key ，时间复杂度为O(1)。
删除单个列表、集合、有序集合或哈希表类型的 key ，时间复杂度为O(M)， M 为以上数据结构内的元素数量。
返回值：
被删除 key 的数量。
*/
func (this *RedisPool) Del(keys ...string) (count int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	count, err = redis.Int(conn.Do("DEL", redis.Args{}.AddFlat(keys)...))
	return
}

/*
从当前数据库中随机返回(不删除)一个 key 。

可用版本：
>= 1.0.0
时间复杂度：
O(1)
返回值：
当数据库不为空时，返回一个 key 。
当数据库为空时，返回 nil 。
*/
func (this *RedisPool) RandomKey() (key string, exist bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var reply interface{}
	reply, err = conn.Do("RANDOMKEY")
	if err != nil {
		return
	}
	if reply == nil {
		return
	}

	key, err = redis.String(reply, err)
	if err != nil {
		return
	}

	exist = true
	return
}

/*
DBSIZE
可用版本： >= 1.0.0
时间复杂度： O(1)
返回当前数据库的 key 的数量。

返回值
当前数据库的 key 的数量。
*/
func (this *RedisPool) DBSize() (keyCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	keyCount, err = redis.Int(conn.Do("DBSIZE"))
	return
}

/*
KEYS pattern

查找所有符合给定模式 pattern 的 key 。

KEYS * 匹配数据库中所有 key 。
KEYS h?llo 匹配 hello ， hallo 和 hxllo 等。
KEYS h*llo 匹配 hllo 和 heeeeello 等。
KEYS h[ae]llo 匹配 hello 和 hallo ，但不匹配 hillo 。
特殊符号用 \ 隔开

KEYS 的速度非常快，但在一个大的数据库中使用它仍然可能造成性能问题，如果你需要从一个数据集中查找特定的 key ，你最好还是用 Redis 的集合结构(set)来代替。
可用版本：
>= 1.0.0
时间复杂度：
O(N)， N 为数据库中 key 的数量。
返回值：
符合给定模式的 key 列表。
*/
func (this *RedisPool) Keys(pattern string) (keyList []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	keyList, err = redis.Strings(conn.Do("KEYS", pattern))
	return
}
