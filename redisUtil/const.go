package redisUtil

// 过期时间类型
type ExpireType string

const (
	// 永不过期
	Expire_NoExpire ExpireType = ""

	// 秒
	Expire_Seond ExpireType = "EX"

	// 毫秒过期时间
	Expire_Millisecond ExpireType = "PX"
)

// 设置类型
type SetType string

const (
	// 直接覆盖
	Set_Write SetType = ""

	// 只在键不存在时，才对键进行设置操作。 SET key value NX 效果等同于 SETNX key value
	Set_NX SetType = "NX"

	// 只在键已经存在时，才对键进行设置操作
	Set_XX SetType = "XX"
)

// redis字段
type RedisField string

// redis关键字
type RedisKey string
