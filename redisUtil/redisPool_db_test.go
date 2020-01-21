package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj_db *RedisPool
)

func init() {
	redisPoolObj_db = NewRedisPool("testPool", "10.1.0.21:6379", "redis_pwd", 5, 500, 200, 10*time.Second, 5*time.Second)
}

func TestExists(t *testing.T) {
	/*
		redis> SET db "redis"
		OK
	*/
	key := "db"
	value := "redis"
	successful, err := redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set the key:%s should be successful, but now it's not.", key)
		return
	}

	/*
		redis> EXISTS db
		(integer) 1
	*/
	exist, err := redisPoolObj_db.Exists(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("Set the key:%s should exist, but now it doesn't.", key)
		return
	}

	/*
		redis> DEL db
		(integer) 1
	*/
	expected := 1
	got, err := redisPoolObj_db.Del(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> EXISTS db
		(integer) 0
	*/
	exist, err = redisPoolObj_db.Exists(key)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("Set the key:%s should not exist, but now it does.", key)
		return
	}
}

func TestType(t *testing.T) {
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
		# 字符串

		redis> SET weather "sunny"
		OK

		redis> TYPE weather
		string
	*/
	key := "weather"
	value := "sunny"

	successful, err := redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if !successful {
		t.Errorf("Set the key:%s should be successful, but now it's not.", key)
		return
	}

	expected := "string"
	got, err := redisPoolObj_db.Type(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		# 列表

		redis> LPUSH book_list "programming in scala"
		(integer) 1

		redis> TYPE book_list
		list
	*/
	key = "book_list"
	value = "programming in scala"
	expected2 := 1
	got2, err := redisPoolObj_db.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d", expected2, got2)
		return
	}

	expected = "list"
	got, err = redisPoolObj_db.Type(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		# 集合

		redis> SADD pat "dog"
		(integer) 1

		redis> TYPE pat
		set
	*/

	key = "pat"
	value = "dog"
	expected3 := 1
	got3, err := redisPoolObj_db.SAdd(key, value)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}

	expected = "set"
	got, err = redisPoolObj_db.Type(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)
}

func TestRename(t *testing.T) {
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
	   # key 存在且 newkey 不存在

	   redis> SET message "hello world"
	   OK

	   redis> RENAME message greeting
	   OK

	   redis> EXISTS message               # message 不复存在
	   (integer) 0

	   redis> EXISTS greeting              # greeting 取而代之
	   (integer) 1
	*/
	key := "message"
	value := "hello world"
	expected := true
	got, err := redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	newkey := "greeting"
	err = redisPoolObj_db.Rename(key, newkey)
	if err != nil {
		t.Fail()
	}

	expected = false
	got, err = redisPoolObj_db.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected = true
	got, err = redisPoolObj_db.Exists(newkey)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, newkey)

	/*
	   # 当 key 不存在时，返回错误

	   redis> RENAME fake_key never_exists
	   (error) ERR no such key
	*/
	key = "fake_key"
	newkey = "never_exists"
	err = redisPoolObj_db.Rename(key, newkey)
	if err == nil {
		t.Errorf("There should be one error, but now there isn't.")
		return
	}

	/*
	   # newkey 已存在时， RENAME 会覆盖旧 newkey

	   redis> SET pc "lenovo"
	   OK

	   redis> SET personal_computer "dell"
	   OK

	   redis> RENAME pc personal_computer
	   OK

	   redis> GET pc
	   (nil)

	   redis:1> GET personal_computer      # 原来的值 dell 被覆盖了
	   "lenovo"
	*/
	key = "pc"
	value = "lenovo"
	expected = true
	got, err = redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	key = "personal_computer"
	value = "dell"
	expected = true
	got, err = redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	key = "pc"
	newkey = "personal_computer"
	err = redisPoolObj_db.Rename(key, newkey)
	if err != nil {
		t.Fail()
	}

	expected = false
	got, err = redisPoolObj_db.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected = true
	got, err = redisPoolObj_db.Exists(newkey)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, newkey)
}

func TestRenameNX(t *testing.T) {
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
	   # newkey 不存在，改名成功

	   redis> SET player "MPlyaer"
	   OK

	   redis> EXISTS best_player
	   (integer) 0

	   redis> RENAMENX player best_player
	   (integer) 1
	*/
	key := "player"
	value := "MPlayer"
	expected := true
	got, err := redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	newkey := "best_player"
	expected = false
	got, err = redisPoolObj_db.Exists(newkey)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected = true
	got, err = redisPoolObj_db.RenameNX(key, newkey)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, newkey)

	/*
	   # newkey存在时，失败

	   redis> SET animal "bear"
	   OK

	   redis> SET favorite_animal "butterfly"
	   OK

	   redis> RENAMENX animal favorite_animal
	   (integer) 0

	   redis> get animal
	   "bear"

	   redis> get favorite_animal
	   "butterfly"
	*/
	key = "animal"
	value = "bear"
	expected = true
	got, err = redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	key = "favorite_animal"
	value = "butterfly"
	expected = true
	got, err = redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	expected = false
	got, err = redisPoolObj_db.RenameNX(key, newkey)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	key2 := "animal"
	expected2 := "bear"
	got2_interface, exist2, err := redisPoolObj_db.Get(key2)
	if err != nil {
		t.Fail()
	}
	if !exist2 {
		t.Errorf("The key:%s should exist, but now it doesn't.", key2)
		return
	}
	got2, err := redisPoolObj_db.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}

	key3 := "animal"
	expected3 := "bear"
	got3_interface, exist3, err := redisPoolObj_db.Get(key3)
	if err != nil {
		t.Fail()
	}
	if !exist3 {
		t.Errorf("The key:%s should exist, but now it doesn't.", key3)
		return
	}
	got3, err := redisPoolObj_db.String(got3_interface)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %s, but now got %s", expected3, got3)
		return
	}
}

func TestDel(t *testing.T) {
	/*
	   #  删除单个 key

	   redis> SET name huangz
	   OK

	   redis> DEL name
	   (integer) 1
	*/
	key := "name"
	value := "huangz"
	expected := true
	got, err := redisPoolObj_db.Set(key, value, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected2 := 1
	got2, err := redisPoolObj_db.Del(key)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d", expected2, got2)
		return
	}

	/*
	   # 删除一个不存在的 key

	   redis> EXISTS phone
	   (integer) 0

	   redis> DEL phone # 失败，没有 key 被删除
	   (integer) 0
	*/
	key = "phone"
	expected = false
	got, err = redisPoolObj_db.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected3 := 0
	got3, err := redisPoolObj_db.Del(key)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}

	/*
	   # 同时删除多个 key

	   redis> SET name "redis"
	   OK

	   redis> SET type "key-value store"
	   OK

	   redis> SET website "redis.com"
	   OK

	   redis> DEL name type website
	   (integer) 3
	*/
	key1 := "name"
	value1 := "redis"
	expected = true
	got, err = redisPoolObj_db.Set(key1, value1, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}
	key2 := "type"
	value2 := "key-value store"
	expected = true
	got, err = redisPoolObj_db.Set(key2, value2, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}
	key3 := "website"
	value3 := "redis.com"
	expected = true
	got, err = redisPoolObj_db.Set(key3, value3, "", 0, "")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t", expected, got)
		return
	}

	expected4 := 3
	got4, err := redisPoolObj_db.Del(key1, key2, key3)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %d, but now got %d", expected4, got4)
		return
	}
}

func TestRandomKey(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_list.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
	   # 数据库为空

	   redis> RANDOMKEY
	   (nil)
	*/
	expected := ""
	got, exist, err := redisPoolObj_db.RandomKey()
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("RandomKey doesn't exist, but now it does.")
		return
	}
	if got != expected {
		t.Errorf("Expected to get %s, but now got %s", expected, got)
		return
	}

	/*
	   # 数据库不为空

	   redis> MSET fruit "apple" drink "beer" food "cookies"   # 设置多个 key
	   OK

	   redis> RANDOMKEY
	   "fruit"

	   redis> RANDOMKEY
	   "food"

	   redis> KEYS *    # 查看数据库内所有key，证明 RANDOMKEY 并不删除 key
	   1) "food"
	   2) "drink"
	   3) "fruit"
	*/
	key_value_map := make(map[string]interface{})
	key_value_map["fruit"] = "apple"
	key_value_map["drink"] = "beer"
	key_value_map["food"] = "cookies"
	err = redisPoolObj_db.MSet(key_value_map)
	if err != nil {
		t.Fail()
	}

	got, exist, err = redisPoolObj_db.RandomKey()
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("RandomKey should exist, but now it doesn't.")
		return
	}
	if _, exist = key_value_map[got]; !exist {
		t.Errorf("RandomKey should exist, but now it doesn't.")
		return
	}

	expected2 := make([]string, 0, len(key_value_map))
	for k := range key_value_map {
		expected2 = append(expected2, k)
		deleteKeys = append(deleteKeys, k)
	}
	got2, err := redisPoolObj_db.Keys("*")
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}
}

func TestDBSize(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_list.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
	   redis> DBSIZE
	   (integer) 0

	   redis> SET new_key "hello_moto"     # 增加一个 key 试试
	   OK

	   redis> DBSIZE
	   (integer) 1
	*/
	expected := 0
	got, err := redisPoolObj_db.DBSize()
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	key := "new_key"
	value := "hello_moto"
	redisPoolObj_db.Set(key, value, "", 0, "")

	deleteKeys = append(deleteKeys, key)

	expected = 1
	got, err = redisPoolObj_db.DBSize()
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
}

func TestKeys(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	defer func() {
		// Delete the test keys
		distinctKeyList := getDistinctKeyList(deleteKeys)
		count, err := redisPoolObj_list.Del(distinctKeyList...)
		if err != nil {
			t.Fail()
		}
		if count != len(distinctKeyList) {
			t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
			return
		}
	}()

	/*
	   redis> MSET one 1 two 2 three 3 four 4  # 一次设置 4 个 key
	   OK
	*/
	key_value_map := make(map[string]interface{})
	key_value_map["one"] = "1"
	key_value_map["two"] = "2"
	key_value_map["three"] = "3"
	key_value_map["four"] = "4"
	err := redisPoolObj_db.MSet(key_value_map)
	if err != nil {
		t.Fail()
	}

	for k := range key_value_map {
		deleteKeys = append(deleteKeys, k)
	}

	/*
	   redis> KEYS *o*
	   1) "four"
	   2) "two"
	   3) "one"
	*/
	pattern := "*o*"
	expected := []string{"four", "two", "one"}
	got, err := redisPoolObj_db.Keys(pattern)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   redis> KEYS t??
	   1) "two"
	*/
	pattern = "t??"
	expected = []string{"two"}
	got, err = redisPoolObj_db.Keys(pattern)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   redis> KEYS t[w]*
	   1) "two"
	*/
	pattern = "t[w]*"
	expected = []string{"two"}
	got, err = redisPoolObj_db.Keys(pattern)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   redis> KEYS *  # 匹配数据库内所有 key
	   1) "four"
	   2) "three"
	   3) "two"
	   4) "one"
	*/
	pattern = "*"
	expected = []string{"two", "one", "three", "four"}
	got, err = redisPoolObj_db.Keys(pattern)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}
}
