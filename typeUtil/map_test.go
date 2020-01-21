package typeUtil

import (
	"testing"
)

func TestMapToString(t *testing.T) {
	var data map[string]int
	separator1 := ","
	separator2 := ";"

	got, err := MapToString(data, separator1, separator2)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}

	data1 := make([]int, 0, 4)
	data1 = append(data1, 1)
	got, err = MapToString(data1, separator1, separator2)
	if err == nil {
		t.Errorf("There should be an error, but now there is not")
		return
	}

	data2 := make(map[string]int, 4)
	data2["Jordan"] = 34
	data2["Thomas"] = 6
	expected1 := "Jordan,34;Thomas,6"
	expected2 := "Thomas,6;Jordan,34"

	got, err = MapToString(data2, separator1, separator2)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected1 && got != expected2 {
		t.Errorf("Expected to get:%s or %s, but got:%s", expected1, expected2, got)
	}
}

func TestMapToString2(t *testing.T) {
	var data map[string]int
	separator1 := ","
	separator2 := ";"

	got, err := MapToString2(data, separator1, separator2, valGetFunc)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}

	data1 := make([]int, 0, 4)
	data1 = append(data1, 1)
	got, err = MapToString2(data1, separator1, separator2, valGetFunc)
	if err == nil {
		t.Errorf("There should be an error, but now there is not")
		return
	}

	data2 := make(map[string]int, 4)
	data2["Jordan"] = 34
	data2["Thomas"] = 6
	expected1 := "Jordan,34;Thomas,6"
	expected2 := "Thomas,6;Jordan,34"

	got, err = MapToString2(data2, separator1, separator2, valGetFunc)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected1 && got != expected2 {
		t.Errorf("Expected to get:%s or %s, but got:%s", expected1, expected2, got)
	}
}

func valGetFunc(val interface{}) interface{} {
	return val
}
