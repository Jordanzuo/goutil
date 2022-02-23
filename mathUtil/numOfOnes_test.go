package mathUtil

import "testing"

func TestGetNumOfOnes1_uint32(t *testing.T) {
	for i, v := range numOfOnesArray {
		getNum := GetNumOfOnes1_uint32(uint32(i))
		expectNum := int(v)
		if getNum != expectNum {
			t.Errorf("got %d, want %d for %d", getNum, expectNum, i)
		}
	}
}

func TestGetNumOfOnes2_uint32(t *testing.T) {
	for i, v := range numOfOnesArray {
		getNum := GetNumOfOnes2_uint32(uint32(i))
		expectNum := int(v)
		if getNum != expectNum {
			t.Errorf("got %d, want %d for %d", getNum, expectNum, i)
		}
	}
}

func BenchmarkGetNumOfOnes1_uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNumOfOnes1_uint32(uint32(i))
	}
}

func BenchmarkGetNumOfOnes2_uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNumOfOnes2_uint32(uint32(i))
	}
}
