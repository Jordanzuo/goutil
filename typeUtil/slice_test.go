package typeUtil

import (
	"testing"
)

func TestSliceToString(t *testing.T) {
	// Test with nil value
	var value interface{} = nil
	expected := ""
	got, err := SliceToString(value, ",")
	if err != nil {
		t.Errorf("There should be no error, but now got one. %s", err)
		return
	}

	// Test with wrong type value
	value = "hello"
	got, err = SliceToString(value, ",")
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with correct value
	value = []int{1, 2, 3, 4, 5}
	expected = "1,2,3,4,5"

	got, err = SliceToString(value, ",")
	if err != nil {
		t.Errorf("There should be no error, but now got one. %s", err)
		return
	}

	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}
}
