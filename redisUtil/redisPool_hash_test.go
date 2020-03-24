package redisUtil

import (
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPoolObj_hash *RedisPool
)

func init() {
	redisPoolObj_hash = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestHSet(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		代码示例
		设置一个新域：

		redis> HSET website google "www.g.cn"
		(integer) 1
	*/
	key := "website"
	field := "google"
	value := "www.g.cn"
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HGET website google
		"www.g.cn"
	*/
	expected2 := "www.g.cn"
	got2, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s and field:%s should exist, but now it doesn't", key, field)
		return
	}
	got2_str, err := redis.String(got2, err)
	if err != nil {
		t.Fail()
	}
	if got2_str != expected2 {
		t.Errorf("Expected to get %s, but got %s", expected2, got2_str)
		return
	}

	/*
		对一个已存在的域进行更新：
		redis> HSET website google "www.google.com"
		(integer) 0

	*/
	field = "google"
	value = "www.google.com"
	expected = 0
	got, err = redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}

	/*
		redis> HGET website google
		"www.google.com"
	*/
	expected3 := "www.google.com"
	got3, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s and field:%s should exist, but now it doesn't", key, field)
		return
	}
	got3_str, err := redis.String(got3, nil)
	if err != nil {
		t.Fail()
	}
	if got3_str != expected3 {
		t.Errorf("Expected to get %s, but got %s", expected3, got3_str)
		return
	}

	/*
		对一个已存在的域进行更新：
		redis> HSET website google 1
		(integer) 0

	*/
	field = "google"
	value4 := 1
	expected = 0
	got, err = redisPoolObj_hash.HSet(key, field, value4)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}

	/*
		redis> HGET website google
		"www.google.com"
	*/
	expected4 := 1
	got4, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s and field:%s should exist, but now it doesn't", key, field)
		return
	}
	got4_int, err := redis.Int(got4, nil)
	if err != nil {
		t.Fail()
	}
	if got4_int != expected4 {
		t.Errorf("Expected to get %d, but got %d", expected4, got4_int)
		return
	}
}

func TestHSetNX(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		代码示例
		域尚未存在， 设置成功：

		redis> HSETNX database key-value-store Redis
		(integer) 1
	*/
	key := "database"
	field := "key-value-store"
	value := "Redis"
	successful, err := redisPoolObj_hash.HSetNX(key, field, value)
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("HSetNX key:%s, field:%s should be successful, but now it's not.", key, field)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HGET database key-value-store
		"Redis"
	*/
	expected := "Redis"
	gottmp, exist, err := redisPoolObj_hash.HGet(key, field)
	got, err := redis.String(gottmp, err)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("Key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	/*
		域已经存在， 设置未成功， 域原有的值未被改变：

		redis> HSETNX database key-value-store Riak
		(integer) 0
	*/
	value = "Riak"
	successful, err = redisPoolObj_hash.HSetNX(key, field, value)
	if err != nil {
		t.Fail()
	}
	if successful {
		t.Errorf("HSetNX key:%s, field:%s should be not successful, but now it is.", key, field)
		return
	}

	/*
		redis> HGET database key-value-store
		"Redis"
	*/
	expected = "Redis"
	gottmp, exist, err = redisPoolObj_hash.HGet(key, field)
	got, err = redis.String(gottmp, err)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("Key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}
}

func TestHGet(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		代码示例
		域存在的情况：

		redis> HSET homepage redis redis.com
		(integer) 1
	*/
	key := "homepage"
	field := "redis"
	value := "redis.com"
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HGET homepage redis
		"redis.com"
	*/
	expected2 := "redis.com"
	got2, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s and field:%s should exist, but now it doesn't", key, field)
		return
	}
	got2_str, err := redis.String(got2, err)
	if err != nil {
		t.Fail()
	}
	if got2_str != expected2 {
		t.Errorf("Expected to get %s, but got %s", expected2, got2_str)
		return
	}

	/*
		域不存在的情况：

		redis> HGET site mysql
		(nil)
	*/
	key = "site"
	field = "mysql"
	_, exist, err = redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s and field:%s should not exist, but now it does", key, field)
		return
	}
}

func TestHExists(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		代码示例
		给定域不存在：

		redis> HEXISTS phone myphone
		(integer) 0
		给定域存在：
	*/
	key := "phone"
	field := "myphone"
	expected := false
	got, err := redisPoolObj_hash.HExists(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	/*
		redis> HSET phone myphone nokia-1110
		(integer) 1
	*/
	value := "nokia-1110"
	expected2 := 1
	got2, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HEXISTS phone myphone
		(integer) 1
	*/
	expected3 := true
	got3, err := redisPoolObj_hash.HExists(key, field)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %t, but now got %t", expected3, got3)
		return
	}
}

func TestHDel(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		# 测试数据
	*/
	key := "abbr"
	redisPoolObj_hash.HSet(key, "a", "apple")
	redisPoolObj_hash.HSet(key, "b", "banana")
	redisPoolObj_hash.HSet(key, "c", "cat")
	redisPoolObj_hash.HSet(key, "d", "dog")
	deleteKeys = append(deleteKeys, key)

	/*
		# 删除单个域

		redis> HDEL abbr a
		(integer) 1
	*/
	field := "a"
	expected := 1
	got, err := redisPoolObj_hash.HDel(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		# 删除不存在的域

		redis> HDEL abbr not-exists-field
		(integer) 0
	*/
	field = "not-exists-field"
	expected = 0
	got, err = redisPoolObj_hash.HDel(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		# 删除多个域

		redis> HDEL abbr b c
		(integer) 2
	*/
	fields := []string{"b", "c"}
	expected = 2
	got, err = redisPoolObj_hash.HDel(key, fields...)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HGETALL abbr
		1) "d"
		2) "dog"
	*/
}

func TestHLen(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		redis> HSET db redis redis.com
		(integer) 1
	*/
	key := "db"
	field := "redis"
	value := "redis.com"
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HSET db mysql mysql.com
		(integer) 1
	*/
	field = "mysql"
	value = "mysql.com"
	expected = 1
	got, err = redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HLEN db
		(integer) 2
	*/
	expected = 2
	got, err = redisPoolObj_hash.HLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HSET db mongodb mongodb.org
		(integer) 1
	*/
	field = "mongodb"
	value = "mongodb.org"
	expected = 1
	got, err = redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HLEN db
		(integer) 3
	*/
	expected = 3
	got, err = redisPoolObj_hash.HLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
}

func TestHStrlen(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		redis> HMSET myhash f1 "HelloWorld" f2 "99" f3 "-256"
		OK
	*/
	key := "myhash"
	redisPoolObj_hash.HSet(key, "f1", "HelloWorld")
	redisPoolObj_hash.HSet(key, "f2", "99")
	redisPoolObj_hash.HSet(key, "f3", "-256")
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HSTRLEN myhash f1
		(integer) 10
	*/
	expected := 10
	field := "f1"
	got, err := redisPoolObj_hash.HStrlen(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HSTRLEN myhash f2
		(integer) 2
	*/
	expected = 2
	field = "f2"
	got, err = redisPoolObj_hash.HStrlen(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HSTRLEN myhash f3
		(integer) 4
	*/
	expected = 4
	field = "f3"
	got, err = redisPoolObj_hash.HStrlen(key, field)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
}

func TestHIncrBy(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		# increment 为正数

		redis> HEXISTS counter page_view    # 对空域进行设置
		(integer) 0
	*/
	key := "counter"
	field := "page_view"
	exist, err := redisPoolObj_hash.HExists(key, field)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s, field:%s should not exist, but now it does.", key, field)
		return
	}

	/*
		redis> HINCRBY counter page_view 200
		(integer) 200
	*/
	increment := int64(200)
	expected := int64(200)
	got, err := redisPoolObj_hash.HIncrBy(key, field, increment)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HGET counter page_view
		"200"
	*/
	expected = 200
	got2, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	got2_int64, err := redis.Int64(got2, err)
	if err != nil {
		t.Fail()
	}
	if got2_int64 != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got2_int64)
		return
	}

	/*
		# increment 为负数

		redis> HGET counter page_view
		"200"
	*/
	/*
		redis> HINCRBY counter page_view -50
		(integer) 150
	*/
	increment = -50
	expected = 150
	got, err = redisPoolObj_hash.HIncrBy(key, field, increment)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> HGET counter page_view
		"150"
	*/
	expected = 150
	got3, exist, err := redisPoolObj_hash.HGet(key, field)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	got3_int64, err := redis.Int64(got3, err)
	if err != nil {
		t.Fail()
	}
	if got3_int64 != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got3_int64)
		return
	}

	/*
		# 尝试对字符串值的域执行HINCRBY命令

		redis> HSET myhash string hello,world       # 设定一个字符串值
		(integer) 1
	*/
	key2 := "myhash"
	field2 := "string"
	value2 := "hello,world"
	expected4 := 1
	got4, err := redisPoolObj_hash.HSet(key2, field2, value2)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %d, but now got %d", expected4, got4)
		return
	}
	deleteKeys = append(deleteKeys, key2)

	/*
		redis> HGET myhash string
		"hello,world"
	*/
	expected5 := "hello,world"
	got5_interface, exist, err := redisPoolObj_hash.HGet(key2, field2)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	got5, err := redis.String(got5_interface, err)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected to get %s, but now got %s", expected5, got5)
		return
	}

	/*
		redis> HINCRBY myhash string 1              # 命令执行失败，错误。
		(error) ERR hash value is not an integer
	*/
	increment = 1
	got, err = redisPoolObj_hash.HIncrBy(key2, field2, increment)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	/*
		redis> HGET myhash string                   # 原值不变
		"hello,world"
	*/
	expected6 := "hello,world"
	got6_interface, exist, err := redisPoolObj_hash.HGet(key2, field2)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field)
		return
	}
	got6, err := redis.String(got6_interface, err)
	if err != nil {
		t.Fail()
	}
	if got6 != expected6 {
		t.Errorf("Expected to get %s, but now got %s", expected6, got6)
		return
	}
}

func TestHIncrByFloat(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		# 值和增量都是普通小数

		redis> HSET mykey field 10.50
		(integer) 1
		redis> HINCRBYFLOAT mykey field 0.1
		"10.6"
	*/
	key := "mykey"
	field := "field"
	value := 10.50
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	increment := 0.1
	expected2 := 10.6
	got2, err := redisPoolObj_hash.HIncrByFloat(key, field, increment)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %f, but now got %f", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		# 值和增量都是指数符号

		redis> HSET mykey field 5.0e3
		(integer) 0
		redis> HINCRBYFLOAT mykey field 2.0e2
		"5200"
	*/
	value3 := 5.0e3
	expected3 := 0
	got3, err := redisPoolObj_hash.HSet(key, field, value3)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}

	increment4 := 2.0e2
	expected4 := 5200.0
	got4, err := redisPoolObj_hash.HIncrByFloat(key, field, increment4)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %f, but now got %f", expected4, got4)
		return
	}

	/*
		# 对不存在的键执行 HINCRBYFLOAT

		redis> EXISTS price
		(integer) 0
		redis> HINCRBYFLOAT price milk 3.5
		"3.5"
		redis> HGETALL price
		1) "milk"
		2) "3.5"
	*/
	key5 := "price"
	exist5, err := redisPoolObj_hash.Exists(key5)
	if err != nil {
		t.Fail()
	}
	if exist5 {
		t.Errorf("The key:%s should not exist, but now it does.", key5)
		return
	}

	field5 := "milk"
	increment5 := 3.5
	expected5 := 3.5
	got5, err := redisPoolObj_hash.HIncrByFloat(key5, field5, increment5)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected to get %f, but now got %f", expected5, got5)
		return
	}

	got5_interface, exist, err := redisPoolObj_hash.HGet(key5, field5)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't", key5, field5)
		return
	}
	got5, err = redis.Float64(got5_interface, err)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected to get %f, but now got %f", expected5, got5)
		return
	}

	deleteKeys = append(deleteKeys, key5)

	/*
		# 对不存在的域进行 HINCRBYFLOAT

		redis> HGETALL price
		1) "milk"
		2) "3.5"
		redis> HINCRBYFLOAT price coffee 4.5   # 新增 coffee 域
		"4.5"
		redis> HGETALL price
		1) "milk"
		2) "3.5"
		3) "coffee"
		4) "4.5"
	*/
	key6 := "price"
	field6 := "coffee"
	increment6 := 4.5
	expected6 := 4.5
	got6, err := redisPoolObj_hash.HIncrByFloat(key6, field6, increment6)
	if err != nil {
		t.Fail()
	}
	if got6 != expected6 {
		t.Errorf("Expected to get %f, but now got %f", expected6, got6)
		return
	}

	got6_interface, exist, err := redisPoolObj_hash.HGet(key6, field6)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't", key6, field6)
		return
	}
	got6, err = redis.Float64(got6_interface, err)
	if got6 != expected6 {
		t.Errorf("Expected to get %f, but now got %f", expected6, got6)
		return
	}
}

func TestHMSet(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	// HMSet
	key := "people"
	people := &People{
		Jack: "Jack Ma",
		Gump: "Gump Li",
	}
	err := redisPoolObj_hash.HMSet(key, people)
	if err != nil {
		t.Fail()
	}

	data := make(map[string]interface{})
	data["Jordan"] = "Jordan Zuo"
	err = redisPoolObj_hash.HMSet(key, data)
	if err != nil {
		t.Fail()
	}

	deleteKeys = append(deleteKeys, key)

	// HGet
	field1 := "Jack"
	expected1 := "Jack Ma"
	got1_interface, exist1, err := redisPoolObj_hash.HGet(key, field1)
	if err != nil {
		t.Fail()
	}
	if !exist1 {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field1)
		return
	}
	got1, err := redis.String(got1_interface, nil)
	if err != nil {
		t.Fail()
	}
	if got1 != expected1 {
		t.Errorf("Expected to get %s, but now got %s", expected1, got1)
		return
	}

	// HGet
	field2 := "Gump"
	expected2 := "Gump Li"
	got2_interface, exist2, err := redisPoolObj_hash.HGet(key, field2)
	if err != nil {
		t.Fail()
	}
	if !exist2 {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field2)
		return
	}
	got2, err := redis.String(got2_interface, nil)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}

	field3 := "Jordan"
	expected3 := "Jordan Zuo"
	got3_interface, exist3, err := redisPoolObj_hash.HGet(key, field3)
	if err != nil {
		t.Fail()
	}
	if !exist3 {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field3)
		return
	}
	got3, err := redis.String(got3_interface, nil)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %s, but now got %s", expected3, got3)
		return
	}
}

func TestHMGet(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		redis> HMSET pet dog "doudou" cat "nounou"    # 一次设置多个域
		OK
	*/
	key := "pet"
	field_value_map := make(map[string]interface{})
	field_value_map["dog"] = "doudou"
	field_value_map["cat"] = "nounou"
	err := redisPoolObj_hash.HMSet(key, field_value_map)
	if err != nil {
		t.Fail()
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HMGET pet dog cat fake_pet             # 返回值的顺序和传入参数的顺序一样
		1) "doudou"
		2) "nounou"
		3) (nil)                                      # 不存在的域返回nil值
	*/
	expected := []string{"doudou", "nounou", ""}
	reply, exist, err := redisPoolObj_hash.HMGet(key, "dog", "cat", "fake_pet")
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should have items, but now it doesn't.", key)
		return
	}
	got, err := redisPoolObj_hash.Strings(reply)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}
}

func TestHKeys(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		# 哈希表非空

		redis> HMSET website google www.google.com yahoo www.yahoo.com
		OK
	*/
	key := "website"
	field_value_map := make(map[string]interface{})
	field_value_map["google"] = "www.google.com"
	field_value_map["yahoo"] = "www.yahoo.com"
	err := redisPoolObj_hash.HMSet(key, field_value_map)
	if err != nil {
		t.Fail()
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HKEYS website
		1) "google"
		2) "yahoo"
	*/
	expected := make([]string, 0, len(field_value_map))
	for k := range field_value_map {
		expected = append(expected, k)
	}
	got, err := redisPoolObj_hash.HKeys(key)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}

	/*
		# 空哈希表/key不存在

		redis> EXISTS fake_key
		(integer) 0
	*/
	key2 := "fake_key"
	exist, err := redisPoolObj_hash.Exists(key2)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key2)
		return
	}

	/*
		redis> HKEYS fake_key
		(empty list or set)*
	*/
	expected2 := make([]string, 0, len(field_value_map))
	got2, err := redisPoolObj_hash.HKeys(key2)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(got2, expected2) == false {
		t.Errorf("Expected to get %v, but got %v", expected2, got2)
	}
}

func TestHVals(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		# 哈希表非空

		redis> HMSET website google www.google.com yahoo www.yahoo.com
		OK
	*/
	key := "website"
	field_value_map := make(map[string]interface{})
	field_value_map["google"] = "www.google.com"
	field_value_map["yahoo"] = "www.yahoo.com"
	err := redisPoolObj_hash.HMSet(key, field_value_map)
	if err != nil {
		t.Fail()
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HVALS website
		1) "www.google.com"
		2) "www.yahoo.com"
	*/
	expected := make([]string, 0, len(field_value_map))
	for _, v := range field_value_map {
		if v_str, ok := v.(string); ok {
			expected = append(expected, v_str)
		}
	}
	reply, err := redisPoolObj_hash.HVals(key)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_hash.Strings(reply)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}

	/*
		# 空哈希表/不存在的key

		redis> EXISTS not_exists
		(integer) 0
	*/
	key2 := "fake_key"
	exist, err := redisPoolObj_hash.Exists(key2)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key2)
		return
	}

	/*
		redis> HVALS not_exists
		(empty list or set)
	*/
	expected2 := make([]string, 0, len(field_value_map))
	for _, v := range field_value_map {
		if v_str, ok := v.(string); ok {
			expected2 = append(expected2, v_str)
		}
	}
	reply2, err := redisPoolObj_hash.HVals(key)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_hash.Strings(reply2)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(got2, expected2) == false {
		t.Errorf("Expected to get %v, but got %v", expected2, got2)
	}
}

func TestHGetAll(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		redis> HSET people jack "Jack Sparrow"
		(integer) 1
	*/
	key := "people"
	field1 := "Jack"
	value1 := "Jack Sparrow"
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field1, value1)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HSET people gump "Forrest Gump"
		(integer) 1
	*/
	field2 := "Gump"
	value2 := "Forrest Gump"
	expected = 1
	got, err = redisPoolObj_hash.HSet(key, field2, value2)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HGETALL people
		1) "Jack"          # 域
		2) "Jack Sparrow"  # 值
		3) "Gump"
		4) "Forrest Gump"
	*/

	reply, err := redisPoolObj_hash.HGetAll(key)
	if err != nil {
		t.Fail()
	}
	field_value_map, err := redisPoolObj_hash.StringMap(reply)
	if err != nil {
		t.Fail()
	}

	expected1 := "Jack Sparrow"
	got1, exist1 := field_value_map[field1]
	if !exist1 {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field1)
		return
	}
	if got1 != expected1 {
		t.Errorf("Expected to get %s, but now got %s", expected1, got1)
		return
	}

	expected2 := "Forrest Gump"
	got2, exist2 := field_value_map[field2]
	if !exist2 {
		t.Errorf("The key:%s, field:%s should exist, but now it doesn't.", key, field2)
		return
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}
}

func TestHGetAll_Struct(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_string.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
		redis> HSET people jack "Jack Sparrow"
		(integer) 1
	*/
	key := "people"
	field := "Jack"
	value := "Jack Sparrow"
	expected := 1
	got, err := redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> HSET people gump "Forrest Gump"
		(integer) 1
	*/
	field = "Gump"
	value = "Forrest Gump"
	expected = 1
	got, err = redisPoolObj_hash.HSet(key, field, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> HGETALL people
		1) "Jack"          # 域
		2) "Jack Sparrow"  # 值
		3) "Gump"
		4) "Forrest Gump"
	*/

	got2 := new(People)
	expected2 := &People{
		Jack: "Jack Sparrow",
		Gump: "Forrest Gump",
	}
	exist, err := redisPoolObj_hash.HGetAll_Struct(key, got2)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2.TheSame(expected2) == false {
		t.Errorf("Expected to get:%s, but now got %s", expected2, got2)
	}
}

type People struct {
	Jack string
	Gump string
}

func (this *People) TheSame(other *People) bool {
	return this.Jack == other.Jack && this.Gump == other.Gump
}

func (this *People) String() string {
	return fmt.Sprintf("{\"Jack\": %s, \"Gump\":%s}", this.Jack, this.Gump)
}
