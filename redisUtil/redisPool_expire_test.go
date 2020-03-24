package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj_expire *RedisPool
)

func init() {
	redisPoolObj_expire = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestExpire(t *testing.T) {
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
	   redis> SET cache_page "www.google.com"
	   OK

	   redis> TTL cache_page
	   (integer) -1

	   redis> EXPIRE cache_page 30  # 设置过期时间为 30 秒
	   (integer) 1

	   redis> TTL cache_page    # 查看剩余生存时间
	   (integer) 23

	   redis> EXPIRE cache_page 30000   # 更新过期时间
	   (integer) 1

	   redis> TTL cache_page
	   (integer) 29996
	*/
	key := "cache_page"
	value := "www.google.com"
	expected := true
	got, err := redisPoolObj_expire.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	expected = true
	_, _, got, err = redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	seconds := int64(30)
	got, err = redisPoolObj_expire.Expire(key, seconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	expected2 := seconds - 5
	got2, exist, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}

	seconds = int64(3000)
	got, err = redisPoolObj_expire.Expire(key, seconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	expected2 = seconds - 5
	got2, exist, _, err = redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}
}

func TestExpireAt(t *testing.T) {
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
	   redis> SET cache www.google.com
	   OK

	   redis> EXPIREAT cache '1609403601'     # 这个 key 将在 Now()+1day 过期
	   (integer) 1

	   redis> TTL cache
	   (integer) 45081860
	*/
	key := "cache_page"
	value := "www.google.com"
	expected := true
	got, err := redisPoolObj_expire.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	timestamp := time.Now().AddDate(0, 0, 1).Unix()
	expected = true
	got, err = redisPoolObj_expire.ExpireAt(key, timestamp)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	expected2 := int64(86400 - 5)
	got2, exist, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}
}

func TestTTL(t *testing.T) {
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
	   # 不存在的 key

	   redis> TTL key
	   (integer) -2
	*/
	key := "key"
	expected := false
	_, got, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	/*
	   # key 存在，但没有设置剩余生存时间

	   redis> SET key value
	   OK

	   redis> TTL key
	   (integer) -1
	*/
	value := "value"
	redisPoolObj_expire.Set(key, value, "", 0, "")
	deleteKeys = append(deleteKeys, key)

	expected2 := int64(-1)
	got2, exist, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but got %d", expected2, got2)
		return
	}

	/*
	   # 有剩余生存时间的 key

	   redis> EXPIRE key 10086
	   (integer) 1

	   redis> TTL key
	   (integer) 10084
	*/
	seconds := int64(10086)
	expected = true
	got, err = redisPoolObj_expire.Expire(key, seconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	expected2 = seconds - 5
	got2, exist, _, err = redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}
}

func TestPersist(t *testing.T) {
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
	   redis> SET mykey "Hello"
	   OK
	*/
	key := "mykey"
	value := "Hello"
	redisPoolObj_expire.Set(key, value, "", 0, "")

	/*
	   redis> EXPIRE mykey 10  # 为 key 设置生存时间
	   (integer) 1
	*/
	seconds := int64(10)
	expected := true
	got, err := redisPoolObj_expire.Expire(key, seconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> TTL mykey
	   (integer) 10
	*/
	expected2 := seconds - 5
	got2, _, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get %d, but got %d", expected2, got2)
		return
	}

	/*
	   redis> PERSIST mykey    # 移除 key 的生存时间
	   (integer) 1
	*/
	redisPoolObj_expire.Persist(key)

	/*
	   redis> TTL mykey
	   (integer) -1
	*/
	expected = true
	_, _, got, err = redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
}

func TestPExpire(t *testing.T) {
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
	   redis> SET mykey "Hello"
	   OK
	*/
	key := "mykey"
	value := "Hello"
	expected := true
	got, err := redisPoolObj_expire.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> PEXPIRE mykey 1500
	   (integer) 1
	*/
	milliseconds := int64(1500)
	expected = true
	got, err = redisPoolObj_expire.PExpire(key, milliseconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	/*
	   redis> TTL mykey    # TTL 的返回值以秒为单位
	   (integer) 2

	   redis> PTTL mykey   # PTTL 可以给出准确的毫秒数
	   (integer) 1499
	*/
	expected2 := int64(1)
	got2, _, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got2 < expected2 {
		t.Errorf("TTL %d should be no less than %d", got2, expected2)
		return
	}

	expected2 = int64(1000)
	got2, _, _, err = redisPoolObj_expire.PTTL(key)
	if err != nil {
		t.Fail()
	}
	if got2 < expected2 {
		t.Errorf("TTL %d should be no less than %d", got2, expected2)
		return
	}
}

func TestPExpireAt(t *testing.T) {
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
	   redis> SET mykey "Hello"
	   OK
	*/
	key := "mykey"
	value := "Hello"
	expected := true
	got, err := redisPoolObj_expire.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> PEXPIREAT mykey 1609403601005
	   (integer) 1
	*/
	milliseconds_timestamp := time.Now().AddDate(0, 0, 1).Unix() * 1000
	redisPoolObj_expire.PExpireAt(key, milliseconds_timestamp)

	/*
	   redis> TTL mykey           # TTL 返回秒
	   (integer) 223157079

	   redis> PTTL mykey          # PTTL 返回毫秒
	   (integer) 223157079318
	*/
	expected2 := int64(86400 - 10)
	got2, exist, _, err := redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}

	expected2 = int64(86400000 - 1000)
	got2, exist, _, err = redisPoolObj_expire.PTTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}
}

func TestPTTL(t *testing.T) {
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
	   # 不存在的 key

	   redis> PTTL key
	   (integer) -2
	*/
	key := "key"
	expected := false
	_, got, _, err := redisPoolObj_expire.PTTL(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	/*
	   # key 存在，但没有设置剩余生存时间

	   redis> SET key value
	   OK

	   redis> PTTL key
	   (integer) -1
	*/
	value := "value"
	redisPoolObj_expire.Set(key, value, "", 0, "")
	deleteKeys = append(deleteKeys, key)

	expected = true
	_, _, got, err = redisPoolObj_expire.TTL(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	/*
	   # 有剩余生存时间的 key

	   redis> PEXPIRE key 10086
	   (integer) 1

	   redis> PTTL key
	   (integer) 6179
	*/
	milliseconds := int64(10086)
	expected = true
	got, err = redisPoolObj_expire.PExpire(key, milliseconds)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but got %t", expected, got)
		return
	}

	expected2 := milliseconds - 500
	got2, exist, _, err := redisPoolObj_expire.PTTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	if got2 <= expected2 {
		t.Errorf("Expected to get a number bigger than %d, but now get %d.", expected2, got2)
		return
	}
}
