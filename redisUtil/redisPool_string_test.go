package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj_string *RedisPool
)

func init() {
	redisPoolObj_string = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestSet(t *testing.T) {
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
		对不存在的键进行设置：

		redis> SET key "value"
		OK
	*/

	key := "key"
	value := "value"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET key
	   "value"
	*/
	got_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err := redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   对已存在的键进行设置：

	   redis> SET key "new-value"
	   OK
	*/
	value = "new-value"
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET key
	   "new-value"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   使用 EX 选项：

	   redis> SET key-with-expire-time "hello" EX 10086
	   OK
	*/
	key = "key-with-expire-time"
	value = "hello"
	successful, err = redisPoolObj_string.Set(key, value, "EX", 10086, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET key-with-expire-time
	   "hello"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", "key-with-expire-time")
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> TTL key-with-expire-time
	   (integer) 10069
	*/
	ttl, exist, _, err := redisPoolObj_string.TTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't", key)
		return
	}
	if ttl < 10086-5 {
		t.Errorf("The TTL of key:%s is wrong. Now it's %d", key, ttl)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   使用 PX 选项：

	   redis> SET key-with-pexpire-time "moto" PX 123321
	   OK
	*/
	key = "key-with-pexpire-time"
	value = "moto"
	successful, err = redisPoolObj_string.Set(key, value, "PX", 123321, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET key-with-pexpire-time
	   "moto"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> PTTL key-with-pexpire-time
	   (integer) 111939
	*/
	ttl, exist, _, err = redisPoolObj_string.PTTL(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't", key)
		return
	}
	if ttl < 123321-5000 {
		t.Errorf("The TTL of key:%s is wrong. Now it's %d", key, ttl)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   使用 NX 选项：

	   redis> SET not-exists-key "value" NX
	   OK      # 键不存在，设置成功
	*/
	key = "not-exists-key"
	value = "value"
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "NX")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET not-exists-key
	   "value"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SET not-exists-key "new-value" NX
	   (nil)   # 键已经存在，设置失败

	*/
	newValue := "new-value"
	successful, err = redisPoolObj_string.Set(key, newValue, "", 0, "NX")
	if err != nil {
		t.Fail()
	}
	if successful {
		t.Errorf("Seting key:%s should fail, but now it doesn't.", key)
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GEt not-exists-key
	   "value" # 维持原值不变
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}

	/*
	   使用 XX 选项：

	   redis> EXISTS exists-key
	   (integer) 0
	*/
	key = "exists-key"
	value = "value"
	exist, err = redisPoolObj_string.Exists(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("Key:%s should not exist, but now it does exist", key)
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SET exists-key "value" XX
	   (nil)   # 因为键不存在，设置失败
	*/
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "XX")
	if err != nil {
		t.Fail()
	}
	if successful {
		t.Errorf("Seting key:%s should fail, but now it doesn't.", key)
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SET exists-key "value"
	   OK      # 先给键设置一个值
	*/
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Seting key:%s should succeed, but now it doesn't.", key)
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SET exists-key "new-value" XX
	   OK      # 设置新值成功
	*/
	value = "new-value"
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "XX")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Seting key:%s should succeed, but now it doesn't.", key)
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> GET exists-key
	   "new-value"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> SET int-key 5
		OK      # 设置新值成功
	*/
	key2 := "int-key"
	value2 := 5
	successful, err = redisPoolObj_string.Set(key2, value2, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Seting key:%s should succeed, but now it doesn't.", key)
	}
	deleteKeys = append(deleteKeys, key2)

	/*
		redis> GET int-key
		5
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key2)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_string.Int(got_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != value2 {
		t.Errorf("Expected to get %d, but now got %d.", value2, got2)
		return
	}
}

func TestGet(t *testing.T) {
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
		对不存在的键 key 或是字符串类型的键 key 执行 GET 命令：

		redis> GET db
		(nil)
	*/
	key := "db"
	_, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> SET db redis
		OK
	*/
	value := "redis"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET db
		"redis"
	*/
	got_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err := redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value {
		t.Errorf("Expected to get %s, but now got %s.", value, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		对不是字符串类型的键 key 执行 GET 命令：

		redis> DEL db
		(integer) 1
	*/
	count, err := redisPoolObj_string.Del(key)
	if err != nil {
		t.Fail()
	}
	if count != 1 {
		t.Errorf("Expected to get 1, but now got %d", count)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LPUSH db redis mongodb mysql
		(integer) 3
	*/
	_, err = redisPoolObj_string.LPush(key, "redis", "mongodb", "mysql")
	if err != nil {
		t.Fail()
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET db
		(error) ERR Operation against a key holding the wrong kind of value
	*/
	_, exist, err = redisPoolObj_string.Get(key)
	if err == nil {
		t.Fail()
	}
	deleteKeys = append(deleteKeys, key)
}

func TestGetSet(t *testing.T) {
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
		redis> GETSET db mongodb    # 没有旧值，返回 nil
		(nil)
	*/
	key := "db"
	value1 := "mongodb"
	got, exist, err := redisPoolObj_string.GetSet(key, value1)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("There should be no old value, but now there is.")
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET db
		"mongodb"
	*/
	got_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value1 {
		t.Errorf("Expected to get %s, but now got %s.", value1, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETSET db redis      # 返回旧值 mongodb
		"mongodb"
	*/
	value2 := "redis"
	got_interface, exist, err = redisPoolObj_string.GetSet(key, value2)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value1 {
		t.Errorf("Expected to get %s, but now got %s.", value1, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET db
		"redis"
	*/
	got_interface, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err = redisPoolObj_string.String(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != value2 {
		t.Errorf("Expected to get %s, but now got %s.", value2, got)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestStrLen(t *testing.T) {
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
		获取字符串值的长度：

		redis> SET mykey "Hello world"
		OK
	*/
	key := "mykey"
	value := "Hello world"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> STRLEN mykey
		(integer) 11
		不存在的键的长度为 0 ：
	*/
	expectedLength := 11
	got, err := redisPoolObj_string.StrLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> STRLEN nonexisting
		(integer) 0
	*/
	key = "nonexisting"
	expectedLength = 0
	got, err = redisPoolObj_string.StrLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, got)
		return
	}
}

func TestAppend(t *testing.T) {
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
		示例代码
		对不存在的 key 执行 APPEND ：

		redis> EXISTS myphone               # 确保 myphone 不存在
		(integer) 0
	*/
	key := "myphone"
	value := "value"
	exist, err := redisPoolObj_string.Exists(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("Key:%s should not exist, but now it does exist", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> APPEND myphone "nokia"       # 对不存在的 key 进行 APPEND ，等同于 SET myphone "nokia"
		(integer) 5                         # 字符长度
		对已存在的字符串进行 APPEND ：
	*/
	value = "nokia"
	expected := 5
	got, err := redisPoolObj_string.Append(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> APPEND myphone " - 1110"     # 长度从 5 个字符增加到 12 个字符
		(integer) 12
	*/
	value = " - 1110"
	expected = 12
	got, err = redisPoolObj_string.Append(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET myphone
		"nokia - 1110"
	*/
	expected2 := "nokia - 1110"
	got2_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("Key:%s should exist, but now it doesn't exist", key)
		return
	}
	got2, err := redisPoolObj_string.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestSetRange(t *testing.T) {
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
		对非空字符串执行 SETRANGE 命令：

		redis> SET greeting "hello world"
		OK
	*/
	key := "greeting"
	value := "hello world"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> SETRANGE greeting 6 "Redis"
		(integer) 11
	*/
	offset := 6
	value = "Redis"
	expected := 11
	got, err := redisPoolObj_string.SetRange(key, offset, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET greeting
		"hello Redis"
	*/
	expected2 := "hello Redis"
	got2_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_string.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		对空字符串/不存在的键执行 SETRANGE 命令：

		redis> EXISTS empty_string
		(integer) 0

		redis> SETRANGE empty_string 5 "Redis!"   # 对不存在的 key 使用 SETRANGE
		(integer) 11

		redis> GET empty_string                   # 空白处被"\x00"填充
		"\x00\x00\x00\x00\x00Redis!"
	*/
	key = "empty_string"
	_, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	offset = 5
	value = "Redis!"
	expected3 := 11
	got3, err := redisPoolObj_string.SetRange(key, offset, value)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}
	deleteKeys = append(deleteKeys, key)

	expected4 := "\x00\x00\x00\x00\x00Redis!"
	got4_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got4, err := redisPoolObj_string.String(got4_interface)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %s, but now got %s", expected4, got4)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestGetRange(t *testing.T) {
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
		redis> SET greeting "hello, my friend"
		OK
	*/
	key := "greeting"
	value := "hello, my friend"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETRANGE greeting 0 4          # 返回索引0-4的字符，包括4。
		"hello"
	*/
	start, end := 0, 4
	expected := "hello"
	got, err := redisPoolObj_string.GetRange(key, start, end)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETRANGE greeting -1 -5        # 不支持回绕操作
		""
	*/
	start, end = -1, -5
	expected = ""
	got, err = redisPoolObj_string.GetRange(key, start, end)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETRANGE greeting -3 -1        # 负数索引
		"end"
	*/
	start, end = -3, -1
	expected = "end"
	got, err = redisPoolObj_string.GetRange(key, start, end)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETRANGE greeting 0 -1         # 从第一个到最后一个
		"hello, my friend"
	*/
	start, end = 0, -1
	expected = "hello, my friend"
	got, err = redisPoolObj_string.GetRange(key, start, end)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GETRANGE greeting 0 1008611    # 值域范围不超过实际字符串，超过部分自动被符略
		"hello, my friend"
	*/
	start, end = 0, 1008611
	expected = "hello, my friend"
	got, err = redisPoolObj_string.GetRange(key, start, end)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestIncr(t *testing.T) {
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
		redis> SET page_view 20
		OK
	*/
	key := "page_view"
	value := "20"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> INCR page_view
		(integer) 21
	*/
	expected := int64(21)
	got, err := redisPoolObj_string.Incr(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET page_view    # 数字值在 Redis 中以字符串的形式保存
		"21"
	*/
	expected2 := 21
	got2_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_string.Int(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestIncrBy(t *testing.T) {
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
		键存在，并且值为数字：

		redis> SET rank 50
		OK

		redis> INCRBY rank 20
		(integer) 70

		redis> GET rank
		"70"
	*/
	key := "rank"
	value := "50"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	increment := int64(20)
	expected := int64(70)
	got, err := redisPoolObj_string.IncrBy(key, increment)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	expected2 := 70
	got2_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_string.Int(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		键不存在：

		redis> EXISTS counter
		(integer) 0
	*/
	key = "counter"
	_, exist, err = redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> INCRBY counter 30
		(integer) 30
	*/
	increment = int64(30)
	expected3 := int64(30)
	got3, err := redisPoolObj_string.IncrBy(key, 30)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET counter
		"30"
	*/
	expected4 := 30
	got4_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got4, err := redisPoolObj_string.Int(got4_interface)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %d, but now got %d", expected4, got4)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		键存在，但值无法被解释为数字：

		redis> SET book "long long ago..."
		OK

		redis> INCRBY book 200
		(error) ERR value is not an integer or out of range
	*/
	// This feature can't be tested, because the IncrBy function needs an int64 parameter.
}

func TestIncrByFloat(t *testing.T) {
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
		redis> GET decimal
		"3.0"
	*/
	key := "decimal"
	value := "3.0"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	expected := 3.0
	got_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got, err := redisPoolObj_string.Float64(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %f, but now got %f", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> INCRBYFLOAT decimal 2.56
		"5.56"
	*/
	increment := 2.56
	expected2 := 5.56
	got2, err := redisPoolObj_string.IncrByFloat(key, increment)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %f, but now got %f", expected2, got2)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> GET decimal
		"5.56"
	*/
	expected3 := 5.56
	got3_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got3, err := redisPoolObj_string.Float64(got3_interface)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %f, but now got %f", expected3, got3)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestDecrBy(t *testing.T) {
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
		对已经存在的键执行 DECRBY 命令：

		redis> SET count 100
		OK
	*/
	key := "count"
	value := 100
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> DECRBY count 20
		(integer) 80
		对不存在的键执行 DECRBY 命令：
	*/
	expected := int64(80)
	decrement := int64(20)
	got_interface, err := redisPoolObj_string.DecrBy(key, decrement)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_string.Int64(got_interface)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> EXISTS pages
		(integer) 0
	*/
	key = "pages"
	exist, err := redisPoolObj_string.Exists(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> DECRBY pages 10
		(integer) -10
	*/
	expected = int64(-10)
	decrement = int64(10)
	got, err = redisPoolObj_string.DecrBy(key, decrement)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)
}

func TestMSet(t *testing.T) {
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
		同时对多个键进行设置：

		redis> MSET date "2012.3.30" time "11:00 a.m." weather "sunny"
		OK
	*/
	key_value_map := make(map[string]interface{})
	key_value_map["date"] = "2012.3.30"
	key_value_map["time"] = "11:00 a.m."
	key_value_map["weather"] = "sunny"
	err := redisPoolObj_string.MSet(key_value_map)
	if err != nil {
		t.Fail()
	}

	/*
		redis> MGET date time weather
		1) "2012.3.30"
		2) "11:00 a.m."
		3) "sunny"
	*/
	keyList := make([]string, 0, len(key_value_map))
	expected := make([]string, 0, len(key_value_map))
	for k, v := range key_value_map {
		keyList = append(keyList, k)
		if v_str, ok := v.(string); ok {
			expected = append(expected, v_str)
		}
	}
	got_interface, err := redisPoolObj_string.MGet(keyList)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_string.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, keyList...)

	/*
		覆盖已有的值：

		redis> MGET k1 k2
		1) "hello"
		2) "world"
	*/
	key_value_map = make(map[string]interface{})
	key_value_map["k1"] = "hello"
	key_value_map["k2"] = "world"
	err = redisPoolObj_string.MSet(key_value_map)
	if err != nil {
		t.Fail()
	}

	/*

		redis> MSET k1 "good" k2 "bye"
		OK
	*/
	key_value_map["k1"] = "good"
	key_value_map["k2"] = "bye"
	err = redisPoolObj_string.MSet(key_value_map)
	if err != nil {
		t.Fail()
	}

	/*
		redis> MGET k1 k2
		1) "good"
		2) "bye"
	*/
	keyList = make([]string, 0, len(key_value_map))
	expected = make([]string, 0, len(key_value_map))
	for k, v := range key_value_map {
		keyList = append(keyList, k)
		if v_str, ok := v.(string); ok {
			expected = append(expected, v_str)
		}
	}

	got_interface, err = redisPoolObj_string.MGet(keyList)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_string.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, keyList...)
}

func TestMSetNX(t *testing.T) {
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
	   对不存在的键执行 MSETNX 命令：

	   redis> MSETNX rmdbs "MySQL" nosql "MongoDB" key-value-store "redis"
	   (integer) 1
	*/
	key_value_map := make(map[string]interface{})
	key_value_map["rmdbs"] = "MySQL"
	key_value_map["nosql"] = "MongoDB"
	key_value_map["key-value-store"] = "redis"
	successful, err := redisPoolObj_string.MSetNX(key_value_map)
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("It should be successful, but now it's not.")
		return
	}

	/*

	   redis> MGET rmdbs nosql key-value-store
	   1) "MySQL"
	   2) "MongoDB"
	   3) "redis"
	   对某个已经存在的键进行设置：
	*/
	keyList := make([]string, 0, len(key_value_map))
	expected := make([]string, 0, len(key_value_map))
	for k, v := range key_value_map {
		keyList = append(keyList, k)
		if v_str, ok := v.(string); ok {
			expected = append(expected, v_str)
		}
	}
	got_interface, err := redisPoolObj_string.MGet(keyList)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_string.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, keyList...)

	/*
	   redis> MSETNX rmdbs "Sqlite" language "python"  # rmdbs 键已经存在，操作失败
	   (integer) 0
	*/
	key_value_map = make(map[string]interface{})
	key_value_map["rmdbs"] = "Sqlite"
	key_value_map["language"] = "python"
	successful, err = redisPoolObj_string.MSetNX(key_value_map)
	if err != nil {
		t.Fail()
	}
	if successful {
		t.Errorf("It should be not successful, but now it is.")
		return
	}

	/*
	   redis> EXISTS language                          # 因为 MSETNX 命令没有成功执行
	   (integer) 0                                     # 所以 language 键没有被设置
	*/
	key := "language"
	exist, err := redisPoolObj_string.Exists(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s should not exist, but now it does.", key)
		return
	}

	/*
	   redis> GET rmdbs                                # rmdbs 键也没有被修改
	   "MySQL"
	*/
	key = "rmdbs"
	expected2 := "MySQL"
	got2_interface, exist, err := redisPoolObj_string.Get(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should exist, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_string.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %v, but got %v", expected2, got2)
		return
	}
}

func TestMGet(t *testing.T) {
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
	   redis> SET redis redis.com
	   OK
	*/
	key := "redis"
	value := "redis.com"
	successful, err := redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SET mongodb mongodb.org
	   OK
	*/
	key = "mongodb"
	value = "mongodb.org"
	successful, err = redisPoolObj_string.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set key:%s should be successful, but now it's not.", key)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> MGET redis mongodb
	   1) "redis.com"
	   2) "mongodb.org"
	*/
	keys := []string{"redis", "mongodb"}
	expected := []string{"redis.com", "mongodb.org"}
	got_interface, err := redisPoolObj_string.MGet(keys)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_string.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
		return
	}

	/*
	   redis> MGET redis mongodb mysql     # 不存在的 mysql 返回 nil
	   1) "redis.com"
	   2) "mongodb.org"
	   3) (nil)
	*/
	keys = []string{"redis", "mongodb", "mysql"}
	expected = []string{"redis.com", "mongodb.org", ""}
	got_interface, err = redisPoolObj_string.MGet(keys)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_string.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(got, expected) == false {
		t.Errorf("Expected to get %v, but got %v", expected, got)
		return
	}
}
