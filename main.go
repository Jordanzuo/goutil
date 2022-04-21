package main

import (
	"fmt"
	"github.com/Jordanzuo/goutil/fileUtil"
	"time"
)

func main() {
	path := "./output"
	batchFilePtr, err := fileUtil.NewBatchFile(path, "test_batch_file.txt", 1024*1024, 5)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		batchFilePtr.WriteString(fmt.Sprintf("line %d", i))
		time.Sleep(10 * time.Millisecond)
	}
}
