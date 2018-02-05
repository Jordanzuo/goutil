package mathUtil

import (
	"fmt"
	"testing"
)

func TestGetRandRangeInt(t *testing.T) {
	lower, upper := 10, 100
	rand := GetRandRangeInt(lower, upper)
	if rand < lower || rand >= upper {
		t.Errorf("Expected a num between %d and %d, but got %d", lower, upper, rand)
	}
}

func TestGetRandInt(t *testing.T) {
	var n int = 100
	var rand int = GetRandInt(n)
	if rand >= n {
		t.Errorf("Expected a num < %d, but got %d", n, rand)
	}
}

func TestGetRandInt32(t *testing.T) {
	var n int32 = 100
	var rand int32 = GetRandInt32(n)
	if rand >= n {
		t.Errorf("Expected a num < %d, but got %d", n, rand)
	}
}

func TestGetRandInt64(t *testing.T) {
	var n int64 = 100
	var rand int64 = GetRandInt64(n)
	if rand >= n {
		t.Errorf("Expected a num < %d, but got %d", n, rand)
	}
}

func TestGetRandFloat32(t *testing.T) {
	var rand float32 = GetRandFloat32()
	if rand >= 1 {
		t.Errorf("Expected a num < 1, but got %f", rand)
	}
}

func TestGetRandFloat64(t *testing.T) {
	var rand float64 = GetRandFloat64()
	if rand >= 1 {
		t.Errorf("Expected a num < 1, but got %f", rand)
	}
}

func TestGetRandNumList(t *testing.T) {
	if _, err := GetRandNumList(11, 10, 11, false); err.Error() != "minValue can't be bigger than maxValue." {
		t.Error("Expected err, but got nil")
	}

	if _, err := GetRandNumList(1, 10, 11, false); err.Error() != "随机的数量超过区间的元素数量" {
		t.Error("Expected err, but got nil")
	}

	if _, err := GetRandNumList(1, 10001, 10, false); err.Error() != "随机数的区间不能大于10000" {
		t.Error("Expected err, but got nil")
	}

	randNumList, _ := GetRandNumList(1, 10, 1, false)
	fmt.Printf("randNumList:%v\n", randNumList)
	randNumList, _ = GetRandNumList(1, 10, 3, false)
	fmt.Printf("randNumList:%v\n", randNumList)
	randNumList, _ = GetRandNumList(1, 10, 5, false)
	fmt.Printf("randNumList:%v\n", randNumList)
	randNumList, _ = GetRandNumList(1, 10, 7, false)
	fmt.Printf("randNumList:%v\n", randNumList)
	randNumList, _ = GetRandNumList(1, 10, 9, false)
	fmt.Printf("randNumList:%v\n", randNumList)
	randNumList, _ = GetRandNumList(1, 10, 10, true)
	fmt.Printf("randNumList:%v\n", randNumList)
}

func TestGetRandIntList(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if _, err := GetRandIntList(source, 11, false); err.Error() != "随机的数量超过列表的元素数量" {
		t.Error("Expected err, but got nil")
	}

	randIntList, _ := GetRandIntList(source, 1, false)
	fmt.Printf("randIntList:%v\n", randIntList)
	randIntList, _ = GetRandIntList(source, 3, false)
	fmt.Printf("randIntList:%v\n", randIntList)
	randIntList, _ = GetRandIntList(source, 5, false)
	fmt.Printf("randIntList:%v\n", randIntList)
	randIntList, _ = GetRandIntList(source, 7, false)
	fmt.Printf("randIntList:%v\n", randIntList)
	randIntList, _ = GetRandIntList(source, 9, false)
	fmt.Printf("randIntList:%v\n", randIntList)
	randIntList, _ = GetRandIntList(source, 10, true)
	fmt.Printf("randIntList:%v\n", randIntList)
}

func TestGetRandInterfaceList(t *testing.T) {
	item1 := NewItem(1, "name1", 1)
	item2 := NewItem(2, "name1", 1)
	item3 := NewItem(3, "name1", 1)
	item4 := NewItem(4, "name1", 1)
	item5 := NewItem(5, "name1", 1)
	item6 := NewItem(6, "name1", 1)
	item7 := NewItem(7, "name1", 1)
	item8 := NewItem(8, "name1", 1)
	item9 := NewItem(9, "name1", 1)
	item10 := NewItem(10, "name1", 1)

	source := make([]interface{}, 0, 10)
	source = append(source, item1)
	source = append(source, item2)
	source = append(source, item3)
	source = append(source, item4)
	source = append(source, item5)
	source = append(source, item6)
	source = append(source, item7)
	source = append(source, item8)
	source = append(source, item9)
	source = append(source, item10)

	if _, err := GetRandInterfaceList(source, 11, false); err.Error() != "随机的数量超过列表的元素数量" {
		t.Error("Expected err, but got nil")
	}

	randInterfaceList, _ := GetRandInterfaceList(source, 1, false)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
	randInterfaceList, _ = GetRandInterfaceList(source, 3, false)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
	randInterfaceList, _ = GetRandInterfaceList(source, 5, false)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
	randInterfaceList, _ = GetRandInterfaceList(source, 7, false)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
	randInterfaceList, _ = GetRandInterfaceList(source, 9, false)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
	randInterfaceList, _ = GetRandInterfaceList(source, 10, true)
	fmt.Printf("randInterfaceList:%v\n", randInterfaceList)
}

func TestGetRandWeight(t *testing.T) {
	source := make([]IWeight, 0, 10)
	if _, err := GetRandWeight(source); err == nil {
		t.Errorf("err should not be nil, but it's nil")
	}

	item1 := NewItem(1, "name1", 1)
	item2 := NewItem(2, "name2", 2)
	item3 := NewItem(3, "name3", 3)
	item4 := NewItem(4, "name4", 4)
	item5 := NewItem(5, "name5", 5)
	item6 := NewItem(6, "name6", 60)
	item7 := NewItem(7, "name7", 70)
	item8 := NewItem(8, "name8", 80)
	item9 := NewItem(9, "name9", 90)
	item10 := NewItem(10, "name10", 100)

	source = append(source, item1)
	source = append(source, item2)
	source = append(source, item3)
	source = append(source, item4)
	source = append(source, item5)
	source = append(source, item6)
	source = append(source, item7)
	source = append(source, item8)
	source = append(source, item9)
	source = append(source, item10)

	data := make(map[int]int)
	for i := 0; i < 10000; i++ {
		if result, err := GetRandWeight(source); err != nil {
			t.Errorf("err should be nil, but it's not:%s", err)
		} else {
			if item, ok := result.(*Item); !ok {
				t.Errorf("convert to Item failed")
			} else {
				if count, ok := data[item.Id]; ok {
					data[item.Id] = count + 1
				} else {
					data[item.Id] = 1
				}
			}
		}
	}

	total := 0
	for _, v := range data {
		total += v
	}

	k := 1
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 2
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 3
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 4
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 5
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 6
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 7
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 8
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 9
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
	k = 10
	if v, ok := data[k]; ok {
		fmt.Printf("%d:%d, ratio:%d\n", k, v, v*100/total)
	}
}

type Item struct {
	Id     int
	Name   string
	Weight int
}

func (item *Item) GetWeight() int {
	return item.Weight
}

func (item *Item) String() string {
	return fmt.Sprintf("Id:%d,Name:%s\n", item.Id, item.Name)
}

func NewItem(id int, name string, weight int) *Item {
	return &Item{
		Id:     id,
		Name:   name,
		Weight: weight,
	}
}
