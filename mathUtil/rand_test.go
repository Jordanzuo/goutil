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
	randNumList, _ = GetRandNumList(1, 10, 6, false)
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
	item1 := NewItem(1, "name1")
	item2 := NewItem(2, "name1")
	item3 := NewItem(3, "name1")
	item4 := NewItem(4, "name1")
	item5 := NewItem(5, "name1")
	item6 := NewItem(6, "name1")
	item7 := NewItem(7, "name1")
	item8 := NewItem(8, "name1")
	item9 := NewItem(9, "name1")
	item10 := NewItem(10, "name1")

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

type Item struct {
	Id   int
	Name string
}

func (item *Item) String() string {
	return fmt.Sprintf("Id:%d,Name:%s\n", item.Id, item.Name)
}

func NewItem(id int, name string) *Item {
	return &Item{
		Id:   id,
		Name: name,
	}
}
