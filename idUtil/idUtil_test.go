package idUtil

import (
	"fmt"
	"sync"
	"testing"
)

var (
	wg       sync.WaitGroup
	mapMutex sync.Mutex
)

func TestGenerateNewId(t *testing.T) {
	idGeneratorObj, err := New(7, 40, 20)
	if err == nil {
		t.Errorf("there should be err, but now not.")
	}

	idGeneratorObj, err = New(3, 40, 20)
	if err != nil {
		t.Errorf("there should be no err, but now there is")
	}

	prefix := int64(127)
	_, err = idGeneratorObj.GenerateNewId(prefix)
	if err == nil {
		t.Errorf("there should be err, but now not.")
	}

	prefix = 5
	_, err = idGeneratorObj.GenerateNewId(prefix)
	if err != nil {
		t.Errorf("there should be no err, but now there is")
	}

	idMap := make(map[int64]struct{})
	count := 1048575

	for num := 0; num < 8; num++ {
		wg.Add(1)
		go func(prefix int64) {
			defer wg.Done()

			fmt.Printf("Prefix:%d\n", prefix)
			for i := 0; i < count; i++ {
				id, err := idGeneratorObj.GenerateNewId(prefix)
				if err != nil {
					t.Errorf("there should be no error, but now it has")
				} else {
					mapMutex.Lock()
					if _, exists := idMap[id]; exists {
						t.Errorf("there should be not duplicate, but now it does.%d", id)
					} else {
						idMap[id] = struct{}{}
					}
					mapMutex.Unlock()
					fmt.Println(id)
				}
			}
		}(int64(num))
	}

	wg.Wait()
}
