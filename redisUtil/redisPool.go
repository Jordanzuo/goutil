/*
redisUtil对Redis的连接池进行了一定程度的封装
将常用的方法进行了内部封装，对于不常见的方法，有两种处理方式：
1、向作者提出请求，由作者添加到代码中
2、调用GetConnection方法，然后自己实现逻辑
在代码中，统一将conn.Do的结果和redis.Int,redis.String等类型转换合并处理
*/
package redisUtil

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 自定义Redis连接池对象
type RedisPool struct {
	name    string
	address string
	pool    *redis.Pool
}

// 获取自定义Redis连接池对象的名称
// 返回值:
// 自定义Redis连接池对象的名称
func (this *RedisPool) GetName() string {
	return this.name
}

// 获取自定义Redis连接池对象的目标地址
// 返回值:
// 自定义Redis连接池对象的目标地址
func (this *RedisPool) GetAddress() string {
	return this.address
}

// 从自定义连接池中获取连接，在使用后需要调用Close方法
// 返回值:
// 连接对象
func (this *RedisPool) GetConnection() redis.Conn {
	return this.pool.Get()
}

// 测试连接
// 返回值:
// 错误对象
func (this *RedisPool) Test() error {
	conn := this.GetConnection()
	defer conn.Close()

	return this.pool.TestOnBorrow(conn, time.Now())
}

// 关闭自定义连接池
func (this *RedisPool) Close() {
	this.pool.Close()
}

// 判断指定的Key是否存在
// key:key
// 返回值:
// 是否存在
// 错误对象
func (this *RedisPool) Exists(key string) (exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("EXISTS", key)); err != nil {
		return
	}

	exists = result == 1

	return
}

// 根据指定的模式获取匹配的key列表
// pattern:模式字符串
// 返回值:
// key列表
// 错误对象
func (this *RedisPool) Keys(pattern string) (keyList []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	keyList, err = redis.Strings(conn.Do("KEYS", pattern))

	return
}

// 删除指定的key列表
// keys:指定的key列表
// 返回值:
// 删除key的数量
// 错误对象
func (this *RedisPool) Del(keys ...string) (count int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := make([]interface{}, len(keys))
	for _, key := range keys {
		args = append(args, key)
	}

	count, err = redis.Int(conn.Do("DEL", args...))

	return
}

// 设置指定key的过期时间
// key:key
// seconds:过期的秒数
// 返回值:
// 是否成功
// 错误对象
func (this *RedisPool) Expire(key string, seconds int) (success bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	var result int
	if result, err = redis.Int(conn.Do("EXPIRE", key, seconds)); err != nil {
		return
	}

	success = result == 1

	return
}

// 获取指定key的内容
// key:key
// 返回值:
// 内容
// 是否存在
// 错误对象
func (this *RedisPool) Get(key string) (value string, exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	value, err = redis.String(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}

		return
	}

	exists = true

	return
}

// 设置key和对应的value
// key:key
// value:value
// 返回值:
// 错误对象
func (this *RedisPool) Set(key string, value interface{}) error {
	conn := this.GetConnection()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)

	return err
}

// 获取指定key的Hash表的field值
// key:key
// field:field
// 返回值:
// string类型的值
// 是否存在
// 错误对象
func (this *RedisPool) HGet(key, field string) (value string, exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	value, err = redis.String(conn.Do("HGET", key, field))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}

		return
	}

	exists = true

	return
}

// 设置指定key的Hash表的field的值
// key:key
// field:field
// value:value
// 返回值:
// 错误对象
func (this *RedisPool) HSet(key, field string, value interface{}) error {
	conn := this.GetConnection()
	defer conn.Close()

	_, err := conn.Do("HSET", key, field, value)

	return err
}

// 获取指定key的Hash表的所有field的值，并将其赋值给value对象
// key:key
// value:对象
// 返回值:
// 错误对象
func (this *RedisPool) HGetAll(key string, value interface{}) (exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	reply, err1 := redis.Values(conn.Do("HGETALL", key))
	if err1 != nil {
		err = err1
		return
	}

	if len(reply) == 0 {
		return
	}

	if err = redis.ScanStruct(reply, value); err != nil {
		return
	}

	exists = true

	return
}

// 将对象value的值赋值给key对应的Hast表
// key:key
// value:对象
// 返回值:
// 错误对象
func (this *RedisPool) HMSet(key string, value interface{}) error {
	conn := this.GetConnection()
	defer conn.Close()

	_, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)

	return err
}

// 获取key对应的List的指定区间的内容
// key:key
// start:开始区间
// stop:结束区间
// 返回值:
// 内容列表
// 错误对象
func (this *RedisPool) LRange(key string, start, stop int) (valueList []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	valueList, err = redis.Strings(conn.Do("LRANGE", key, start, stop))
	if err != nil {
		return
	}

	return
}

// 移除key对应的List中与value相等的指定数量的元素
// key:key
// count:指定数量
// value:匹配的内容
// 返回值:
// 删除的数量
// 错误对象
func (this *RedisPool) LRem(key string, count int, value string) (removeCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	removeCount, err = redis.Int(conn.Do("LREM", key, count, value))

	return
}

// 从左侧向key对应的List中追加数据
// key:key
// values:不定数量的值
// 返回值:
// 执行 LPUSH 命令后，列表的长度
// 错误对象
func (this *RedisPool) LPush(key string, values ...string) (newCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, value := range values {
		args = append(args, value)
	}

	newCount, err = redis.Int(conn.Do("LPUSH", args...))

	return
}

// 从右侧向key对应的List中追加数据
// key:key
// values:不定数量的值
// 返回值:
// 执行 LPUSH 命令后，列表的长度
// 错误对象
func (this *RedisPool) RPush(key string, values ...string) (newCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, value := range values {
		args = append(args, value)
	}

	newCount, err = redis.Int(conn.Do("RPush", args...))

	return
}

// 从左侧向key对应的List中移除数据并返回
// key:key
// 返回值:
// 列表的头元素
// 是否存在
// 错误对象
func (this *RedisPool) LPop(key string) (item string, exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.String(conn.Do("LPOP", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}

		return
	}

	exists = true

	return
}

// 从右侧向key对应的List中移除数据并返回
// key:key
// 返回值:
// 列表的尾部元素
// 是否存在
// 错误对象
func (this *RedisPool) RPop(key string) (item string, exists bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.String(conn.Do("RPOP", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}

		return
	}

	exists = true

	return
}

// 让指定Key递增1
// key:key
// 返回值:
// 递增之后的值
// 错误对象
func (this *RedisPool) Incr(key string) (item int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.Int64(conn.Do("INCR", key))
	return
}

// 让指定Key递增指定值
// key:key
// increment:增加值
// 返回值:
// 增加之后的值
// 错误对象
func (this *RedisPool) IncrBy(key string, increment int64) (item int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.Int64(conn.Do("INCRBY", key, increment))
	return
}

// 让指定key递减1
// key:key
// 返回值:
// 递减之后的值
// 错误对象
func (this *RedisPool) Decr(key string) (item int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.Int64(conn.Do("DECR", key))
	return
}

// 让指定key递减指定值
// key:key
// 返回值:
// 递减之后的值
// 错误对象
func (this *RedisPool) DecrBy(key string, decrement int64) (item int64, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	item, err = redis.Int64(conn.Do("DECRBY", key, decrement))
	return
}

// 创建新的Redis连接池对象
// _name:连接池对象名称
// _address:Redis服务器连接地址
// password：Redis服务器连接密码
// database：Redis服务器选择的数据库
// maxActive：Redis连接池允许的最大活跃连接数量
// maxIdle：Redis连接池允许的最大空闲数量
// idleTimeout：连接被回收前的空闲时间
// dialConnectTimeout：连接Redis服务器超时时间
func NewRedisPool(_name, _address, password string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) *RedisPool {
	_pool := &redis.Pool{
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0, 4)
			options = append(options, redis.DialConnectTimeout(dialConnectTimeout))
			if password != "" {
				options = append(options, redis.DialPassword(password))
			}
			options = append(options, redis.DialDatabase(database))
			c, err := redis.Dial("tcp", _address, options...)
			if err != nil {
				return nil, fmt.Errorf("Dial failed, err:%s", err)
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &RedisPool{
		name:    _name,
		address: _address,
		pool:    _pool,
	}
}
