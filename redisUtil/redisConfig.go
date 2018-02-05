package redisUtil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Redis配置对象
type RedisConfig struct {
	// 连接字符串
	ConnectionString string

	// 密码
	Password string

	// 数据库编号
	Database int

	// 最大活跃连接数
	MaxActive int

	// 最大空闲连接数
	MaxIdle int

	// 空闲超时
	IdleTimeout time.Duration

	// 连接超时
	DialConnectTimeout time.Duration
}

// 将redis连接字符串转化为redis config对象
// 格式：ConnectionString=10.1.0.21:6379;Password=redis_pwd;Database=3;MaxActive=50;MaxIdle=20;IdleTimeout=300;DialConnectTimeout=10;
// redisConfigStr：redis连接字符串
// 返回值：
// redis config对象
// 错误对象
func NewRedisConfig(redisConfigStr string) (redisConfig *RedisConfig, err error) {
	var connectionString string
	var password string
	var database int
	var maxActive int
	var maxIdle int
	var idleTimeout time.Duration
	var dialConectTimeout time.Duration
	var count int = 7
	var subCount int = 2

	itemList := strings.Split(redisConfigStr, ";")
	// 去掉最后的空数据
	if itemList[len(itemList)-1] == "" {
		itemList = itemList[0 : len(itemList)-1]
	}
	if len(itemList) != count {
		err = fmt.Errorf("%s格式不正确，需要包含%d个部分，现在有%d个部分", redisConfigStr, count, len(itemList))
		return
	}

	for _, item := range itemList {
		subItemList := strings.Split(item, "=")
		if len(subItemList) != subCount {
			err = fmt.Errorf("%s格式不正确，需要包含%d个部分", item, subCount)
			return
		}

		// 分别进行判断
		switch strings.ToLower(subItemList[0]) {
		case strings.ToLower("ConnectionString"):
			connectionString = subItemList[1]
		case strings.ToLower("Password"):
			password = subItemList[1]
		case strings.ToLower("Database"):
			if database, err = strconv.Atoi(subItemList[1]); err != nil {
				err = fmt.Errorf("%s转化为int型失败", subItemList[1])
				return
			}
		case strings.ToLower("MaxActive"):
			if maxActive, err = strconv.Atoi(subItemList[1]); err != nil {
				err = fmt.Errorf("%s转化为int型失败", subItemList[1])
				return
			}
		case strings.ToLower("MaxIdle"):
			if maxIdle, err = strconv.Atoi(subItemList[1]); err != nil {
				err = fmt.Errorf("%s转化为int型失败", subItemList[1])
				return
			}
		case strings.ToLower("IdleTimeout"):
			if idleTimeout_int, err1 := strconv.Atoi(subItemList[1]); err1 != nil {
				err = fmt.Errorf("%s转化为int型失败", subItemList[1])
				return
			} else {
				idleTimeout = time.Duration(idleTimeout_int) * time.Second
			}
		case strings.ToLower("DialConnectTimeout"):
			if dialConectTimeout_int, err1 := strconv.Atoi(subItemList[1]); err1 != nil {
				err = fmt.Errorf("%s转化为int型失败", subItemList[1])
				return
			} else {
				dialConectTimeout = time.Duration(dialConectTimeout_int) * time.Second
			}
		}
	}

	redisConfig = NewRedisConfig2(connectionString, password, database, maxActive, maxIdle, idleTimeout, dialConectTimeout)
	return
}

func NewRedisConfig2(connectionString, password string,
	database, maxActive, maxIdle int,
	idleTimeout, dialConnectTimeout time.Duration) *RedisConfig {

	return &RedisConfig{
		ConnectionString:   connectionString,
		Password:           password,
		Database:           database,
		MaxActive:          maxActive,
		MaxIdle:            maxIdle,
		IdleTimeout:        idleTimeout,
		DialConnectTimeout: dialConnectTimeout,
	}
}
