package fileUtil

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var (
	mutex sync.Mutex
)

// 文件是否存在
// 文件路径
// 返回值：
// 是否存在
// 错误对象
func IsFileExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir() == false, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		}
	}

	return true, err
}

// 文件夹是否存在
// 文件夹路径
// 返回值：
// 是否存在
// 错误对象
func IsDirectoryExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir(), nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		}
	}

	return true, err
}

// 文件夹是否存在(obsolete)
// 文件夹路径
// 返回值：
// 是否存在
func IsDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
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
func GetFileList(path string) (fileList []string, err error) {
	if exists, err1 := IsDirectoryExists(path); err1 != nil {
		err = err1
		return
	} else if !exists {
		return
	}

	// 遍历目录，获取所有文件列表
	err = filepath.Walk(path, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if fi.IsDir() {
			return nil
		}

		// 添加到列表
		fileList = append(fileList, fileName)

		return nil
	})

	return
}

// 获取目标文件列表（完整路径）
// path：文件夹路径
// prefix：文件前缀
// suffix：文件后缀
// 返回值：文件列表（完整路径）
func GetFileList2(path, prefix, suffix string) (fileList []string, err error) {
	if exists, err1 := IsDirectoryExists(path); err1 != nil {
		err = err1
		return
	} else if !exists {
		return
	}

	// 遍历目录，获取所有文件列表
	err = filepath.Walk(path, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if fi.IsDir() {
			return nil
		}

		// 添加到列表
		baseName := filepath.Base(fileName)
		if prefix != "" && strings.HasPrefix(baseName, prefix) == false {
			return nil
		}

		if suffix != "" && strings.HasSuffix(baseName, suffix) == false {
			return nil
		}

		fileList = append(fileList, fileName)

		return nil
	})

	return
}

// 按行读取每一个文件的内容
// fileName:文件的绝对路径
// 返回值：
// 行内容列表
// 错误信息
func ReadFileLineByLine(fileName string) (lineList []string, err error) {
	// 打开文件
	file, err1 := os.Open(fileName)
	if err1 != nil {
		err = err1
		return
	}
	defer file.Close()

	// 读取文件
	buf := bufio.NewReader(file)
	for {
		// 按行读取
		line, _, err2 := buf.ReadLine()
		if err2 == io.EOF {
			break
		}

		//将byte[]转换为string，并添加到列表中
		lineList = append(lineList, string(line))
	}

	return
}

// 读取文件内容（字符串）
// fileName：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileContent(fileName string) (content string, err error) {
	bytes, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		err = err1
		return
	}

	content = string(bytes)
	return
}

// 读取文件内容（字符数组）
// fileName：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileBytes(fileName string) (content []byte, err error) {
	content, err = ioutil.ReadFile(fileName)
	return
}

// 写入文件
// filePath：文件夹路径
// fileName：文件名称
// ifAppend：是否追加内容
// args：可变参数
// 返回值:
// error:错误信息
func WriteFile(filePath, fileName string, ifAppend bool, args ...string) error {
	// 得到最终的fileName
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	mutex.Lock()
	if !IsDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}
	mutex.Unlock()

	// 打开文件(如果文件存在就以写模式打开，并追加写入；如果文件不存在就创建，然后以写模式打开。)
	var f *os.File
	var err error
	if ifAppend == false {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm|os.ModeTemporary)
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	}

	if err != nil {
		return err
	}
	defer f.Close()

	// 写入内容
	for _, arg := range args {
		_, err = f.WriteString(arg)
		if err != nil {
			return err
		}
	}

	return nil
}

// 写入文件
// filePath：文件夹路径
// fileName：文件名称
// ifAppend：是否追加内容
// args：可变参数
// 返回值:
// error:错误信息
func WriteFile4Byte(filePath, fileName string, ifAppend bool, args ...[]byte) error {
	// 得到最终的fileName
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	mutex.Lock()
	if !IsDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}
	mutex.Unlock()

	// 打开文件(如果文件存在就以写模式打开，并追加写入；如果文件不存在就创建，然后以写模式打开。)
	var f *os.File
	var err error
	if ifAppend == false {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm|os.ModeTemporary)
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	}

	if err != nil {
		return err
	}
	defer f.Close()

	// 写入内容
	for _, arg := range args {
		_, err = f.Write(arg)
		if err != nil {
			return err
		}
	}

	return nil
}

// 删除文件
// fileName：文件的绝对路径
// 返回值：
// 错误对象
func DeleteFile(fileName string) error {
	return os.Remove(fileName)
}
