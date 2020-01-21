/*
redisUtil对Redis的连接池进行了一定程度的封装
将常用的方法进行了内部封装，对于不常见的方法，有两种处理方式：
1、向作者提出请求，由作者添加到代码中
2、调用GetConnection方法，然后自己实现逻辑
在代码中，统一将conn.Do的结果和redis.Int,redis.String等类型转换合并处理

redis的命令请参考：https://redis.readthedocs.io/en/2.6/index.html
*/
package redisUtil

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
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

// 关闭自定义连接池
func (this *RedisPool) Close() {
	this.pool.Close()
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

// 创建新的Redis连接池对象(obsolete，建议使用NewRedisPool2)
// name:连接池对象名称
// connectionString:Redis服务器连接地址
// password：Redis服务器连接密码
// database：Redis服务器选择的数据库
// maxActive：Redis连接池允许的最大活跃连接数量
// maxIdle：Redis连接池允许的最大空闲数量
// idleTimeout：连接被回收前的空闲时间
// dialConnectTimeout：连接Redis服务器超时时间
// 返回值：
// Redis连接池对象
func NewRedisPool(name, connectionString, password string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) *RedisPool {
	redisConfig := NewRedisConfig2(connectionString, password, database, maxActive, maxIdle, idleTimeout, dialConnectTimeout)
	return NewRedisPool2(name, redisConfig)
}

// 创建新的Redis连接池对象
// name:连接池对象名称
// redisConfig:Redis配置对象
// 返回值：
// Redis连接池对象
func NewRedisPool2(name string, redisConfig *RedisConfig) *RedisPool {
	poolObj := &redis.Pool{
		MaxActive:   redisConfig.MaxActive,
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: redisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0, 4)
			options = append(options, redis.DialConnectTimeout(redisConfig.DialConnectTimeout))
			options = append(options, redis.DialDatabase(redisConfig.Database))
			if redisConfig.Password != "" {
				options = append(options, redis.DialPassword(redisConfig.Password))
			}

			conn, err := redis.Dial("tcp", redisConfig.ConnectionString, options...)
			if err != nil {
				return nil, fmt.Errorf("Dial failed, err:%s", err)
			}

			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}

	return &RedisPool{
		name:    name,
		address: redisConfig.ConnectionString,
		pool:    poolObj,
	}
}
