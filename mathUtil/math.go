package mathUtil

import (
	"sort"
)

// 判断传入的byte型数组是否连续
func IsContinuous_byte(list []byte) bool {
	if len(list) == 0 || len(list) == 1 {
		return true
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsContinuous_int64(list_int64)
}

// 判断传入的int型数组是否连续
func IsContinuous_int(list []int) bool {
	if len(list) == 0 || len(list) == 1 {
		return true
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsContinuous_int64(list_int64)
}

// 判断传入的int型数组是否连续
func IsContinuous_int32(list []int32) bool {
	if len(list) == 0 || len(list) == 1 {
		return true
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsContinuous_int64(list_int64)
}

// 判断传入的int型数组是否连续
func IsContinuous_int64(list []int64) bool {
	if len(list) == 0 || len(list) == 1 {
		return true
	}

	if int64(len(list)) != list[len(list)-1]-list[0]+1 {
		return false
	}

	for i := 0; i < len(list)-1; i++ {
		if list[i]+1 != list[i+1] {
			return false
		}
	}

	return true
}

// 判断区间是否连续
func IsContinuous_Region(list []*IntRegion) bool {
	if len(list) == 0 || len(list) == 1 {
		return true
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Lower < list[j].Lower })

	for i := 0; i < len(list)-1; i++ {
		if list[i].IsSorted() == false || list[i+1].IsSorted() == false {
			return false
		}

		if list[i].Upper+1 != list[i+1].Lower {
			return false
		}
	}

	return true
}

// 判断传入的概率是否全覆盖
func IsOddFullConvered(list []*IntRegion, min, max int) bool {
	if len(list) == 0 {
		return false
	}

	if list[0].Lower != min || list[len(list)-1].Upper != max {
		return false
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Lower < list[j].Lower })

	for i := 0; i < len(list)-1; i++ {
		if list[i].IsSorted() == false || list[i+1].IsSorted() == false {
			return false
		}

		if list[i].Upper+1 != list[i+1].Lower {
			return false
		}
	}

	return true
}

func IsDistinct_byte(list []byte) (result bool) {
	if len(list) == 0 || len(list) == 1 {
		result = true
		return
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsDistinct_int64(list_int64)
}

func IsDistinct_int(list []int) (result bool) {
	if len(list) == 0 || len(list) == 1 {
		result = true
		return
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsDistinct_int64(list_int64)
}

func IsDistinct_int32(list []int32) (result bool) {
	if len(list) == 0 || len(list) == 1 {
		result = true
		return
	}

	list_int64 := make([]int64, len(list), len(list))
	for i, v := range list {
		list_int64[i] = int64(v)
	}

	return IsDistinct_int64(list_int64)
}

// 判断int64数组是否内容唯一
func IsDistinct_int64(list []int64) (result bool) {
	if len(list) == 0 || len(list) == 1 {
		result = true
		return
	}

	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })

	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			result = false
			return
		}
	}

	result = true
	return
}
