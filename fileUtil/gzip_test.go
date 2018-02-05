package fileUtil

import (
	"fmt"
	"testing"
)

func TestGzip(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt")
	if err := WriteFile(path, "test.txt", true, "first line\nHello world"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	if err := Gzip(fileName, ""); err != nil {
		// if err := Gzip(fileName, path); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}

	if fileList, err := GetFileList(path); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	} else {
		for _, item := range fileList {
			fmt.Printf("item:%s\n", item)
		}
	}

	DeleteFile(fileName)
}

func TestUnGzip(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt.gz")
	if err := UnGzip(fileName, ""); err != nil {
		// if err := UnGzip(fileName, path); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}

	content, err := ReadFileContent(fmt.Sprintf("%s/%s", path, "test.txt"))
	if err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	} else {
		fmt.Printf("content:%s\n", content)
	}

	DeleteFile(fileName)

	fileName = fmt.Sprintf("%s/%s", path, "test.txt")
	DeleteFile(fileName)
}
