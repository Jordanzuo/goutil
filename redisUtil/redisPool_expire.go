/*
未实现的哈希表方法：
MOVE、SCAN、SORT、FLUSHDB、FLUSHALL、SELECT、SWAPDB
*/
package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

/*
EXPIRE key seconds
可用版本： >= 1.0.0
时间复杂度： O(1)
为给定 key 设置生存时间，当 key 过期时(生存时间为 0 )，它会被自动删除。

在 Redis 中，带有生存时间的 key 被称为『易失的』(volatile)。

生存时间可以通过使用 DEL 命令来删除整个 key 来移除，或者被 SET 和 GETSET 命令覆写(overwrite)，这意味着，如果一个命令只是修改(alter)一个带生存时间的 key 的值而不是用一个新的 key 值来代替(replace)它的话，那么生存时间不会被改变。

比如说，对一个 key 执行 INCR 命令，对一个列表进行 LPUSH 命令，或者对一个哈希表执行 HSET 命令，这类操作都不会修改 key 本身的生存时间。

另一方面，如果使用 RENAME 对一个 key 进行改名，那么改名后的 key 的生存时间和改名前一样。

RENAME 命令的另一种可能是，尝试将一个带生存时间的 key 改名成另一个带生存时间的 another_key ，这时旧的 another_key (以及它的生存时间)会被删除，然后旧的 key 会改名为 another_key ，因此，新的 another_key 的生存时间也和原本的 key 一样。

使用 PERSIST 命令可以在不删除 key 的情况下，移除 key 的生存时间，让 key 重新成为一个『持久的』(persistent) key 。

更新生存时间
可以对一个已经带有生存时间的 key 执行 EXPIRE 命令，新指定的生存时间会取代旧的生存时间。

过期时间的精确度
在 Redis 2.4 版本中，过期时间的延迟在 1 秒钟之内 —— 也即是，就算 key 已经过期，但它还是可能在过期之后一秒钟之内被访问到，而在新的 Redis 2.6 版本中，延迟被降低到 1 毫秒之内。

Redis 2.1.3 之前的不同之处
在 Redis 2.1.3 之前的版本中，修改一个带有生存时间的 key 会导致整个 key 被删除，这一行为是受当时复制(replication)层的限制而作出的，现在这一限制已经被修复。

返回值
设置成功返回 1 。 当 key 不存在或者不能为 key 设置生存时间时(比如在低于 2.1.3 版本的 Redis 中你尝试更新 key 的生存时间)，返回 0 。
*/
func (this *RedisPool) Expire(key string, seconds int64) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("EXPIRE", key, seconds)); err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
EXPIREAT key timestamp
可用版本： >= 1.2.0
时间复杂度： O(1)
EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置生存时间。

不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。

返回值
如果生存时间设置成功，返回 1 ； 当 key 不存在或没办法设置生存时间，返回 0 。
*/
func (this *RedisPool) ExpireAt(key string, timestamp int64) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("EXPIREAT", key, timestamp)); err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
TTL key
可用版本： >= 1.0.0
时间复杂度： O(1)
以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。

返回值
当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以秒为单位，返回 key 的剩余生存时间。

Note

在 Redis 2.8 以前，当 key 不存在，或者 key 没有设置剩余生存时间时，命令都返回 -1 。
*/
func (this *RedisPool) TTL(key string) (ttl int64, exist, persisted bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	ttl, err = redis.Int64(conn.Do("TTL", key))
	if err != nil {
		return
	}

	if ttl == -2 {
		exist = false
		persisted = false
	} else if ttl == -1 {
		exist = true
		persisted = true
	} else {
		exist = true
		persisted = false
	}

	return
}

/*
PERSIST key
可用版本： >= 2.2.0
时间复杂度： O(1)
移除给定 key 的生存时间，将这个 key 从“易失的”(带生存时间 key )转换成“持久的”(一个不带生存时间、永不过期的 key )。

返回值
当生存时间移除成功时，返回 1 . 如果 key 不存在或 key 没有设置生存时间，返回 0 。
*/
func (this *RedisPool) Persist(key string) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("PERSIST", key)); err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
PEXPIRE key milliseconds
可用版本： >= 2.6.0
时间复杂度： O(1)
这个命令和 EXPIRE 命令的作用类似，但是它以毫秒为单位设置 key 的生存时间，而不像 EXPIRE 命令那样，以秒为单位。

返回值
设置成功，返回 1 key 不存在或设置失败，返回 0
*/
func (this *RedisPool) PExpire(key string, milliseconds int64) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("PEXPIRE", key, milliseconds)); err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
PEXPIREAT key milliseconds-timestamp
可用版本： >= 2.6.0
时间复杂度： O(1)
这个命令和 expireat 命令类似，但它以毫秒为单位设置 key 的过期 unix 时间戳，而不是像 expireat 那样，以秒为单位。

返回值
如果生存时间设置成功，返回 1 。 当 key 不存在或没办法设置生存时间时，返回 0 。(查看 EXPIRE key seconds 命令获取更多信息)
*/
func (this *RedisPool) PExpireAt(key string, milliseconds_timestamp int64) (successful bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("PEXPIREAT", key, milliseconds_timestamp)); err != nil {
		return
	}

	if result == 1 {
		successful = true
	}

	return
}

/*
PTTL key
可用版本： >= 2.6.0
复杂度： O(1)
这个命令类似于 TTL 命令，但它以毫秒为单位返回 key 的剩余生存时间，而不是像 TTL 命令那样，以秒为单位。

返回值
当 key 不存在时，返回 -2 。

当 key 存在但没有设置剩余生存时间时，返回 -1 。

否则，以毫秒为单位，返回 key 的剩余生存时间。

Note

在 Redis 2.8 以前，当 key 不存在，或者 key 没有设置剩余生存时间时，命令都返回 -1 。
*/
func (this *RedisPool) PTTL(key string) (pttl int64, exist bool, persisted bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	pttl, err = redis.Int64(conn.Do("PTTL", key))
	if err != nil {
		return
	}

	if pttl == -2 {
		exist = false
		persisted = false
	} else if pttl == -1 {
		exist = true
		persisted = true
	} else {
		exist = true
		persisted = false
	}

	return
}
