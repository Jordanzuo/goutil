package main

import (
	"fmt"
	"time"

	"github.com/Jordanzuo/goutil/fileUtil"
)

func main() {
	path := "./output"
	batchFilePtr, err := fileUtil.NewBatchFile(path, "test_batch_file_9999", "txt", 1024*1024, 5)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		batchFilePtr.WriteString(fmt.Sprintf("line %d", i))
		time.Sleep(10 * time.Millisecond)
	}

	bakFileNameList, err := batchFilePtr.GetBakFilePathList()
	if err != nil {
		panic(err)
	}

	for _, item := range bakFileNameList {
		fmt.Printf("Bak file: %s\n", item)
	}
}
