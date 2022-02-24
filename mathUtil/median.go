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
		return getMedian2(list, length/2+1)
	} else {
		value1 := getMedian2(list, length/2)
		value2 := getMedian2(list, length/2+1)
		return (value1 + value2) / 2
	}
}

func getMedian2(list []int, nth int) int {
	length := len(list)
	pivot := getPivot(list)
	pivot = splitListByPivotal(list, pivot)
	switch {
	case pivot == 0:
		if nth == 1 {
			return list[pivot]
		} else {
			return getMedian2(list[1:], nth-1)
		}
	case pivot == length-1:
		if nth == length {
			return list[length-1]
		} else {
			return getMedian2(list[:length-1], nth)
		}
	default:
		left, right := list[:pivot], list[pivot+1:]
		leftCount := len(left)
		switch {
		case leftCount < nth:
			if leftCount+1 == nth {
				return list[pivot]
			} else {
				return getMedian2(right, nth-leftCount-1)
			}
		case leftCount > nth:
			return getMedian2(left, nth)
		default:
			return getMedian2(left, nth)
		}
	}
}

func splitListByPivotal(list []int, pivot int) int {
	pivotNum := list[pivot]
	for i := 0; i < len(list); i++ {
		switch {
		case i < pivot:
			if list[i] > pivotNum {
				for j := pivot - 1; j >= i; j-- {
					list[j+1] = list[j]
				}
				list[i] = pivotNum
				pivot = i
			}
		case i > pivot:
			if list[i] < pivotNum {
				tmp := list[i]
				for j := i - 1; j >= pivot; j-- {
					list[j+1] = list[j]
				}
				list[pivot] = tmp
				pivot = pivot + 1
			}
		default:
			// Do nothing
		}
	}

	return pivot
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
