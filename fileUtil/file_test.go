package fileUtil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func BenchmarkWriteFile(b *testing.B) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	for i := 0; i < b.N; i++ {
		WriteFile(path, "test.txt", true, fmt.Sprintf("line %d", i))
	}
}

func TestIsFileExists(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)
	fileName := fmt.Sprintf("%s/%s", path, "test.txt")
	fmt.Printf("FileName:%s\n", fileName)
	if exists, err := IsFileExists(fileName); err != nil || exists {
		t.Errorf("the file %s should not be exists, but now it's exists", fileName)
	}

	if err := WriteFile(path, "test.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	if exists, err := IsFileExists(fileName); err != nil || !exists {
		t.Errorf("the file %s should be exists, but now it's not exists", fileName)
	}

	if content, err := ReadFileContent(fileName); err != nil {
		t.Errorf("there should be no error, but now err:%s", err)
	} else {
		fmt.Printf("Content:%s\n", content)
	}

	DeleteFile(fileName)
}

func TestIsDirectoryExists(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)
	filePath := filepath.Join(path, "Parent")
	if exists, err := IsDirectoryExists(filePath); err != nil || exists {
		t.Errorf("the file %s should not be exists, but now it's exists", filePath)
	}

	fileName := fmt.Sprintf("%s/%s", filePath, "test.txt")

	if err := WriteFile(filePath, "test.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	if exists, err := IsDirectoryExists(filePath); err != nil || !exists {
		t.Errorf("the file %s should  be exists, but now it's not exists", filePath)
	}

	if content, err := ReadFileContent(fileName); err != nil {
		t.Errorf("there should be no error, but now err:%s", err)
	} else {
		fmt.Printf("Content:%s\n", content)
	}

	DeleteFile(fileName)
}

func TestIsDirExists(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)
	filePath := filepath.Join(path, "Parent2")
	if IsDirExists(filePath) {
		t.Errorf("the file %s should not be exists, but now it's exists", filePath)
	}

	fileName := fmt.Sprintf("%s/%s", filePath, "test.txt")
	if err := WriteFile(filePath, "test.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	if IsDirExists(filePath) == false {
		t.Errorf("the file %s should  be exists, but now it's not exists", filePath)
	}

	if content, err := ReadFileContent(fmt.Sprintf("%s/%s", filePath, "test.txt")); err != nil {
		t.Errorf("there should be no error, but now err:%s", err)
	} else {
		fmt.Printf("Content:%s\n", content)
	}

	DeleteFile(fileName)
}

func TestGetFileList(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName1 := "2017-09-12-12.txt"
	fileName2 := "2017-09-12-13.txt"
	fileName3 := "2017-09-12-14.txt"
	fileName4 := "2017-09-12.tar.bz2"

	seperator := "\\"
	if runtime.GOOS != "windows" {
		seperator = "/"
	}

	filePath1 := fmt.Sprintf("%s%s%s", path, seperator, fileName1)
	filePath2 := fmt.Sprintf("%s%s%s", path, seperator, fileName2)
	filePath3 := fmt.Sprintf("%s%s%s", path, seperator, fileName3)
	filePath4 := fmt.Sprintf("%s%s%s", path, seperator, fileName4)

	if err := WriteFile(path, fileName1, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName2, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName3, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName4, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	fileList, err := GetFileList(path)
	if err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if fileList[0] != filePath1 {
		t.Errorf("Expected:%s, now got:%s", filePath1, fileList[0])
	}
	if fileList[1] != filePath2 {
		t.Errorf("Expected:%s, now got:%s", filePath2, fileList[1])
	}
	if fileList[2] != filePath3 {
		t.Errorf("Expected:%s, now got:%s", filePath3, fileList[2])
	}
	if fileList[3] != filePath4 {
		t.Errorf("Expected:%s, now got:%s", filePath4, fileList[3])
	}

	DeleteFile(filePath1)
	DeleteFile(filePath2)
	DeleteFile(filePath3)
	DeleteFile(filePath4)
}

func TestGetFileList2(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName1 := "2017-09-12-12.txt"
	fileName2 := "2017-09-12-13.txt"
	fileName3 := "2017-09-12-14.txt"
	fileName4 := "2017-09-12.tar.bz2"

	seperator := "\\"
	if runtime.GOOS != "windows" {
		seperator = "/"
	}

	filePath1 := fmt.Sprintf("%s%s%s", path, seperator, fileName1)
	filePath2 := fmt.Sprintf("%s%s%s", path, seperator, fileName2)
	filePath3 := fmt.Sprintf("%s%s%s", path, seperator, fileName3)
	filePath4 := fmt.Sprintf("%s%s%s", path, seperator, fileName4)

	if err := WriteFile(path, fileName1, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName2, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName3, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, fileName4, true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	fileList, err := GetFileList2(path, "2017-09-12", "txt")
	if err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	fmt.Printf("fileList:%v\n", fileList)
	if fileList[0] != filePath1 {
		t.Errorf("Expected:%s, now got:%s", filePath1, fileList[0])
	}
	if fileList[1] != filePath2 {
		t.Errorf("Expected:%s, now got:%s", filePath2, fileList[1])
	}
	if fileList[2] != filePath3 {
		t.Errorf("Expected:%s, now got:%s", filePath3, fileList[2])
	}

	fileList2, err := GetFileList2(path, "2017-09-12", "tar.bz2")
	if err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	fmt.Printf("fileList2:%v\n", fileList2)
	if fileList2[0] != filePath4 {
		t.Errorf("Expected:%s, now got:%s", filePath4, fileList2[0])
	}

	DeleteFile(filePath1)
	DeleteFile(filePath2)
	DeleteFile(filePath3)
	DeleteFile(filePath4)
}

func TestReadFileLineByLine(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt")
	if err := WriteFile(path, "test.txt", true, "first line\n"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, "test.txt", true, "second line\n"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	expectedFirstLine := "first line"
	expectedSecondLine := "second line"
	lineList, err := ReadFileLineByLine(fileName)
	if err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if lineList[0] != expectedFirstLine {
		t.Errorf("Expected:%s, but now got:%s", expectedFirstLine, lineList[0])
	}
	if lineList[1] != expectedSecondLine {
		t.Errorf("Expected:%s, but now got:%s", expectedSecondLine, lineList[1])
	}

	if err := DeleteFile(fileName); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}
}

func TestReadFileContent(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt")
	if err := WriteFile(path, "test.txt", true, "first line\n"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}
	if err := WriteFile(path, "test.txt", true, "second line\n"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	expectedContent := "first line\nsecond line\n"
	if content, err := ReadFileContent(fileName); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	} else if content != expectedContent {
		t.Errorf("Expected:%s, but now got:%s", expectedContent, content)
	}

	if err := DeleteFile(fileName); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}
}

func TestDeleteFile(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt")
	if err := WriteFile(path, "test.txt", true, "first line"); err != nil {
		t.Errorf("there should be no error, but now it is:%s", err)
	}

	if err := DeleteFile(fileName); err != nil {
		t.Errorf("There should be no error, but now it has:%s", err)
	}
}

func TestReadWriteSimultaneously(t *testing.T) {
	path := GetCurrentPath()
	fmt.Printf("CurrPath:%s\n", path)

	fileName := fmt.Sprintf("%s/%s", path, "test.txt")

	file1, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	if err != nil {
		t.Errorf("1:there should be no err, but now err:%s", err)
	}

	// for i := 0; i < 10; i++ {
	// 	file1.WriteString(fmt.Sprintf("line %d\n", i))
	// }

	go func() {
		for i := 0; i < 10; i++ {
			file1.WriteString(fmt.Sprintf("line %d\n", i))
			time.Sleep(time.Second)
		}
	}()

	file2, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		t.Errorf("2:there should be no err, but now err:%s", err)
	}

	go func() {
		offset := 0

		// 读取文件
		buf := bufio.NewReader(file2)

		for {
			// 按行读取
			line, _, err2 := buf.ReadLine()
			if err2 == io.EOF {
				time.Sleep(500 * time.Millisecond)
				continue
			}

			if len(line) == 0 {
				continue
			}

			//将byte[]转换为string，并添加到列表中
			fmt.Printf("line %d:%s\n", offset, string(line))

			offset += 1
			if offset >= 10 {
				break
			}
		}
	}()

	time.Sleep(30 * time.Second)

	fmt.Println("end")
}
