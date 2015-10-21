package logUtil

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
	"time"
)

const (
	FILE_SUFFIX = "txt"
	SEPERATOR   = "------------------------------------------------------"
	MIN_SKIP    = 1
	MAX_SKIP    = 10
)

var (
	mLogPath     string
	mLogObjectCh = make(chan *LogObject, 128)
)

func init() {
	go flushLog()
}

func flushLog() {
	for {
		select {
		case logObj := <-mLogObjectCh:
			writeLog(logObj)
		default:
			// 休眠一下，防止CPU过高
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func isDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

func writeLog(logObj *LogObject) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	// 获取当前时间
	now := time.Now()

	// 获得年、月、日、时的字符串形式
	yearString := strconv.Itoa(now.Year())
	monthString := strconv.Itoa(int(now.Month()))
	dayString := strconv.Itoa(now.Day())
	hourString := strconv.Itoa(now.Hour())

	// 构造文件路径和文件名
	filePath := filepath.Join(mLogPath, yearString, monthString)
	fileName := ""
	if logObj.IfIncludeHour {
		fileName = fmt.Sprintf("%s-%s-%s-%s.%s.%s", yearString, monthString, dayString, hourString, logObj.Level, FILE_SUFFIX)
	} else {
		fileName = fmt.Sprintf("%s-%s-%s.%s.%s", yearString, monthString, dayString, logObj.Level, FILE_SUFFIX)
	}

	// 得到最终的fileName
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	if !isDirExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
	}

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以读写模式打开。)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer f.Close()

	// 写入内容
	f.WriteString(logObj.LogInfo)
}

// 设置日志存放的路径
// path：日志文件存放路径
func SetLogPath(path string) {
	mLogPath = path
}

// 获取日志文件存放路径
// 返回值：日志文件存放路径
func GetLogPath() string {
	return mLogPath
}

// 记录日志
// logInfo：需要记录的日志信息
// level：日志级别
// ifIncludeHour：日志文件名称是否包含小时
// 返回值：无
func Log(logInfo string, level LogType, ifIncludeHour bool) {
	// 判断路径是否为空
	if mLogPath == "" {
		panic(errors.New("日志存放路径不能为空，请先设置"))
	}

	// 获取当前时间
	now := time.Now()

	// 组装所有需要写入的内容
	newLogInfo := fmt.Sprintf("%s---->", timeUtil.Format(now, "yyyy-MM-dd HH:mm:ss"))
	newLogInfo += stringUtil.GetNewLineString()
	newLogInfo += fmt.Sprintf("%s", logInfo)
	newLogInfo += stringUtil.GetNewLineString()

	// 加上最后的分隔符
	newLogInfo += SEPERATOR
	newLogInfo += stringUtil.GetNewLineString()

	// 构造对象并添加到队列中
	logObj := NewLogObject(newLogInfo, level, ifIncludeHour)
	mLogObjectCh <- logObj
}

// 记录未知错误日志
// r：recover对象
// 返回值：无
func LogUnknownError(r interface{}, args ...string) {
	// 获取当前时间
	now := time.Now()

	// 组装所有需要写入的内容
	logInfo := fmt.Sprintf("%s---->", timeUtil.Format(now, "yyyy-MM-dd HH:mm:ss"))
	logInfo += stringUtil.GetNewLineString()
	logInfo += fmt.Sprintf("通过recover捕捉到的未处理异常：%v", r)
	logInfo += stringUtil.GetNewLineString()

	// 获取附加信息
	if len(args) > 0 {
		logInfo += fmt.Sprintf("附加信息：%s", strings.Join(args, "-"))
		logInfo += stringUtil.GetNewLineString()
	}

	// 获取堆栈信息
	for skip := MIN_SKIP; skip <= MAX_SKIP; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		logInfo += fmt.Sprintf("skip = %d, file = %s, line = %d", skip, file, line)
		logInfo += stringUtil.GetNewLineString()
	}

	// 加上最后的分隔符
	logInfo += SEPERATOR
	logInfo += stringUtil.GetNewLineString()

	// 构造对象并添加到队列中
	logObj := NewLogObject(logInfo, Error, true)
	mLogObjectCh <- logObj
}
