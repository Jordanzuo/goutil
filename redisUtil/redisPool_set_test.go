package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj_set *RedisPool
)

func init() {
	redisPoolObj_set = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestSAdd(t *testing.T) {
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
	   # 添加单个元素

	   redis> SADD bbs "discuz.net"
	   (integer) 1
	*/
	key := "bbs"
	value := "discuz.net"
	expected := 1
	got, err := redisPoolObj_set.SAdd(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   # 添加重复元素

	   redis> SADD bbs "discuz.net"
	   (integer) 0
	*/
	expected = 0
	got, err = redisPoolObj_set.SAdd(key, value)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	/*
	   # 添加多个元素

	   redis> SADD bbs "tianya.cn" "groups.google.com"
	   (integer) 2
	*/
	expected = 2
	got, err = redisPoolObj_set.SAdd(key, "tianya.cn", "groups.google.com")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	/*
	   redis> SMEMBERS bbs
	   1) "discuz.net"
	   2) "groups.google.com"
	   3) "tianya.cn"
	*/
	expected2 := []string{"discuz.net", "groups.google.com", "tianya.cn"}
	got2_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}
}

func TestSIsMember(t *testing.T) {
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

	key := "joe's_movies"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key, "hi, lady", "Fast Five", "2012")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> SMEMBERS joe's_movies
		1) "hi, lady"
		2) "Fast Five"
		3) "2012"
	*/
	expected2 := []string{"hi, lady", "Fast Five", "2012"}
	got2_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		redis> SISMEMBER joe's_movies "bet man"
		(integer) 0
	*/
	expected3 := false
	value := "bet man"
	got3, err := redisPoolObj_set.SIsMember(key, value)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %t, but now got %t", expected3, got3)
		return
	}

	/*
		redis> SISMEMBER joe's_movies "Fast Five"
		(integer) 1
	*/
	expected3 = true
	value = "Fast Five"
	got3, err = redisPoolObj_set.SIsMember(key, value)
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %t, but now got %t", expected3, got3)
		return
	}

}

func TestSPop(t *testing.T) {
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

	key := "db"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key, "MySQL", "MongoDB", "Redis")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		redis> SMEMBERS db
		1) "MySQL"
		2) "MongoDB"
		3) "Redis"
	*/
	expected2 := []string{"MySQL", "MongoDB", "Redis"}
	got2_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
		redis> SPOP db
		"Redis"

		redis> SMEMBERS db
		1) "MySQL"
		2) "MongoDB"
	*/
	expected3 := make(map[string]struct{})
	for _, item := range expected2 {
		expected3[item] = struct{}{}
	}
	got3_interface, err := redisPoolObj_set.SPop(key)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_set.String(got3_interface)
	if err != nil {
		t.Fail()
	}
	if _, exist := expected3[got3]; !exist {
		t.Errorf("Expected to get one of key from %v, but now get %s.", expected3, got3)
		return
	}
	delete(expected3, got3)

	expected4 := make([]string, 0, len(expected3))
	for k := range expected3 {
		expected4 = append(expected4, k)
	}

	got4_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
		redis> SPOP db
		"MySQL"

		redis> SMEMBERS db
		1) "MongoDB"
	*/

	expected5 := make(map[string]struct{})
	for _, item := range expected4 {
		expected5[item] = struct{}{}
	}
	got5_interface, err := redisPoolObj_set.SPop(key)
	if err != nil {
		t.Fail()
	}
	got5, err := redisPoolObj_set.String(got5_interface)
	if err != nil {
		t.Fail()
	}
	if _, exist := expected5[got5]; !exist {
		t.Errorf("Expected to get one of key from %v, but now get %s.", expected5, got5)
		return
	}
	delete(expected5, got5)

	expected6 := make([]string, 0, len(expected5))
	for k := range expected5 {
		expected6 = append(expected6, k)
	}

	got6_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got6, err := redisPoolObj_set.Strings(got6_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected6, got6) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected6, got6)
		return
	}
}

func TestSRandMember(t *testing.T) {
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
		# 添加元素

		redis> SADD fruit apple banana cherry
		(integer) 3
	*/
	key := "fruit"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key, "apple", "banana", "cherry")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		# 只给定 key 参数，返回一个随机元素

		redis> SRANDMEMBER fruit
		"cherry"
	*/
	expected2 := make(map[string]struct{})
	expected2["apple"] = struct{}{}
	expected2["banana"] = struct{}{}
	expected2["cherry"] = struct{}{}

	count := 1
	got2_interface, err := redisPoolObj_set.SRandMember(key, count)
	if err != nil {
		t.Fail()
	}
	got2_slice, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if len(got2_slice) != count {
		t.Errorf("Expected to have %d items, but now get %d.", count, len(got2_slice))
		return
	}
	got2 := got2_slice[0]
	if _, exist := expected2[got2]; !exist {
		t.Errorf("Expected length %s, but got %s", expected2, got2)
		return
	}

	/*
		# 给定 3 为 count 参数，返回 3 个随机元素
		# 每个随机元素都不相同

		redis> SRANDMEMBER fruit 3
		1) "apple"
		2) "banana"
		3) "cherry"
	*/
	count = 3
	expected3 := make([]string, 0, len(expected2))
	for k := range expected2 {
		expected3 = append(expected3, k)
	}
	got3_interface, err := redisPoolObj_set.SRandMember(key, count)
	if err != nil {
		t.Fail()
	}
	got3_slice, err := redisPoolObj_set.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if len(got3_slice) != count {
		t.Errorf("Expected to have %d items, but now get %d.", count, len(got3_slice))
		return
	}
	if isTwoUnorderedSliceEqual(expected3, got3_slice) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3_slice)
		return
	}

	/*
		# 给定 -3 为 count 参数，返回 3 个随机元素
		# 元素可能会重复出现多次

		redis> SRANDMEMBER fruit -3
		1) "banana"
		2) "cherry"
		3) "apple"

		# 如果 count 是整数，且大于等于集合基数，那么返回整个集合

		redis> SRANDMEMBER fruit 10
		1) "apple"
		2) "banana"
		3) "cherry"

		# 如果 count 是负数，且 count 的绝对值大于集合的基数
		# 那么返回的数组的长度为 count 的绝对值

		redis> SRANDMEMBER fruit -10
		1) "banana"
		2) "apple"
		3) "banana"
		4) "cherry"
		5) "apple"
		6) "apple"
		7) "cherry"
		8) "apple"
		9) "apple"
		10) "banana"

		# SRANDMEMBER 并不会修改集合内容

		redis> SMEMBERS fruit
		1) "apple"
		2) "cherry"
		3) "banana"

		# 集合为空时返回 nil 或者空数组

		redis> SRANDMEMBER not-exists
		(nil)

		redis> SRANDMEMBER not-eixsts 10
		(empty list or set)
	*/
}

func TestSRem(t *testing.T) {
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
		# 添加元素

		redis> SADD languages c lisp python ruby
		(integer) 4
	*/
	key := "languages"
	expected := 4
	got, err := redisPoolObj_set.SAdd(key, "c", "lisp", "python", "ruby")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
		# 测试数据

		redis> SMEMBERS languages
		1) "c"
		2) "lisp"
		3) "python"
		4) "ruby"
	*/
	expected2 := make(map[string]struct{})
	expected2["c"] = struct{}{}
	expected2["lisp"] = struct{}{}
	expected2["python"] = struct{}{}
	expected2["ruby"] = struct{}{}

	expected4 := make([]string, 0, len(expected2))
	for k := range expected2 {
		expected4 = append(expected4, k)
	}

	got4_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
		# 移除单个元素

		redis> SREM languages ruby
		(integer) 1
	*/
	value := "ruby"
	expected5 := 1
	got5, err := redisPoolObj_set.SRem(key, value)
	if err != nil {
		t.Fail()
	}
	if got5 != expected5 {
		t.Errorf("Expected length %d, but got %d", expected5, got5)
		return
	}

	/*
		# 移除不存在元素

		redis> SREM languages non-exists-language
		(integer) 0
	*/
	value = "non-exists-language"
	expected6 := 0
	got6, err := redisPoolObj_set.SRem(key, value)
	if err != nil {
		t.Fail()
	}
	if got6 != expected6 {
		t.Errorf("Expected length %d, but got %d", expected6, got6)
		return
	}

	/*
		# 移除多个元素

		redis> SREM languages lisp python c
		(integer) 3

		redis> SMEMBERS languages
		(empty list or set)
	*/
	expected7 := 3
	got7, err := redisPoolObj_set.SRem(key, "c", "lisp", "python")
	if err != nil {
		t.Fail()
	}
	if got7 != expected7 {
		t.Errorf("Expected length %d, but got %d", expected7, got7)
		return
	}

	expected8 := make([]string, 0)
	got8_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got8, err := redisPoolObj_set.Strings(got8_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected8, got8) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected8, got8)
		return
	}
}

func TestSMove(t *testing.T) {
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
		# 添加元素

		redis> SADD languages c lisp python ruby
		(integer) 4
	*/
	source := "songs"
	expected := 2
	got, err := redisPoolObj_set.SAdd(source, "Billie Jean", "Believe Me")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, source)

	/*
	   redis> SMEMBERS songs
	   1) "Billie Jean"
	   2) "Believe Me"
	*/
	expected2 := []string{"Billie Jean", "Believe Me"}
	got2_interface, err := redisPoolObj_set.SMembers(source)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
	   redis> SMEMBERS my_songs
	   (empty list or set)
	*/
	destination := "my_songs"
	expected3 := []string{}
	got3_interface, err := redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got3, err := redisPoolObj_set.Strings(got3_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected3, got3) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected3, got3)
		return
	}

	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> SMOVE songs my_songs "Believe Me"
	   (integer) 1
	*/
	value := "Believe Me"
	expected4 := true
	got4, err := redisPoolObj_set.SMove(source, destination, value)
	if err != nil {
		t.Fail()
	}
	if got4 != expected4 {
		t.Errorf("Expected length %t, but got %t", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS songs
	   1) "Billie Jean"
	*/
	expected5 := []string{"Billie Jean"}
	got5_interface, err := redisPoolObj_set.SMembers(source)
	if err != nil {
		t.Fail()
	}
	got5, err := redisPoolObj_set.Strings(got5_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected5, got5) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected5, got5)
		return
	}

	/*
	   redis> SMEMBERS my_songs
	   1) "Believe Me"
	*/
	expected6 := []string{"Believe Me"}
	got6_interface, err := redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got6, err := redisPoolObj_set.Strings(got6_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected6, got6) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected6, got6)
		return
	}
}

func TestSCard(t *testing.T) {
	/*
	   redis> SADD tool pc printer phone
	   (integer) 3
	*/
	key := "tool"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key, "pc", "printer", "phone")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	/*
	   redis> SCARD tool   # 非空集合
	   (integer) 3
	*/
	expected = 3
	got, err = redisPoolObj_set.SCard(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	/*
	   redis> DEL tool
	   (integer) 1
	*/
	expected = 1
	got, err = redisPoolObj_set.Del(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	/*
	   redis> SCARD tool   # 空集合
	   (integer) 0
	*/
	expected = 0
	got, err = redisPoolObj_set.SCard(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}
}

func TestSMembers(t *testing.T) {
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
	   # key 不存在或集合为空

	   redis> EXISTS not_exists_key
	   (integer) 0
	*/
	key := "not_exists_key"
	expected := false
	got, err := redisPoolObj_set.Exists(key)
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %t, but now got %t.", expected, got)
		return
	}

	/*
	   redis> SMEMBERS not_exists_key
	   (empty list or set)
	*/
	expected2 := make([]string, 0)
	got2_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got2, err := redisPoolObj_set.Strings(got2_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected2, got2) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected2, got2)
		return
	}

	/*
	   # 非空集合

	   redis> SADD language Ruby Python Clojure
	   (integer) 3
	*/
	key = "language"
	expected3 := 3
	got3, err := redisPoolObj_set.SAdd(key, "Clojure", "Python", "Ruby")
	if err != nil {
		t.Fail()
	}
	if got3 != expected3 {
		t.Errorf("Expected to get %d, but now got %d.", expected3, got3)
		return
	}

	deleteKeys = append(deleteKeys, key)

	/*
	   redis> SMEMBERS language
	   1) "Python"
	   2) "Ruby"
	   3) "Clojure"
	*/
	expected4 := []string{"Clojure", "Python", "Ruby"}
	got4_interface, err := redisPoolObj_set.SMembers(key)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSInter(t *testing.T) {
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

	key1 := "group_1"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key1, "LI LEI", "TOM", "JACK")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	key2 := "group_2"
	expected = 2
	got, err = redisPoolObj_set.SAdd(key2, "HAN MEIMEI", "JACK")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	/*
	   redis> SMEMBERS group_1
	   1) "LI LEI"
	   2) "TOM"
	   3) "JACK"
	*/
	expected4 := []string{"LI LEI", "TOM", "JACK"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS group_2
	   1) "HAN MEIMEI"
	   2) "JACK"
	*/
	expected4 = []string{"HAN MEIMEI", "JACK"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SINTER group_1 group_2
	   1) "JACK"
	*/
	expected4 = []string{"JACK"}
	got4_interface, err = redisPoolObj_set.SInter(key1, key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSInterStore(t *testing.T) {
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
	   redis> SMEMBERS songs
	   1) "good bye joe"
	   2) "hello,peter"
	*/
	key1 := "songs"
	expected := 2
	got, err := redisPoolObj_set.SAdd(key1, "good bye joe", "hello,peter")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	expected4 := []string{"good bye joe", "hello,peter"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS my_songs
	   1) "good bye joe"
	   2) "falling"
	*/
	key2 := "my_songs"
	expected = 2
	got, err = redisPoolObj_set.SAdd(key2, "good bye joe", "falling")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	expected4 = []string{"good bye joe", "falling"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SINTERSTORE song_interset songs my_songs
	   (integer) 1
	*/
	destination := "song_interset"
	expected2 := 1
	got2, err := redisPoolObj_set.SInterStore(destination, key1, key2)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d.", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> SMEMBERS song_interset
	   1) "good bye joe"
	*/
	expected4 = []string{"good bye joe"}
	got4_interface, err = redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSUnion(t *testing.T) {
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
	   redis> SMEMBERS songs
	   1) "Billie Jean"
	*/
	key1 := "songs"
	expected := 1
	got, err := redisPoolObj_set.SAdd(key1, "Billie Jean")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	expected4 := []string{"Billie Jean"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS my_songs
	   1) "Believe Me"
	*/
	key2 := "my_songs"
	expected = 1
	got, err = redisPoolObj_set.SAdd(key2, "Believe Me")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	expected4 = []string{"Believe Me"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SUNION songs my_songs
	   1) "Billie Jean"
	   2) "Believe Me"
	*/
	expected4 = []string{"Billie Jean", "Believe Me"}
	got4_interface, err = redisPoolObj_set.SUnion(key1, key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSUnionStore(t *testing.T) {
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
	   redis> SMEMBERS NoSQL
	   1) "MongoDB"
	   2) "Redis"
	*/
	key1 := "NoSQL"
	expected := 2
	got, err := redisPoolObj_set.SAdd(key1, "MongoDB", "Redis")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	expected4 := []string{"MongoDB", "Redis"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS SQL
	   1) "sqlite"
	   2) "MySQL"
	*/
	key2 := "SQL"
	expected = 2
	got, err = redisPoolObj_set.SAdd(key2, "sqlite", "MySQL")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	expected4 = []string{"sqlite", "MySQL"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SUNIONSTORE db NoSQL SQL
	   (integer) 4
	*/
	destination := "db"
	expected2 := 4
	got2, err := redisPoolObj_set.SUnionStore(destination, key1, key2)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d.", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> SMEMBERS db
	   1) "MySQL"
	   2) "sqlite"
	   3) "MongoDB"
	   4) "Redis"
	*/
	expected4 = []string{"MongoDB", "Redis", "sqlite", "MySQL"}
	got4_interface, err = redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSDiff(t *testing.T) {
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
	   redis> SMEMBERS peter's_movies
	   1) "bet man"
	   2) "start war"
	   3) "2012"
	*/
	key1 := "peter's_movies"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key1, "bet man", "start war", "2012")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	expected4 := []string{"bet man", "start war", "2012"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS joe's_movies
	   1) "hi, lady"
	   2) "Fast Five"
	   3) "2012"
	*/
	key2 := "joe's_movies"
	expected = 3
	got, err = redisPoolObj_set.SAdd(key2, "hi, lady", "Fast Five", "2012")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	expected4 = []string{"hi, lady", "Fast Five", "2012"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SDIFF peter's_movies joe's_movies
	   1) "bet man"
	   2) "start war"
	*/
	expected4 = []string{"bet man", "start war"}
	got4_interface, err = redisPoolObj_set.SDiff(key1, key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SDIFF joe's_movies peter's_movies
	   1) "hi, lady"
	   2) "Fast Five"
	*/
	expected4 = []string{"hi, lady", "Fast Five"}
	got4_interface, err = redisPoolObj_set.SDiff(key2, key1)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}

func TestSDiffStore(t *testing.T) {
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
	   redis> SMEMBERS peter's_movies
	   1) "bet man"
	   2) "start war"
	   3) "2012"
	*/
	key1 := "peter's_movies"
	expected := 3
	got, err := redisPoolObj_set.SAdd(key1, "bet man", "start war", "2012")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key1)

	expected4 := []string{"bet man", "start war", "2012"}
	got4_interface, err := redisPoolObj_set.SMembers(key1)
	if err != nil {
		t.Fail()
	}
	got4, err := redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SMEMBERS joe's_movies
	   1) "hi, lady"
	   2) "Fast Five"
	   3) "2012"
	*/
	key2 := "joe's_movies"
	expected = 3
	got, err = redisPoolObj_set.SAdd(key2, "hi, lady", "Fast Five", "2012")
	if err != nil {
		t.Fail()
	}
	if got != expected {
		t.Errorf("Expected to get %d, but now got %d.", expected, got)
		return
	}

	deleteKeys = append(deleteKeys, key2)

	expected4 = []string{"hi, lady", "Fast Five", "2012"}
	got4_interface, err = redisPoolObj_set.SMembers(key2)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SDIFFSTORE joe_diff_peter joe's_movies peter's_movies
	   (integer) 2
	*/
	destination := "joe_diff_peter"
	expected2 := 2
	got2, err := redisPoolObj_set.SDiffStore(destination, key1, key2)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d.", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> SMEMBERS joe_diff_peter
	   1) "bet man"
	   2) "start war"
	*/
	expected4 = []string{"bet man", "start war"}
	got4_interface, err = redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}

	/*
	   redis> SDIFFSTORE peter_diff_joe peter's_movies, joe's_movies
	   (integer) 2
	*/
	destination = "peter_diff_joe"
	expected2 = 2
	got2, err = redisPoolObj_set.SDiffStore(destination, key2, key1)
	if err != nil {
		t.Fail()
	}
	if got2 != expected2 {
		t.Errorf("Expected to get %d, but now got %d.", expected2, got2)
		return
	}

	deleteKeys = append(deleteKeys, destination)

	/*
	   redis> SMEMBERS joe_diff_peter
	   1) "hi, lady"
	   2) "Fast Five"
	*/
	expected4 = []string{"hi, lady", "Fast Five"}
	got4_interface, err = redisPoolObj_set.SMembers(destination)
	if err != nil {
		t.Fail()
	}
	got4, err = redisPoolObj_set.Strings(got4_interface)
	if err != nil {
		t.Fail()
	}
	if isTwoUnorderedSliceEqual(expected4, got4) == false {
		t.Errorf("Expected to get %v, but got %v\n", expected4, got4)
		return
	}
}
