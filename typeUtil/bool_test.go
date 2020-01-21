package typeUtil

import (
	"testing"
)

func TestBoolToInt(t *testing.T) {
	// Test with true value
	value := true
	expected := 1
	got := BoolToInt(value)
	if expected != got {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with false value
	value = false
	expected = 0
	got = BoolToInt(value)
	if expected != got {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestIntToBool(t *testing.T) {
	// Test with 0 value
	value := 0
	expected := false
	got := IntToBool(value)
	if expected != got {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with negative value
	value = -1
	expected = false
	got = IntToBool(value)
	if expected != got {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with positive value
	value = 1
	expected = true
	got = IntToBool(value)
	if expected != got {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with bigger positive value
	value = 100
	expected = true
	got = IntToBool(value)
	if expected != got {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}
}
