package stringUtil

import (
	"testing"
)

func TestStringListToString(t *testing.T) {
	list := make([]string, 0, 4)
	list = append(list, "Hello")
	list = append(list, "World")
	list = append(list, "Hello")
	list = append(list, "Apple")

	expected := "Hello,World,Hello,Apple"
	got := StringListToString(list, ",")
	if expected != got {
		t.Errorf("Expected:%s, but got:%s", expected, got)
	}
}

func TestIntListToString(t *testing.T) {
	list := make([]int, 0, 4)
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)

	expected := "1,2,3,4"
	got := IntListToString(list, ",")
	if expected != got {
		t.Errorf("Expected:%s, but got:%s", expected, got)
	}
}

func TestInt64ListToString(t *testing.T) {
	list := make([]int64, 0, 4)
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)

	expected := "1,2,3,4"
	got := Int64ListToString(list, ",")
	if expected != got {
		t.Errorf("Expected:%s, but got:%s", expected, got)
	}
}

func TestInt32ListToString(t *testing.T) {
	list := make([]int32, 0, 4)
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)

	expected := "1,2,3,4"
	got := Int32ListToString(list, ",")
	if expected != got {
		t.Errorf("Expected:%s, but got:%s", expected, got)
	}
}
