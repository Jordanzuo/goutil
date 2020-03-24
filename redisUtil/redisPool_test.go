package redisUtil

import (
	"testing"
	"time"
)

var (
	redisPoolObj *RedisPool
)

func init() {
	redisPoolObj = NewRedisPool("testPool", "localhost:6379", "redis_pwd", 0, 500, 200, 10*time.Second, 5*time.Second)
}

func TestGetName(t *testing.T) {
	expected := "testPool"
	got := redisPoolObj.GetName()
	if expected != got {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}
}

func TestGetAddress(t *testing.T) {
	expected := "localhost:6379"
	got := redisPoolObj.GetAddress()
	if expected != got {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}
}

func converInterfaceSliceToStringSlice(sourceList []interface{}) []string {
	targetList := make([]string, 0, len(sourceList))
	for _, item := range sourceList {
		if item == nil {
			targetList = append(targetList, "")
		} else if item_str, ok := item.(string); ok {
			targetList = append(targetList, item_str)
		} else if item_bytes, ok2 := item.([]byte); ok2 {
			targetList = append(targetList, string(item_bytes))
		}
	}

	return targetList
}

func isTwoOrderedSliceEqual(list1, list2 []string) bool {
	if list1 == nil && list2 == nil {
		return true
	}

	if list1 == nil || list2 == nil {
		return false
	}

	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func isTwoUnorderedSliceEqual(list1, list2 []string) bool {
	if list1 == nil && list2 == nil {
		return true
	}

	if list1 == nil || list2 == nil {
		return false
	}

	if len(list1) != len(list2) {
		return false
	}

	map1 := make(map[string]struct{})
	map2 := make(map[string]struct{})

	for _, item := range list1 {
		map1[item] = struct{}{}
	}
	for _, item := range list2 {
		map2[item] = struct{}{}
	}

	for k := range map1 {
		if _, exist := map2[k]; !exist {
			return false
		}
	}

	return true
}

func getDistinctKeyList(keyList []string) []string {
	distinctKeyList := make([]string, 0, len(keyList))
	keyMap := make(map[string]struct{})
	for _, key := range keyList {
		if _, exist := keyMap[key]; !exist {
			distinctKeyList = append(distinctKeyList, key)
			keyMap[key] = struct{}{}
		}
	}

	return distinctKeyList
}
