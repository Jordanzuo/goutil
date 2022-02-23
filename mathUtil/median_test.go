package mathUtil

import (
	"testing"
)

func TestGetMedian(t *testing.T) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = GetRand().Intn(count)
	}

	value1 := GetMedian1(list)
	value2 := GetMedian2(list)
	if value1 != value2 {
		t.Errorf("Expect to get %d, but get %d instead.", value1, value2)
		return
	}
}

func TestGetMedian1(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := 5
	get := GetMedian1(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect = 5
	get = GetMedian1(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expect = 6
	get = GetMedian1(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}
}

func TestGetMedian2(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := 5
	get := GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect = 5
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect = 5
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect = 5
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{4, 5, 6, 1, 2, 3, 7, 8, 9}
	expect = 5
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{4, 7, 8, 9, 5, 6, 1, 2, 3}
	expect = 5
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}

	list = []int{4, 10, 8, 9, 5, 7, 1, 2, 3, 11}
	expect = 6
	get = GetMedian2(list)
	if get != expect {
		t.Errorf("got %v, want %v", get, expect)
		return
	}
}

func BenchmarkGetMedian11(b *testing.B) {
	count := 9999
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = GetRand().Intn(count)
	}

	for i := 0; i < b.N; i++ {
		GetMedian1(list)
	}
}

func BenchmarkGetMedian21(b *testing.B) {
	count := 9999
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = GetRand().Intn(count)
	}

	for i := 0; i < b.N; i++ {
		GetMedian2(list)
	}
}

func BenchmarkGetMedian12(b *testing.B) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = GetRand().Intn(count)
	}

	for i := 0; i < b.N; i++ {
		GetMedian1(list)
	}
}

func BenchmarkGetMedian22(b *testing.B) {
	count := 10000
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = GetRand().Intn(count)
	}

	for i := 0; i < b.N; i++ {
		GetMedian2(list)
	}
}
