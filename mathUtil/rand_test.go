package mathUtil

import (
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
