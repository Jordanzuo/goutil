package fileUtil

import (
	"fmt"
	"strings"
	"testing"
)

func TestTar(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName1 := fmt.Sprintf("%s/%s", path, "test1.txt")
	fileName2 := fmt.Sprintf("%s/%s", path, "test2.txt")

	if err := WriteFile(path, "test1.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, "test2.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	sourceList := make([]string, 0, 2)
	sourceList = append(sourceList, fileName1)
	sourceList = append(sourceList, fileName2)
	target := fmt.Sprintf("%s/%s", path, "test.tar")
	if err := Tar(sourceList, target); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}

	if fileList, err := GetFileList(path); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	} else {
		for _, item := range fileList {
			fmt.Printf("item:%s\n", item)
		}
	}

	DeleteFile(fileName1)
	DeleteFile(fileName2)
}

func TestUntar(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	source := fmt.Sprintf("%s/%s", path, "test.tar")
	// target := path
	target := ""
	if err := Untar(source, target); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}

	if fileList, err := GetFileList(path); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	} else {
		for _, item := range fileList {
			fmt.Printf("item:%s\n", item)

			if strings.HasSuffix(item, "txt") {
				if content, err := ReadFileContent(item); err != nil {
					t.Errorf("There should be no error, but now it has:%s", err)
				} else {
					fmt.Printf("content:%s\n", content)
				}

				DeleteFile(item)
			}
		}

		DeleteFile(source)
	}
}
