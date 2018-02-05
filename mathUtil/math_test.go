package mathUtil

import (
	"fmt"
	"testing"
)

func TestIsContinuous_byte(t *testing.T) {
	list := make([]byte, 0, 8)
	if IsContinuous_byte(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 1)
	if IsContinuous_byte(list) == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 5)
	if IsContinuous_byte(list) == false {
		t.Errorf("it's should be true, but now false-------3")
	}

	list = append(list, 10)
	if IsContinuous_byte(list) == true {
		t.Errorf("it's should be false, but now true-------3")
	}
}

func TestIsContinuous_int(t *testing.T) {
	list := make([]int, 0, 8)
	if IsContinuous_int(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 11)
	if IsContinuous_int(list) == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 12)
	list = append(list, 13)
	list = append(list, 14)
	list = append(list, 15)
	if IsContinuous_int(list) == false {
		t.Errorf("it's should be true, but now false-------3")
	}

	list = append(list, 10)
	if IsContinuous_int(list) == true {
		t.Errorf("it's should be false, but now true-------3")
	}
}

func TestIsContinuous_int32(t *testing.T) {
	list := make([]int32, 0, 8)
	if IsContinuous_int32(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 1)
	if IsContinuous_int32(list) == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 5)
	if IsContinuous_int32(list) == false {
		t.Errorf("it's should be true, but now false-------3")
	}

	list = append(list, 10)
	if IsContinuous_int32(list) == true {
		t.Errorf("it's should be false, but now true-------3")
	}
}

func TestIsContinuous_int64(t *testing.T) {
	list := make([]int64, 0, 8)
	if IsContinuous_int64(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 1)
	if IsContinuous_int64(list) == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 5)
	if IsContinuous_int64(list) == false {
		t.Errorf("it's should be true, but now false-------3")
	}

	list = append(list, 10)
	if IsContinuous_int64(list) == true {
		t.Errorf("it's should be false, but now true-------3")
	}
}

func TestIsContinuous_Region(t *testing.T) {
	list := make([]*IntRegion, 0, 8)
	if IsContinuous_Region(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, NewIntRegion(101, 110))
	if IsContinuous_Region(list) == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, NewIntRegion(1, 10))
	if IsContinuous_Region(list) == true {
		t.Errorf("it's should be false, but now true-------2")
	}

	list = append(list, NewIntRegion(11, 100))
	if IsContinuous_Region(list) == false {
		t.Errorf("it's should be true, but now false-------3")
	}

}

func TestIsOddFullConvered(t *testing.T) {
	list := make([]*IntRegion, 0, 8)
	min, max := 1, 100
	if IsOddFullConvered(list, min, max) {
		t.Errorf("it's should be false, but now true-------1")
	}

	list = append(list, NewIntRegion(1, 10))
	if IsOddFullConvered(list, min, max) == true {
		t.Errorf("it's should be false, but now true-------2")
	}

	list = append(list, NewIntRegion(11, 100))
	if IsOddFullConvered(list, min, max) == false {
		t.Errorf("it's should be true, but now false-------1")
	}
}

func TestIsDistinct_byte(t *testing.T) {
	list := make([]byte, 0, 8)
	result := IsDistinct_byte(list)
	fmt.Printf("list:%v,result:%v-------1\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 10)
	result = IsDistinct_byte(list)
	fmt.Printf("list:%v,result:%v-------2\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 10)
	result = IsDistinct_byte(list)
	fmt.Printf("list:%v,result:%v-------3\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------3")
	}

	list = append(list, 0)
	result = IsDistinct_byte(list)
	fmt.Printf("list:%v,result:%v-------4\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------4")
	}
}

func TestIsDistinct_int(t *testing.T) {
	list := make([]int, 0, 8)
	result := IsDistinct_int(list)
	fmt.Printf("list:%v,result:%v-------1\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 10)
	result = IsDistinct_int(list)
	fmt.Printf("list:%v,result:%v-------2\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 10)
	result = IsDistinct_int(list)
	fmt.Printf("list:%v,result:%v-------3\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------3")
	}

	list = append(list, 0)
	result = IsDistinct_int(list)
	fmt.Printf("list:%v,result:%v-------4\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------4")
	}
}

func TestIsDistinct_int32(t *testing.T) {
	list := make([]int32, 0, 8)
	result := IsDistinct_int32(list)
	fmt.Printf("list:%v,result:%v-------1\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 10)
	result = IsDistinct_int32(list)
	fmt.Printf("list:%v,result:%v-------2\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 10)
	result = IsDistinct_int32(list)
	fmt.Printf("list:%v,result:%v-------3\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------3")
	}

	list = append(list, 0)
	result = IsDistinct_int32(list)
	fmt.Printf("list:%v,result:%v-------4\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------4")
	}
}

func TestIsDistinct_int64(t *testing.T) {
	list := make([]int64, 0, 8)
	result := IsDistinct_int64(list)
	fmt.Printf("list:%v,result:%v-------1\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, 10)
	result = IsDistinct_int64(list)
	fmt.Printf("list:%v,result:%v-------2\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, 10)
	result = IsDistinct_int64(list)
	fmt.Printf("list:%v,result:%v-------3\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------3")
	}

	list = append(list, 0)
	result = IsDistinct_int64(list)
	fmt.Printf("list:%v,result:%v-------4\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------4")
	}
}
