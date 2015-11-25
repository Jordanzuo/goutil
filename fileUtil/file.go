package fileUtil

import (
	"errors"
	"fmt"
	"github.com/Jordanzuo/goutil/stringUtil"
	"github.com/Jordanzuo/goutil/timeUtil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

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
	if !isDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}
	mutex.Unlock()

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以读写模式打开。)
	var f *os.File
	var err error
	if ifAppend == false {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
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
