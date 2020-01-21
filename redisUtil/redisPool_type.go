package redisUtil

import (
	"github.com/gomodule/redigo/redis"
)

func (this *RedisPool) Int(reply interface{}) (int, error) {
	return redis.Int(reply, nil)
}

func (this *RedisPool) Int64(reply interface{}) (int64, error) {
	return redis.Int64(reply, nil)
}

func (this *RedisPool) Uint64(reply interface{}) (uint64, error) {
	return redis.Uint64(reply, nil)
}

func (this *RedisPool) Float64(reply interface{}) (float64, error) {
	return redis.Float64(reply, nil)
}

func (this *RedisPool) String(reply interface{}) (string, error) {
	return redis.String(reply, nil)
}

func (this *RedisPool) Bytes(reply interface{}) ([]byte, error) {
	return redis.Bytes(reply, nil)
}

func (this *RedisPool) Bool(reply interface{}) (bool, error) {
	return redis.Bool(reply, nil)
}

func (this *RedisPool) Values(reply interface{}) ([]interface{}, error) {
	return redis.Values(reply, nil)
}

func (this *RedisPool) Ints(reply interface{}) ([]int, error) {
	return redis.Ints(reply, nil)
}

func (this *RedisPool) Int64s(reply interface{}) ([]int64, error) {
	return redis.Int64s(reply, nil)
}

func (this *RedisPool) Float64s(reply interface{}) ([]float64, error) {
	return redis.Float64s(reply, nil)
}

func (this *RedisPool) Strings(reply interface{}) ([]string, error) {
	return redis.Strings(reply, nil)
}

func (this *RedisPool) ByteSlices(reply interface{}) ([][]byte, error) {
	return redis.ByteSlices(reply, nil)
}

func (this *RedisPool) IntMap(reply interface{}) (map[string]int, error) {
	return redis.IntMap(reply, nil)
}

func (this *RedisPool) Int64Map(reply interface{}) (map[string]int64, error) {
	return redis.Int64Map(reply, nil)
}

func (this *RedisPool) StringMap(reply interface{}) (map[string]string, error) {
	return redis.StringMap(reply, nil)
}

func (this *RedisPool) Positions(reply interface{}) ([]*[2]float64, error) {
	return redis.Positions(reply, nil)
}
