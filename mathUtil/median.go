package mathUtil

import (
	"math"
	"sort"
)

func GetMedian1(list []int) int {
	length := len(list)
	switch length {
	case 0:
		return 0
	case 1:
		return list[0]
	case 2:
		return (list[0] + list[1]) / 2
	}

	sort.Ints(list)
	if length%2 == 1 {
		return list[(length)/2]
	} else {
		return (list[length/2-1] + list[length/2]) / 2
	}
}

func GetMedian2(list []int) int {
	length := len(list)
	switch length {
	case 0:
		return 0
	case 1:
		return list[0]
	case 2:
		return (list[0] + list[1]) / 2
	}

	if length%2 == 1 {
		value := getMedian2(list, length, length/2+1)
		// fmt.Printf("value: %d\n", value)
		return value
	} else {
		list1 := make([]int, len(list))
		list2 := make([]int, len(list))
		for i, v := range list {
			list1[i] = v
			list2[i] = v
		}
		value1 := getMedian2(list1, length, length/2)
		// fmt.Printf("value1: %d\n", value1)
		value2 := getMedian2(list2, length, length/2+1)
		// fmt.Printf("value2: %d\n", value2)
		return (value1 + value2) / 2
	}
}

func getMedian2(list []int, length, count int) int {
	// fmt.Printf("list: %+v, length: %d, count: %d\n", list, length, count)
	pivot := getPivot(list)
	// fmt.Printf("pivot: %d:(%d) before\n", pivot, list[pivot])
	list, pivot = splitListByPivotal(list, pivot)
	// fmt.Printf("list: %+v\n", list)
	// fmt.Printf("pivot: %d:(%d) after\n", pivot, list[pivot])

	switch {
	case pivot == 0:
		if count == 1 {
			return list[pivot]
		} else {
			return getMedian2(list[1:], length-1, count-1)
		}
	case pivot == length-1:
		if count == length {
			return list[length-1]
		} else {
			return getMedian2(list[:length-1], length-1, count)
		}
	default:
		left, right := list[:pivot], list[pivot+1:]
		leftCount, rightCount := len(left), len(right)
		// fmt.Printf("leftCount: %d, rightCount: %d\n", leftCount, rightCount)
		switch {
		case leftCount < count:
			if leftCount+1 == count {
				return list[pivot]
			} else {
				return getMedian2(right, rightCount, count-leftCount-1)
			}
		case leftCount > count:
			return getMedian2(left, leftCount, count)
		default:
			return getMedian2(left, leftCount, count)
		}
		// switch {
		// case leftCount < rightCount:
		// 	switch {
		// 	case leftCount < count:
		// 		if leftCount+1 == count {
		// 			return list[pivot]
		// 		} else {
		// 			return getMedian2(right, rightCount, count-leftCount-1)
		// 		}
		// 	case leftCount > count:
		// 		return getMedian2(left, leftCount, count)
		// 	default:
		// 		return getMedian2(left, leftCount, count)
		// 	}
		// case leftCount > rightCount:
		// 	switch {
		// 	case rightCount < count:
		// 		if rightCount+1 == count {
		// 			return list[pivot]
		// 		} else {
		// 			return getMedian2(left, leftCount, count)
		// 		}
		// 	case rightCount > count:
		// 		return getMedian2(right, rightCount, count-leftCount-1)
		// 	default:
		// 		return getMedian2(right, rightCount, count-leftCount-1)
		// 	}
		// default:
		// 	switch {
		// 	case leftCount < count:
		// 		if leftCount+1 == count {
		// 			return list[pivot]
		// 		} else {
		// 			return getMedian2(right, rightCount, count-leftCount-1)
		// 		}
		// 	case leftCount > count:
		// 		return getMedian2(left, leftCount, count)
		// 	default:
		// 		return getMedian2(left, leftCount, count)
		// 	}
		// }
	}
}

// There are two implementations of splitListByPivotal
// The first one is on place adjust, which takes more time;
// The second one uses extra space, which takes less time.

// func splitListByPivotal(list []int, pivot int) int {
// 	pivotNum := list[pivot]
// 	for i := 0; i < len(list); i++ {
// 		switch {
// 		case i < pivot:
// 			if list[i] > pivotNum {
// 				// fmt.Printf("i<pivot: %d<%d\n", i, pivot)
// 				for j := pivot - 1; j >= i; j-- {
// 					list[j+1] = list[j]
// 				}
// 				list[i] = pivotNum
// 				pivot = i
// 				// fmt.Printf("list in split: %+v, pivot: %d\n", list, pivot)
// 			}
// 		case i > pivot:
// 			if list[i] < pivotNum {
// 				// fmt.Printf("i>pivot: %d>%d\n", i, pivot)
// 				tmp := list[i]
// 				for j := i - 1; j >= pivot; j-- {
// 					list[j+1] = list[j]
// 				}
// 				list[pivot] = tmp
// 				pivot = pivot + 1
// 				// fmt.Printf("list in split: %+v, pivot: %d\n", list, pivot)
// 			}
// 		default:
// 			// fmt.Printf("Do nothing: %d\n", pivot)
// 			// Do nothing
// 		}
// 	}

// 	return pivot
// }

func splitListByPivotal(list []int, pivot int) ([]int, int) {
	pivotNum := list[pivot]
	lessCount := 0
	greaterThanOrEqualCount := 0
	for i, v := range list {
		if i == pivot {
			continue
		}
		if v <= pivotNum {
			lessCount++
		} else {
			greaterThanOrEqualCount++
		}
	}

	newList := make([]int, len(list))
	lessIndex, pivotIndex, greaterThanOrEqualIndex := 0, lessCount, lessCount+1
	for i, v := range list {
		if i == pivot {
			newList[pivotIndex] = v
			continue
		}
		if v <= pivotNum {
			newList[lessIndex] = v
			lessIndex++
		} else {
			newList[greaterThanOrEqualIndex] = v
			greaterThanOrEqualIndex++
		}
	}

	return newList, pivotIndex
}

func getPivot(list []int) int {
	length := len(list)
	if length < 3 {
		return 0
	}

	index1 := 0
	index2 := length / 2
	index3 := length - 1

	pivotal1 := list[index1]
	pivotal2 := list[index2]
	pivotal3 := list[index3]

	max := math.MinInt
	if pivotal1 > max {
		max = pivotal1
	}
	if pivotal2 > max {
		max = pivotal2
	}
	if pivotal3 > max {
		max = pivotal3
	}

	switch max {
	case pivotal1:
		if pivotal2 < pivotal3 {
			return index3
		} else {
			return index2
		}
	case pivotal2:
		if pivotal1 < pivotal3 {
			return index3
		} else {
			return index1
		}
	default:
		if pivotal1 < pivotal2 {
			return index2
		} else {
			return index1
		}
	}
}
