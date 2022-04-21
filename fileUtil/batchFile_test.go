package fileUtil

import (
	"fmt"
	"testing"
)

func TestBatchFileWriteString(t *testing.T) {
	path := "/home/jordan/Documents/github.com/Jordanzuo/goutil/fileUtil/output"
	fmt.Printf("TestBatchFileWriteString CurrPath:%s\n", path)
	batchFilePtr, err := NewBatchFile(path, "test_batch_file.txt", 100, 30)
	if err != nil {
		t.Errorf("there should no err, but not there is:%s", err)
	}

	for i := 0; i < 100000; i++ {
		batchFilePtr.WriteString(fmt.Sprintf("line %d", i))
	}

	fileList, err := GetFileList(path)
	for _, item := range fileList {
		fmt.Printf("file:%s\n", item)
	}
}

func BenchmarkBatchFileWriteString(b *testing.B) {
	path := "/home/jordan/Documents/github.com/Jordanzuo/goutil/fileUtil/output"
	fmt.Printf("TestBatchFileWriteString CurrPath:%s\n", path)
	batchFilePtr, err := NewBatchFile(path, "test_batch_file.txt", 1024*1024*1024, 30)
	if err != nil {
		b.Errorf("there should no err, but not there is:%s", err)
	}

	for i := 0; i < b.N; i++ {
		batchFilePtr.WriteString(fmt.Sprintf("line %d", i))
	}
}
