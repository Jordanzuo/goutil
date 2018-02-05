/*
redisUtil对Redis的连接池进行了一定程度的封装
将常用的方法进行了内部封装，对于不常见的方法，有两种处理方式：
1、向作者提出请求，由作者添加到代码中
2、调用GetConnection方法，然后自己实现逻辑
在代码中，统一将conn.Do的结果和redis.Int,redis.String等类型转换合并处理
*/
package redisUtil

import (
	"errors"
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

// 测试连接情况
// 返回值:
// 错误对象
func (this *RedisPool) TestConnection() error {
	conn := this.GetConnection()
	defer conn.Close()

	_, err := conn.Do("PING")

	return err
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

// 设置key和对应的value
// key:key
// value:value
// expireVal:超时时间值
// expireType:设置的超时类型
// 返回值:
// 错误对象
func (this *RedisPool) Set2(key string, value interface{}, expireType ExpireType, expireVal int) error {
	if expireType != Expire_Millisecond && expireType != Expire_Seond {
		return errors.New("ExpireTypeError")
	}

	conn := this.GetConnection()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, string(expireType), expireVal)

	return err
}

// 设置key和对应的value
// key:key
// value:value
// expireVal:超时时间值
// expireType:设置的超时类型
// setType:值存储类型
// 返回值:
// 错误对象
func (this *RedisPool) SetDetail(key string, value interface{}, expireType ExpireType, expireVal int, setType SetType) error {
	conn := this.GetConnection()
	defer conn.Close()

	paramList := []interface{}{key, value}

	// 超时设置组装
	if expireType != Expire_NoExpire {
		paramList = append(paramList, string(expireType), expireVal)
	}

	// 设置组装
	if setType != Set_Write {
		paramList = append(paramList, string(setType))
	}

	_, err := conn.Do("SET", paramList...)

	return err
}

// 当Key不存在时，设置key和对应的value
// key:key
// value:value
// 返回值:
// 错误对象
func (this *RedisPool) SetNX(key string, value interface{}) (success bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	result, err := redis.Int(conn.Do("SETNX", key, value))
	if err != nil {
		return
	}

	success = result == 1

	return
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

// 往set中添加value
// key:set的key
// values:待添加的值，如果value已经存在，则添加操作仍然成功，只是影响的记录数不会增加
// 返回值:
// newCount:影响的记录数
// err:错误信息
func (this *RedisPool) SAdd(key string, values ...string) (newCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, value := range values {
		args = append(args, value)
	}

	newCount, err = redis.Int(conn.Do("SADD", args...))

	return
}

// 获取指定key的set的记录数
// key:set的key
// values:待添加的值
// 返回值:
// nowCount:当前的记录数
// err:错误信息
func (this *RedisPool) SCard(key string) (nowCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	nowCount, err = redis.Int(conn.Do("SCARD", key))

	return
}

// 查看指定值是否在set中
// key:set的key
// values:待添加的值
// 返回值:
// isMember:是否是集合的成员
// err:错误信息
func (this *RedisPool) SIsMember(key string, value string) (isMember bool, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	isMember, err = redis.Bool(conn.Do("SISMEMBER", key, value))

	return
}

// 获取指定key的set的记录数
// key:set的key
// 返回值:
// values:获取到的值的个数
// err:错误信息
func (this *RedisPool) SMembers(key string) (values []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	values, err = redis.Strings(conn.Do("SMEMBERS", key))

	return
}

// 从set中随机获取一定数量的元素(获取而不移除)
// key:set的key
// count:需要获取的记录数，
//    如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合。
//    如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值。
// 返回值:
// values:获取的记录集合
// err:错误信息
func (this *RedisPool) SRandMember(key string, count int) (values []string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	values, err = redis.Strings(conn.Do("SRANDMEMBER", key, count))

	return
}

// 从set中获取并移除一项
// key:set的key
// 返回值:
// value:获取的记录项
// err:错误信息
func (this *RedisPool) SPop(key string) (value string, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	value, err = redis.String(conn.Do("SPOP", key))

	return
}

// 从set移除指定的value
// key:set的key
// values:需要移除的value集合（如果value项不存在，执行仍然成功）
// 返回值:
// delCount:成功移除的项的个数
// err:错误信息
func (this *RedisPool) SRem(key string, values ...string) (delCount int, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, value := range values {
		args = append(args, value)
	}

	delCount, err = redis.Int(conn.Do("SREM", args...))

	return
}

// 创建新的Redis连接池对象(obsolete，建议使用NewRedisPool2)
// _name:连接池对象名称
// connectionString:Redis服务器连接地址
// password：Redis服务器连接密码
// database：Redis服务器选择的数据库
// maxActive：Redis连接池允许的最大活跃连接数量
// maxIdle：Redis连接池允许的最大空闲数量
// idleTimeout：连接被回收前的空闲时间
// dialConnectTimeout：连接Redis服务器超时时间
// 返回值：
// Redis连接池对象
func NewRedisPool(_name, connectionString, password string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) *RedisPool {
	redisConfig := NewRedisConfig2(connectionString, password, database, maxActive, maxIdle, idleTimeout, dialConnectTimeout)
	return NewRedisPool2(_name, redisConfig)
}

// 创建新的Redis连接池对象
// _name:连接池对象名称
// redisConfig:Redis配置对象
// 返回值：
// Redis连接池对象
func NewRedisPool2(_name string, redisConfig *RedisConfig) *RedisPool {
	_pool := &redis.Pool{
		MaxActive:   redisConfig.MaxActive,
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: redisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0, 4)
			options = append(options, redis.DialConnectTimeout(redisConfig.DialConnectTimeout))
			if redisConfig.Password != "" {
				options = append(options, redis.DialPassword(redisConfig.Password))
			}
			options = append(options, redis.DialDatabase(redisConfig.Database))
			c, err := redis.Dial("tcp", redisConfig.ConnectionString, options...)
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
		address: redisConfig.ConnectionString,
		pool:    _pool,
	}
}
