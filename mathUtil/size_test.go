package mathUtil

import (
	"testing"
)

func TestGetSizeDesc(t *testing.T) {
	var size int64
	var expectedStr string
	var finalStr string

	size = 1
	expectedStr = "1B"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1KB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1MB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1.00GB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1.00TB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1.00PB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

	size *= 1024
	expectedStr = "1.00EB"
	finalStr = GetSizeDesc(size)
	if finalStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, finalStr)
	}

}
