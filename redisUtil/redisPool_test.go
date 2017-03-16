package redisUtil

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	redisPoolObj *RedisPool
)

func init() {
	redisPoolObj = NewRedisPool("testPool", "10.1.0.21:6379", "redis_pwd", 3, 500, 200, 300*time.Second, 10*time.Second)
}

func TestGetName(t *testing.T) {
	expectedName := "testPool"
	if actualName := redisPoolObj.GetName(); actualName != expectedName {
		t.Errorf("GetName should be %s, but got %s", expectedName, actualName)
	}
}

func TestGetAddress(t *testing.T) {
	exptectedAddress := "10.1.0.21:6379"
	if actualAddress := redisPoolObj.GetAddress(); actualAddress != exptectedAddress {
		t.Errorf("GetAddress should be %s, but got %s", exptectedAddress, actualAddress)
	}
}

func TestExists(t *testing.T) {
	key := "testExists"
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if exists {
		t.Errorf("it should be not exists, but now exists")
	}

	redisPoolObj.Set(key, "test")
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if !exists {
		t.Errorf("it should be exists, but now not exists")
	}

	if count, err := redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if count != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestKeys(t *testing.T) {
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 0 {
		t.Errorf("there should be no keys, but now got %d.", len(keyList))
		return
	}

	redisPoolObj.Set("key1", "key1")
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 1 {
		t.Errorf("there should be 1 keys, but now got %d.", len(keyList))
		return
	}

	redisPoolObj.Set("key2", "key2")
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 2 {
		t.Errorf("there should be 2 keys, but now got %d.", len(keyList))
		return
	}

	if count, err := redisPoolObj.Del("key1", "key2"); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if count != 2 {
		t.Errorf("Del should return 2, but now return %d", count)
	}
}

func TestExpire(t *testing.T) {
	key := "expire"
	if success, err := redisPoolObj.Expire(key, 2); err != nil {
		t.Errorf("Expire failed, err:%s", err)
	} else if success {
		t.Errorf("Expire expected fail, but now got success")
	}

	redisPoolObj.Set(key, "test")

	if success, err := redisPoolObj.Expire(key, 2); err != nil {
		t.Errorf("Expire failed, err:%s", err)
	} else if !success {
		t.Errorf("Expire expected success, but now got fail")
	}

	time.Sleep(2 * time.Second)
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if exists {
		t.Errorf("it should be not exists, but now exists")
	}
}

func TestGet(t *testing.T) {
	key := "get"

	if _, exists, err := redisPoolObj.Get(key); err != nil {
		t.Errorf("Get failed, err:%s", err)
	} else if exists {
		t.Errorf("Get should be not exists, but now exists.")
	}

	redisPoolObj.Set(key, "set")

	if value, exists, err := redisPoolObj.Get(key); err != nil {
		t.Errorf("Get failed, err:%s", err)
	} else if !exists {
		t.Errorf("Get should be  exists, but now not exists.")
	} else if value != "set" {
		t.Errorf("Get value should be %s, but now got %s", "set", value)
	}

	redisPoolObj.Del(key)
}

func TestHGet(t *testing.T) {
	key := "hget"
	field := "name"
	value := "jordan"

	if _, exists, err := redisPoolObj.HGet(key, field); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if exists {
		t.Errorf("HGET should be not exists, but now exists")
	}

	redisPoolObj.HSet(key, field, value)

	if actualValue, exists, err := redisPoolObj.HGet(key, field); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if value != actualValue {
		t.Errorf("HGET expected got %s, but now got %s", value, actualValue)
	}

	redisPoolObj.Del(key)
}

func TestHSet(t *testing.T) {
	key := "hget"
	field1 := "name"
	value1 := "jordan"
	field2 := "age"
	value2 := 32

	if err := redisPoolObj.HSet(key, field1, value1); err != nil {
		t.Errorf("HSET failed, err:%s", err)
	}

	if actualValue, exists, err := redisPoolObj.HGet(key, field1); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if actualValue != value1 {
		t.Errorf("HGET expected got %d, but now got %s", value1, actualValue)
	}

	if err := redisPoolObj.HSet(key, field2, value2); err != nil {
		t.Errorf("HSET failed, err:%s", err)
	}

	if actualValue, exists, err := redisPoolObj.HGet(key, field2); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if actualValue2, err := strconv.Atoi(actualValue); err != nil || actualValue2 != value2 {
		t.Errorf("HGET expected got %d, but now got %s", value2, actualValue)
	}

	redisPoolObj.Del(key)
}

type Player struct {
	Name string
	Age  int
}

func TestHGetAll(t *testing.T) {
	key := "player"

	p := &Player{}
	if exists, err := redisPoolObj.HGetAll(key, p); err != nil {
		t.Errorf("HGETALL failed, err:%s", err)
	} else if exists {
		t.Errorf("HGETALL should be not exists, but now exists.")
	}

	p1 := &Player{
		Name: "jordan",
		Age:  32,
	}

	if err := redisPoolObj.HMSet(key, p1); err != nil {
		t.Errorf("HMSET failed, err:%s", err)
	}

	if exists, err := redisPoolObj.HGetAll(key, p); err != nil {
		t.Errorf("HGETALL failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGETALL should be exists, but now not exists.")
	} else {
		fmt.Printf("player:%v\n", p)
	}

	redisPoolObj.Del(key)
}

func TestLRange(t *testing.T) {
	key := "list"
	start := 0
	stop := -1

	if list, err := redisPoolObj.LRange(key, start, stop); err != nil {
		t.Errorf("LRANGE failed, err:%s", err)
	} else if len(list) != 0 {
		t.Errorf("LRANGE expected 0 item, but now %d items", len(list))
	}

	if newCount, err := redisPoolObj.LPush(key, "1"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 1 {
		t.Errorf("LPUSH expected got 1, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "2"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 2 {
		t.Errorf("LPUSH expected got 2, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "3"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 3 {
		t.Errorf("LPUSH expected got 3, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "3"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 4 {
		t.Errorf("LPUSH expected got 4, but now got %d", newCount)
	}

	if list, err := redisPoolObj.LRange(key, start, stop); err != nil {
		t.Errorf("LRANGE failed, err:%s", err)
	} else if len(list) != 4 {
		t.Errorf("LRANGE expected 4 item, but now %d items", len(list))
	} else {
		for _, item := range list {
			fmt.Println(item)
		}
	}

	if removeCount, err := redisPoolObj.LRem(key, 0, "3"); err != nil {
		t.Errorf("LRem failed, err:%s", err)
	} else if removeCount != 2 {
		t.Errorf("LREM expected got 2, but now got %d", removeCount)
	}

	if item, exists, err := redisPoolObj.LPop(key); err != nil {
		t.Errorf("LPOP failed, err:%s", err)
	} else if !exists {
		t.Errorf("LPOP should be exists, but now not exists")
	} else if item != "1" {
		t.Errorf("LPOP should got 1, but now got %s", item)
	}

	if item, exists, err := redisPoolObj.RPop(key); err != nil {
		t.Errorf("RPOP failed, err:%s", err)
	} else if !exists {
		t.Errorf("RPOP should be exists, but now not exists")
	} else if item != "2" {
		t.Errorf("RPOP should got 2, but now got %s", item)
	}

	if _, exists, err := redisPoolObj.RPop(key); err != nil {
		t.Errorf("RPOP failed, err:%s", err)
	} else if exists {
		t.Errorf("RPOP should be not exists, but now exists")
	}
}
