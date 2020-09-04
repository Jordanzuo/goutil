package stringUtil

import (
	"testing"
)

func TestStringToMap_String_String(t *testing.T) {
	str := ""
	seps := []string{",", "|"}
	data, err := StringToMap_String_String(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3"
	data, err = StringToMap_String_String(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,5"
	data, err = StringToMap_String_String(str, seps)
	if err != nil {
		t.Errorf("Expected to get no error. But now there is one:%v.", err)
		return
	}

	expected := make(map[string]string, 2)
	expected["1"] = "2"
	expected["3"] = "5"

	if len(expected) != len(data) {
		t.Errorf("The length of expected:%d is not equals to length of data:%d", len(expected), len(data))
		return
	}

	for k, v := range data {
		if v1, exists := expected[k]; !exists {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		} else if v != v1 {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		}
	}
}

func TestStringToMap_String_Int(t *testing.T) {
	str := ""
	seps := []string{",", "|"}
	data, err := StringToMap_String_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3"
	data, err = StringToMap_String_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,abc"
	data, err = StringToMap_String_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,5"
	data, err = StringToMap_String_Int(str, seps)
	if err != nil {
		t.Errorf("Expected to get no error. But now there is one:%v.", err)
		return
	}

	expected := make(map[string]int, 2)
	expected["1"] = 2
	expected["3"] = 5

	if len(expected) != len(data) {
		t.Errorf("The length of expected:%d is not equals to length of data:%d", len(expected), len(data))
		return
	}

	for k, v := range data {
		if v1, exists := expected[k]; !exists {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		} else if v != v1 {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		}
	}
}

func TestStringToMap_Int_Int(t *testing.T) {
	str := ""
	seps := []string{",", "|"}
	data, err := StringToMap_Int_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3"
	data, err = StringToMap_Int_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,abc"
	data, err = StringToMap_Int_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|a,3"
	data, err = StringToMap_Int_Int(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,5"
	data, err = StringToMap_Int_Int(str, seps)
	if err != nil {
		t.Errorf("Expected to get no error. But now there is one:%v.", err)
		return
	}

	expected := make(map[int]int, 2)
	expected[1] = 2
	expected[3] = 5

	if len(expected) != len(data) {
		t.Errorf("The length of expected:%d is not equals to length of data:%d", len(expected), len(data))
		return
	}

	for k, v := range data {
		if v1, exists := expected[k]; !exists {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		} else if v != v1 {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		}
	}
}

func TestStringToMap_Int32_Int32(t *testing.T) {
	str := ""
	seps := []string{",", "|"}
	data, err := StringToMap_Int32_Int32(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3"
	data, err = StringToMap_Int32_Int32(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,abc"
	data, err = StringToMap_Int32_Int32(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|a,3"
	data, err = StringToMap_Int32_Int32(str, seps)
	if err == nil {
		t.Errorf("Expected to get an error. But now there isn't.")
		return
	}

	str = "1,2|3,5"
	data, err = StringToMap_Int32_Int32(str, seps)
	if err != nil {
		t.Errorf("Expected to get no error. But now there is one:%v.", err)
		return
	}

	expected := make(map[int32]int32, 2)
	expected[1] = 2
	expected[3] = 5

	if len(expected) != len(data) {
		t.Errorf("The length of expected:%d is not equals to length of data:%d", len(expected), len(data))
		return
	}

	for k, v := range data {
		if v1, exists := expected[k]; !exists {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		} else if v != v1 {
			t.Errorf("data is not equals to expected. %v, %v", expected, data)
		}
	}
}
