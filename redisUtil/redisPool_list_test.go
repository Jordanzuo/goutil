package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj_list *RedisPool
)

func init() {
	redisPoolObj_list = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestLPush(t *testing.T) {
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
		# 加入单个元素

		redis> LPUSH languages python
		(integer) 1
	*/
	key := "languages"
	value := "python"
	expected := 1
	got, err := redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		# 加入重复元素

		redis> LPUSH languages python
		(integer) 2
	*/
	value = "python"
	expected = 2
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LRANGE languages 0 -1     # 列表允许重复元素
		1) "python"
		2) "python"
	*/
	expected2 := []string{"python", "python"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		# 加入多个元素

		redis> LPUSH mylist a b c
		(integer) 3
	*/
	key3 := "mylist"
	expected3 := 3
	got3, err := redisPoolObj_list.LPush(key3, "a", "b", "c")
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected length %d, but got %d", expected3, got3)
		return
	}
	deleteKeys = append(deleteKeys, key3)

	/*
		redis> LRANGE mylist 0 -1
		1) "c"
		2) "b"
		3) "a"
	*/
	expected4 := []string{"c", "b", "a"}
	got4_interface, err := redisPoolObj_list.LRange(key3, 0, -1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_list.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestLPushX(t *testing.T) {
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
		# 对空列表执行 LPUSHX

		redis> LLEN greet                       # greet 是一个空列表
		(integer) 0
	*/
	key := "greet"
	expected := 0
	got, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> LPUSHX greet "hello"             # 尝试 LPUSHX，失败，因为列表为空
		(integer) 0
	*/
	value := "hello"
	expected = 0
	got, err = redisPoolObj_list.LPushX(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		# 对非空列表执行 LPUSHX

		redis> LPUSH greet "hello"              # 先用 LPUSH 创建一个有一个元素的列表
		(integer) 1
	*/
	expected = 1
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LPUSHX greet "good morning"      # 这次 LPUSHX 执行成功
		(integer) 2
	*/
	value = "good morning"
	expected = 2
	got, err = redisPoolObj_list.LPushX(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> LRANGE greet 0 -1
		1) "good morning"
		2) "hello"
	*/
	expected2 := []string{"good morning", "hello"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}
}

func TestRPush(t *testing.T) {
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
	   # 添加单个元素

	   redis> RPUSH languages c
	   (integer) 1
	*/
	key := "languages"
	value := "c"
	expected := 1
	got, err := redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   # 添加重复元素

	   redis> RPUSH languages c
	   (integer) 2
	*/
	key = "languages"
	value = "c"
	expected = 2
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> LRANGE languages 0 -1 # 列表允许重复元素
	   1) "c"
	   2) "c"
	*/
	expected2 := []string{"c", "c"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
	   # 添加多个元素

	   redis> RPUSH mylist a b c
	   (integer) 3
	*/
	key3 := "mylist"
	expected3 := 3
	got3, err := redisPoolObj_list.RPush(key3, "a", "b", "c")
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected length %d, but got %d", expected3, got3)
		return
	}
	deleteKeys = append(deleteKeys, key3)

	/*
	   redis> LRANGE mylist 0 -1
	   1) "a"
	   2) "b"
	   3) "c"
	*/
	expected4 := []string{"a", "b", "c"}
	got4_interface, err := redisPoolObj_list.LRange(key3, 0, -1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_list.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestRPushX(t *testing.T) {
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
	   # key不存在

	   redis> LLEN greet
	   (integer) 0
	*/
	key := "greet"
	expected := 0
	got, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> RPUSHX greet "hello"     # 对不存在的 key 进行 RPUSHX，PUSH 失败。
	   (integer) 0
	*/
	value := "hello"
	expected = 0
	got, err = redisPoolObj_list.RPushX(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   # key 存在且是一个非空列表

	   redis> RPUSH greet "hi"         # 先用 RPUSH 插入一个元素
	   (integer) 1
	*/
	value = "hi"
	expected = 1
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
	   redis> RPUSHX greet "hello"     # greet 现在是一个列表类型，RPUSHX 操作成功。
	   (integer) 2
	*/
	value = "hello"
	expected = 2
	got, err = redisPoolObj_list.RPushX(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> LRANGE greet 0 -1
	   1) "hi"
	   2) "hello"
	*/
	expected2 := []string{"hi", "hello"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}
}

func TestLPop(t *testing.T) {
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
	   redis> LLEN course
	   (integer) 0
	*/
	key := "course"
	expected := 0
	got, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> RPUSH course algorithm001
	   (integer) 1
	*/
	value := "algorithm001"
	expected = 1
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   redis> RPUSH course c++101
	   (integer) 2
	*/
	value = "c++101"
	expected = 2
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> LPOP course  # 移除头元素
	   "algorithm001"
	*/
	expected2 := "algorithm001"
	got2_interface, exist, err := redisPoolObj_list.LPop(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should have item, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_list.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}
}

func TestRPop(t *testing.T) {
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
	   redis> RPUSH mylist "one"
	   (integer) 1
	*/
	key := "mylist"
	value := "one"
	expected := 1
	got, err := redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   redis> RPUSH mylist "two"
	   (integer) 2
	*/
	value = "two"
	expected = 2
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> RPUSH mylist "three"
	   (integer) 3
	*/
	value = "three"
	expected = 3
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
	   redis> RPOP mylist           # 返回被弹出的元素
	   "three"
	*/
	expected2 := "three"
	got2_interface, exist, err := redisPoolObj_list.RPop(key)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s should have item, but now it doesn't.", key)
		return
	}
	got2, err := redisPoolObj_list.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}

	/*
	   redis> LRANGE mylist 0 -1    # 列表剩下的元素
	   1) "one"
	   2) "two"
	*/
	expected3 := []string{"one", "two"}
	got3_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_list.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected3, got3) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3)
		return
	}
}

func TestRPopLPush(t *testing.T) {
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
	   # source 和 destination 不同

	   redis> LRANGE alpha 0 -1         # 查看所有元素
	   1) "a"
	   2) "b"
	   3) "c"
	   4) "d"
	*/
	key := "alpha"
	redisPoolObj_list.RPush(key, "a")
	redisPoolObj_list.RPush(key, "b")
	redisPoolObj_list.RPush(key, "c")
	redisPoolObj_list.RPush(key, "d")
	deleteKeys = append(deleteKeys, key)

	expected := []string{"a", "b", "c", "d"}
	got_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   redis> RPOPLPUSH alpha reciver   # 执行一次 RPOPLPUSH 看看
	   "d"
	*/
	source := "alpha"
	destination := "receiver"
	expected2 := "d"
	got2_interface, err := redisPoolObj_list.RPopLPush(source, destination)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %s, but now got %s", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, source)
	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> LRANGE alpha 0 -1
	   1) "a"
	   2) "b"
	   3) "c"
	*/
	expected3 := []string{"a", "b", "c"}
	got3_interface, err := redisPoolObj_list.LRange(source, 0, -1)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_list.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected3, got3) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3)
		return
	}

	/*
	   redis> LRANGE reciver 0 -1
	   1) "d"
	*/
	expected4 := []string{"d"}
	got4_interface, err := redisPoolObj_list.LRange(destination, 0, -1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_list.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got)
		return
	}

	/*
	   redis> RPOPLPUSH alpha reciver   # 再执行一次，证实 RPOP 和 LPUSH 的位置正确
	   "c"
	*/
	expected5 := "c"
	got5_interface, err := redisPoolObj_list.RPopLPush(source, destination)
	if err != nil {
		t.Fail()
	}
	got5, err := redisPoolObj_list.String(got5_interface)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected to get %s, but now got %s", expected5, got5)
		return
	}

	/*
	   redis> LRANGE alpha 0 -1
	   1) "a"
	   2) "b"
	*/
	expected6 := []string{"a", "b"}
	got6_interface, err := redisPoolObj_list.LRange(source, 0, -1)
	if err != nil {
		t.Fail()
	}
	got6, err := redisPoolObj_list.Strings(got6_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected6, got6) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected6, got6)
		return
	}

	/*
	   redis> LRANGE reciver 0 -1
	   1) "c"
	   2) "d"
	*/
	expected7 := []string{"c", "d"}
	got7_interface, err := redisPoolObj_list.LRange(destination, 0, -1)
	if err != nil {
		t.Fail()
	}
	got7, err := redisPoolObj_list.Strings(got7_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected7, got7) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected7, got7)
		return
	}

	/*
	   # source 和 destination 相同

	   redis> LRANGE number 0 -1
	   1) "1"
	   2) "2"
	   3) "3"
	   4) "4"
	*/
	source = "number"
	destination = "number"

	redisPoolObj_list.RPush(source, "1")
	redisPoolObj_list.RPush(source, "2")
	redisPoolObj_list.RPush(source, "3")
	redisPoolObj_list.RPush(source, "4")

	deleteKeys = append(deleteKeys, source)
	deleteKeys = append(deleteKeys, destination)

	expected8 := []string{"1", "2", "3", "4"}
	got8_interface, err := redisPoolObj_list.LRange(source, 0, -1)
	if err != nil {
		t.Fail()
	}
	got8, err := redisPoolObj_list.Strings(got8_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected8, got8) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected8, got8)
		return
	}

	/*
	   redis> RPOPLPUSH number number
	   "4"
	*/
	expected9 := 4
	got9_interface, err := redisPoolObj_list.RPopLPush(source, destination)
	if err != nil {
		t.Fail()
	}
	got9, err := redisPoolObj_list.Int(got9_interface)
	if err != nil {
		t.Fail()
	}
	if got9 != expected9 {
		t.Errorf("Expected to get %d, but now got %d", expected9, got9)
		return
	}

	/*
	   redis> LRANGE number 0 -1           # 4 被旋转到了表头
	   1) "4"
	   2) "1"
	   3) "2"
	   4) "3"
	*/
	expected10 := []string{"4", "1", "2", "3"}
	got10_interface, err := redisPoolObj_list.LRange(source, 0, -1)
	if err != nil {
		t.Fail()
	}
	got10, err := redisPoolObj_list.Strings(got10_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected10, got10) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected10, got10)
		return
	}

	/*
	   redis> RPOPLPUSH number number
	   "3"
	*/
	expected11 := 3
	got11_interface, err := redisPoolObj_list.RPopLPush(source, destination)
	if err != nil {
		t.Fail()
	}
	got11, err := redisPoolObj_list.Int(got11_interface)
	if err != nil {
		t.Fail()
	}
	if got11 != expected11 {
		t.Errorf("Expected to get %d, but now got %d", expected11, got11)
		return
	}

	/*
	   redis> LRANGE number 0 -1           # 这次是 3 被旋转到了表头
	   1) "3"
	   2) "4"
	   3) "1"
	   4) "2"
	*/
	expected12 := []string{"3", "4", "1", "2"}
	got12_interface, err := redisPoolObj_list.LRange(source, 0, -1)
	if err != nil {
		t.Fail()
	}
	got12, err := redisPoolObj_list.Strings(got12_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected12, got12) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected12, got12)
		return
	}
}

func TestLRem(t *testing.T) {
	deleteKeys := make([]string, 0, 8)
	// defer func() {
	// 	// Delete the test keys
	// 	distinctKeyList := getDistinctKeyList(deleteKeys)
	// 	count, err := redisPoolObj_list.Del(distinctKeyList...)
	// 	if err != nil {
	// 		t.Fail()
	// 	}
	// 	if count != len(distinctKeyList) {
	// 		t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
	// 		return
	// 	}
	// }()

	/*
		# 先创建一个表，内容排列是
		# morning hello morning helllo morning

		redis> LPUSH greet "morning"
		(integer) 1
		redis> LPUSH greet "hello"
		(integer) 2
		redis> LPUSH greet "morning"
		(integer) 3
		redis> LPUSH greet "hello"
		(integer) 4
		redis> LPUSH greet "morning"
		(integer) 5
	*/
	key := "greet"
	value := "morning"
	expected := 1
	got, err := redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	value = "hello"
	expected = 2
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	value = "morning"
	expected = 3
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	value = "hello"
	expected = 4
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	value = "morning"
	expected = 5
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> LRANGE greet 0 4         # 查看所有元素
		1) "morning"
		2) "hello"
		3) "morning"
		4) "hello"
		5) "morning"
	*/
	expected2 := []string{"morning", "hello", "morning", "hello", "morning"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		redis> LREM greet 2 morning     # 移除从表头到表尾，最先发现的两个 morning
		(integer) 2                     # 两个元素被移除
	*/
	value3 := "morning"
	expected3 := 2
	got3, err := redisPoolObj_list.LRem(key, 2, value3)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d", expected3, got3)
		return
	}

	/*
		redis> LLEN greet               # 还剩 3 个元素
		(integer) 3
	*/
	expected4 := 3
	got4, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %d, but now got %d", expected4, got4)
		return
	}

	/*
		redis> LRANGE greet 0 2
		1) "hello"
		2) "hello"
		3) "morning"
	*/
	expected5 := []string{"hello", "hello", "morning"}
	got5_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got5, err := redisPoolObj_list.Strings(got5_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected5, got5) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected5, got5)
		return
	}

	/*
		redis> LREM greet -1 morning    # 移除从表尾到表头，第一个 morning
		(integer) 1
	*/

	value6 := "morning"
	expected6 := 1
	got6, err := redisPoolObj_list.LRem(key, -1, value6)
	if err != nil {
		t.Fail()
	}
	if got6 != expected6 {
		t.Errorf("Expected to get %d, but now got %d", expected6, got6)
		return
	}

	/*
		redis> LLEN greet               # 剩下两个元素
		(integer) 2
	*/
	expected7 := 2
	got7, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got7 != expected7 {
		t.Errorf("Expected to get %d, but now got %d", expected7, got7)
		return
	}

	/*
		redis> LRANGE greet 0 1
		1) "hello"
		2) "hello"
	*/
	expected8 := []string{"hello", "hello"}
	got8_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got8, err := redisPoolObj_list.Strings(got8_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected8, got8) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected8, got8)
		return
	}

	/*
		redis> LREM greet 0 hello      # 移除表中所有 hello
		(integer) 2                    # 两个 hello 被移除
	*/
	value9 := "hello"
	expected9 := 2
	got9, err := redisPoolObj_list.LRem(key, 0, value9)
	if err != nil {
		t.Fail()
	}
	if got9 != expected9 {
		t.Errorf("Expected to get %d, but now got %d", expected9, got9)
		return
	}

	/*
		redis> LLEN greet
		(integer) 0
	*/
	expected10 := 0
	got10, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got10 != expected10 {
		t.Errorf("Expected to get %d, but now got %d", expected10, got10)
		return
	}

	expected11 := false
	got11, err := redisPoolObj_list.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got11 != expected11 {
		t.Errorf("Expected to get %t, but now got %t", expected11, got11)
		return
	}
}

func TestLLen(t *testing.T) {
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
		# 空列表

		redis> LLEN job
		(integer) 0
	*/
	key := "job"
	expected := 0
	got, err := redisPoolObj_list.LLen(key)
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		# 非空列表

		redis> LPUSH job "cook food"
		(integer) 1
	*/
	value := "cook food"
	expected = 1
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LPUSH job "have lunch"
		(integer) 2
	*/
	value = "have lunch"
	expected = 2
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LLEN job
		(integer) 2
	*/
	expected = 2
	got, err = redisPoolObj_list.LLen(key)
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
}

func TestLIndex(t *testing.T) {
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
		redis> LPUSH mylist "World"
		(integer) 1
	*/
	key := "mylist"
	value := "World"
	expected := 1
	got, err := redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LPUSH mylist "Hello"
		(integer) 2
	*/
	value = "Hello"
	expected = 2
	got, err = redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}

	/*
		redis> LINDEX mylist 0
		"Hello"
	*/
	expected2 := "Hello"
	index := 0
	got2_interface, exist, err := redisPoolObj_list.LIndex(key, index)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, index:%d should exist, but now it doesn't.", key, index)
		return
	}
	got2, err := redisPoolObj_list.String(got2_interface)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected length %s, but got %s", expected2, got2)
		return
	}

	/*
		redis> LINDEX mylist -1
		"World"
	*/
	expected3 := "World"
	index = -1
	got3_interface, exist, err := redisPoolObj_list.LIndex(key, index)
	if err != nil {
		t.Fail()
	}
	if !exist {
		t.Errorf("The key:%s, index:%d should exist, but now it doesn't.", key, index)
		return
	}
	got3, err := redisPoolObj_list.String(got3_interface)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected length %s, but got %s", expected3, got3)
		return
	}

	/*
		redis> LINDEX mylist 3        # index不在 mylist 的区间范围内
		(nil)
	*/
	index = 3
	_, exist, err = redisPoolObj_list.LIndex(key, index)
	if err != nil {
		t.Fail()
	}
	if exist {
		t.Errorf("The key:%s, index:%d should not exist, but now it does.", key, index)
		return
	}
}

func TestLInsert(t *testing.T) {
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
		redis> RPUSH mylist "Hello"
		(integer) 1

		redis> RPUSH mylist "World"
		(integer) 2
	*/
	key := "mylist"
	value := "Hello"
	expected := 1
	got, err := redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}

	value = "World"
	expected = 2
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> LINSERT mylist BEFORE "World" "There"
		(integer) 3
	*/
	pivot := "World"
	value = "There"
	expected = 3
	got, err = redisPoolObj_list.LInsert(key, "BEFORE", pivot, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}

	/*
		redis> LRANGE mylist 0 -1
		1) "Hello"
		2) "There"
		3) "World"
	*/
	expected2 := []string{"Hello", "There", "World"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		# 对一个非空列表插入，查找一个不存在的 pivot

		redis> LINSERT mylist BEFORE "go" "let's"
		(integer) -1                                    # 失败
	*/
	pivot = "go"
	value = "let's"
	expected = -1
	got, err = redisPoolObj_list.LInsert(key, "BEFORE", pivot, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %d, but got %d", expected, got)
		return
	}

	/*
		# 对一个空列表执行 LINSERT 命令

		redis> EXISTS fake_list
		(integer) 0
	*/
	key3 := "fake_list"
	expected3 := false
	got3, err := redisPoolObj_list.Exists(key3)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected length %t, but got %t", expected3, got3)
		return
	}

	/*
		redis> LINSERT fake_list BEFORE "nono" "gogogog"
		(integer) 0                                      # 失败
	*/
	pivot = "nono"
	value = "gogogog"
	expected4 := 0
	got4, err := redisPoolObj_list.LInsert(key3, "BEFORE", pivot, value)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected to get %d, but got %d", expected4, got4)
		return
	}
}

func TestLSet(t *testing.T) {
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
	   # 对空列表(key 不存在)进行 LSET

	   redis> EXISTS list
	   (integer) 0
	*/
	key := "list"
	expected := false
	got, err := redisPoolObj_list.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected length %t, but got %t", expected, got)
		return
	}

	/*
	   redis> LSET list 0 item
	   (error) ERR no such key
	*/
	index := 0
	value := "item"
	err = redisPoolObj_list.LSet(key, index, value)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	/*
	   # 对非空列表进行 LSET

	   redis> LPUSH job "cook food"
	   (integer) 1
	*/
	key = "job"
	value = "cook food"
	expected2 := 1
	got2, err := redisPoolObj_list.LPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected length %d, but got %d", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   redis> LRANGE job 0 0
	   1) "cook food"
	*/
	expected3 := []string{"cook food"}
	got3_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_list.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected3, got3) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3)
		return
	}

	/*
	   redis> LSET job 0 "play game"
	   OK
	*/
	index = 0
	value = "play game"
	err = redisPoolObj_list.LSet(key, index, value)
	if err != nil {
		t.Fail()
	}

	/*
	   redis> LRANGE job  0 0
	   1) "play game"
	*/
	expected4 := []string{"play game"}
	got4_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_list.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   # index 超出范围

	   redis> LLEN list                    # 列表长度为 1
	   (integer) 1
	*/
	key = "list"
	redisPoolObj_list.RPush(key, "init")

	expected5 := 1
	got5, err := redisPoolObj_list.LLen(key)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected length %d, but got %d", expected5, got5)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   redis> LSET list 3 'out of range'
	   (error) ERR index out of range
	*/
	index = 3
	value = "out of range"
	err = redisPoolObj_list.LSet(key, index, value)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}
}

func TestLRange(t *testing.T) {
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
		redis> RPUSH fp-language lisp
		(integer) 1
	*/
	key := "fp-language"
	value := "lisp"
	expected := 1
	got, err := redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}
	deleteKeys = append(deleteKeys, key)

	/*
		redis> LRANGE fp-language 0 0
		1) "lisp"
	*/
	expected2 := []string{"lisp"}
	got2_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_list.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		redis> RPUSH fp-language scheme
		(integer) 2
	*/
	value = "scheme"
	expected = 2
	got, err = redisPoolObj_list.RPush(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d", expected, got)
		return
	}

	/*
		redis> LRANGE fp-language 0 1
		1) "lisp"
		2) "scheme"
	*/
	expected3 := []string{"lisp", "scheme"}
	got3_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_list.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected3, got3) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3)
		return
	}
}

func TestLTrim(t *testing.T) {
	// deleteKeys := make([]string, 0, 8)
	// defer func() {
	// 	// Delete the test keys
	// 	distinctKeyList := getDistinctKeyList(deleteKeys)
	// 	count, err := redisPoolObj_list.Del(distinctKeyList...)
	// 	if err != nil {
	// 		t.Fail()
	// 	}
	// 	if count != len(distinctKeyList) {
	// 		t.Errorf("Expected to get %d, but now got %d", len(distinctKeyList), count)
	// 		return
	// 	}
	// }()

	/*
	   # 情况 1： 常见情况， start 和 stop 都在列表的索引范围之内
	*/
	key := "alpha"
	redisPoolObj_list.RPush(key, "h", "e", "l", "l", "o")

	/*
	   redis> LRANGE alpha 0 -1       # alpha 是一个包含 5 个字符串的列表
	   1) "h"
	   2) "e"
	   3) "l"
	   4) "l"
	   5) "o"
	*/
	expected := []string{"h", "e", "l", "l", "o"}
	got_interface, err := redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err := redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   redis> LTRIM alpha 1 -1        # 删除 alpha 列表索引为 0 的元素
	   OK

	   redis> LRANGE alpha 0 -1       # "h" 被删除了
	   1) "e"
	   2) "l"
	   3) "l"
	   4) "o"
	*/
	start := 1
	stop := -1
	err = redisPoolObj_list.LTrim(key, start, stop)
	if err != nil {
		t.Fail()
	}

	expected = []string{"e", "l", "l", "o"}
	got_interface, err = redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   # 情况 2： stop 比列表的最大下标还要大

	   redis> LTRIM alpha 1 10086     # 保留 alpha 列表索引 1 至索引 10086 上的元素
	   OK

	   redis> LRANGE alpha 0 -1       # 只有索引 0 上的元素 "e" 被删除了，其他元素还在
	   1) "l"
	   2) "l"
	   3) "o"
	*/
	start = 1
	stop = 10086
	err = redisPoolObj_list.LTrim(key, start, stop)
	if err != nil {
		t.Fail()
	}

	expected = []string{"l", "l", "o"}
	got_interface, err = redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   # 情况 3： start 和 stop 都比列表的最大下标要大，并且 start < stop

	   redis> LTRIM alpha 10086 123321
	   OK

	   redis> LRANGE alpha 0 -1        # 列表被清空
	   (empty list or set)
	*/
	start = 10086
	stop = 123321
	err = redisPoolObj_list.LTrim(key, start, stop)
	if err != nil {
		t.Fail()
	}

	expected = []string{}
	got_interface, err = redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	/*
	   # 情况 4： start 和 stop 都比列表的最大下标要大，并且 start > stop

	   redis> RPUSH new-alpha "h" "e" "l" "l" "o"     # 重新建立一个新列表
	   (integer) 5

	   redis> LRANGE new-alpha 0 -1
	   1) "h"
	   2) "e"
	   3) "l"
	   4) "l"
	   5) "o"

	   redis> LTRIM new-alpha 123321 10086    # 执行 LTRIM
	   OK

	   redis> LRANGE new-alpha 0 -1           # 同样被清空
	   (empty list or set)
	*/
	key = "new-alpha"
	redisPoolObj_list.RPush(key, "h", "e", "l", "l", "o")
	expected = []string{"h", "e", "l", "l", "o"}
	got_interface, err = redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}

	start = 123321
	stop = 10086
	err = redisPoolObj_list.LTrim(key, start, stop)
	if err != nil {
		t.Fail()
	}

	expected = []string{}
	got_interface, err = redisPoolObj_list.LRange(key, 0, -1)
	if err != nil {
		t.Fail()
	}
	got, err = redisPoolObj_list.Strings(got_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoOrderedSliceEqual(expected, got) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected, got)
		return
	}
}
