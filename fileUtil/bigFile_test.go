package fileUtil

import (
	"fmt"
	"testing"
)

func BenchmarkSaveMessage(b *testing.B) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)
	bigFileObj, err := NewBigFile(path, 1024*1024*1024)
	if err != nil {
		b.Errorf("there should no err, but not there is:%s", err)
	}

	for i := 0; i < b.N; i++ {
		bigFileObj.SaveMessage(fmt.Sprintf("line %d", i))
	}
}

func TestSaveMessage(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)
	bigFileObj, err := NewBigFile(path, 1024)
	if err != nil {
		t.Errorf("there should no err, but not there is:%s", err)
	}

	for i := 0; i < 100000; i++ {
		bigFileObj.SaveMessage(fmt.Sprintf("line %d", i))
	}

	fileList, err := GetFileList(path)
	for _, item := range fileList {
		fmt.Printf("file:%s\n", item)
	}
}
