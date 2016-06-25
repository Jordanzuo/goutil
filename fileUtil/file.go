package fileUtil

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	mutex sync.Mutex
)

// 文件夹是否存在
// 文件夹路径
// 返回值：
// 是否存在
func IsDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

// 获取当前路径
// 返回值：
// 当前路径
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	fileAbsPath, _ := filepath.Abs(file)

	return filepath.Dir(fileAbsPath)
}

// 获取目标文件列表（完整路径）
// path：文件夹路径
// 返回值：文件列表（完整路径）
func GetFileList(path string) ([]string, error) {
	files := make([]string, 0, 100)

	//遍历目录，获取所有文件列表
	filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//忽略目录
		if fi.IsDir() {
			return nil
		}

		// 添加到列表
		files = append(files, filename)

		return nil
	})

	return files, nil
}

// 按行读取每一个文件的内容
// filename:文件的绝对路径
// 返回值：
// 行内容列表
// 错误信息
func ReadFileLineByLine(filename string) ([]string, error) {
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//读取文件
	lineList := make([]string, 0, 100)
	buf := bufio.NewReader(file)
	for {
		//按行读取
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}

		//将byte[]转换为string，并添加到列表中
		lineList = append(lineList, string(line))
	}

	return lineList, nil
}

// 读取文件内容（字符串）
// filename：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileContent(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// 读取文件内容（字符数组）
// filename：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileBytes(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)

	return bytes, err
}

// 写入文件
// filePath：文件夹路径
// fileName：文件名称
// ifAppend：是否追加内容
// args：可变参数
func WriteFile(filePath, fileName string, ifAppend bool, args ...string) {
	// 得到最终的fileName
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	mutex.Lock()
	if !IsDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}
	mutex.Unlock()

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以读写模式打开。)
	var f *os.File
	var err error
	if ifAppend == false {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm|os.ModeTemporary)
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	}

	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer f.Close()

	// 写入内容
	for _, arg := range args {
		f.WriteString(arg)
	}
}

// 删除文件
// filename：文件的绝对路径
// 返回值：
// 错误对象
func DeleteFile(filename string) error {
	return os.Remove(filename)
}
