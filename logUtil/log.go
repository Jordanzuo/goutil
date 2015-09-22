package logUtil

import (
	"errors"
	"fmt"
	"github.com/Jordanzuo/goutil/timeUtil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	SEPERATOR = "------------------------------------------------------\n"
	MAX_SKIP  = 5
)

var (
	LogPath  string
	LogMutex sync.Mutex
)

// 设置日志存放的路径
// path：日志文件存放路径
func SetLogPath(path string) {
	LogPath = path
}

// 获取日志文件存放路径
// 返回值：日志文件存放路径
func GetLogPath() string {
	return LogPath
}

// 判断目录是否存在
// path：文件路径
// 返回值：目录是否存在2
func isDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

// 记录日志
// logInfo：需要记录的日志信息
// level：日志级别
// ifIncludeHour：日志文件名称是否包含小时
// 返回值：无
func Log(logInfo string, level LogType, ifIncludeHour bool) {
	// 判断路径是否为空
	if len(LogPath) == 0 {
		panic(errors.New("日志存放路径不能为空，请先设置"))
	}

	// 获取当前时间
	now := time.Now()

	// 获得年、月、日、时的字符串形式
	yearString := strconv.Itoa(now.Year())
	monthString := strconv.Itoa(int(now.Month()))
	dayString := strconv.Itoa(now.Day())
	hourString := strconv.Itoa(now.Hour())

	// 构造文件路径和文件名
	filePath := filepath.Join(LogPath, yearString, monthString)
	fileName := ""
	if ifIncludeHour {
		fileName = fmt.Sprintf("%s-%s-%s-%s.%s.%s", yearString, monthString, dayString, hourString, level, "txt")
	} else {
		fileName = fmt.Sprintf("%s-%s-%s.%s.%s", yearString, monthString, dayString, level, "txt")
	}

	// 得到最终的fileName
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	LogMutex.Lock()
	if !isDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}
	LogMutex.Unlock()

	// 组装所有需要写入的内容
	content := fmt.Sprintf("%s---->\n", timeUtil.Format(now, "yyyy-MM-dd HH:mm:ss"))
	content += fmt.Sprintf("%s\n", logInfo)

	// 如果是Error类型，则同时记录堆栈信息
	if level == Error {
		for skip := 0; skip <= MAX_SKIP; skip++ {
			pc, file, line, ok := runtime.Caller(skip)
			if ok {
				content += fmt.Sprintf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
			}
		}
	}

	// 加上最后的分隔符
	content += SEPERATOR

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以读写模式打开。)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer f.Close()

	// 写入内容
	LogMutex.Lock()
	f.WriteString(content)
	LogMutex.Unlock()
}
