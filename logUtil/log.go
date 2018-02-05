package logUtil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/fileUtil"
	"github.com/Jordanzuo/goutil/stringUtil"
	"github.com/Jordanzuo/goutil/timeUtil"
)

const (
	con_FILE_SUFFIX = "txt"
	con_SEPERATOR   = "------------------------------------------------------"
	con_MIN_SKIP    = 1
	con_MAX_SKIP    = 10
)

var (
	logPath = "DefaultLogPath"

	// 用于控制压缩文件的并发逻辑
	mutex sync.Mutex

	// 上一次日志压缩的日期 时间戳
	preCompressDate int64

	// 压缩锁对象
	compressLock sync.Mutex
)

// 日志压缩
func compress() {
	defer func() {
		if r := recover(); r != nil {
			// 将错误输出，而不是记录到文件，是因为可能导致死循环
			fmt.Println(r)
		}
	}()

	// 获取昨天的日期，并获取昨天对应的文件夹
	yesterday := time.Now().AddDate(0, 0, -1)
	dateString := timeUtil.Format(yesterday, "yyyy-MM-dd")
	fileAbsoluteDirectory := filepath.Join(logPath, strconv.Itoa(yesterday.Year()), strconv.Itoa(int(yesterday.Month())))

	mutex.Lock()
	defer mutex.Unlock()

	// 判断是否已经存在压缩文件
	compressFileName := fmt.Sprintf("%s.tar.gz", dateString)
	compressAbsolutePath := filepath.Join(fileAbsoluteDirectory, compressFileName)
	if exists, err := fileUtil.IsFileExists(compressAbsolutePath); err == nil && exists {
		return
	}

	// 获取昨天的文件列表
	fileList, err := fileUtil.GetFileList2(fileAbsoluteDirectory, dateString, con_FILE_SUFFIX)
	if err != nil {
		fmt.Printf("logUtil.compress.fileUtil.GetFileList2 err:%s\n", err)
		return
	}
	if len(fileList) == 0 {
		return
	}

	// 进行tar操作，得到yyyy-MM-dd.tar
	tarFileName := fmt.Sprintf("%s.tar", dateString)
	tarAbsolutePath := filepath.Join(fileAbsoluteDirectory, tarFileName)
	if err := fileUtil.Tar(fileList, tarAbsolutePath); err != nil {
		fmt.Printf("logUtil.compress.fileUtil.Tar err:%s\n", err)
	}

	// 进行gzip操作，得到yyyy-MM-dd.tar.gz
	if err := fileUtil.Gzip(tarAbsolutePath, ""); err != nil {
		fmt.Printf("logUtil.compress.fileUtil.Gzip err:%s\n", err)
	}

	// 删除原始文件
	for _, item := range fileList {
		fileUtil.DeleteFile(item)
	}

	// 删除tar文件
	fileUtil.DeleteFile(tarAbsolutePath)
}

// 写日志
// logObj:日志对象
func writeLog(logObj *logObject) {
	if logObj.level == Warn || logObj.level == Error || logObj.level == Fatal {
		debugUtil.Println(logObj.logInfo)
	}

	// 获取当前时间
	now := time.Now()
	fileAbsoluteDirectory := filepath.Join(logPath, strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())))
	fileName := ""
	fileAbsolutePath := ""

	if logObj.ifIncludeHour {
		fileName = fmt.Sprintf("%s.%s.%s", timeUtil.Format(now, "yyyy-MM-dd-HH"), logObj.level, con_FILE_SUFFIX)
	} else {
		fileName = fmt.Sprintf("%s.%s.%s", timeUtil.Format(now, "yyyy-MM-dd"), logObj.level, con_FILE_SUFFIX)
	}

	// 得到最终的文件绝对路径
	fileAbsolutePath = filepath.Join(fileAbsoluteDirectory, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	if !fileUtil.IsDirExists(fileAbsoluteDirectory) {
		if err := os.MkdirAll(fileAbsoluteDirectory, os.ModePerm|os.ModeTemporary); err != nil {
			return
		}
	}

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以写模式打开。)
	f, err := os.OpenFile(fileAbsolutePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return
	}
	defer f.Close()

	// 写入内容
	f.WriteString(logObj.logInfo)

	// 检查是否需要进行数据压缩
	nowDate := timeUtil.GetDate(now).Unix()
	if nowDate == preCompressDate {
		return
	}

	compressLock.Lock()
	defer compressLock.Unlock()

	// 上一次压缩的时间
	if nowDate == preCompressDate {
		return
	}
	preCompressDate = nowDate

	// 日志压缩
	go compress()
}

// 设置日志存放的路径
// _logPath：日志文件存放路径
func SetLogPath(_logPath string) {
	logPath = _logPath
}

// 获取日志文件存放路径
// 返回值：日志文件存放路径
func GetLogPath() string {
	return logPath
}

// 记录日志
// logInfo：需要记录的日志信息
// level：日志级别
// ifIncludeHour：日志文件名称是否包含小时
// 返回值：无
func Log(logInfo string, level logType, ifIncludeHour bool) {
	// 获取当前时间
	now := time.Now()

	// 组装所有需要写入的内容
	newLogInfo := fmt.Sprintf("%s---->", timeUtil.Format(now, "yyyy-MM-dd HH:mm:ss"))
	newLogInfo += stringUtil.GetNewLineString()
	newLogInfo += fmt.Sprintf("%s", logInfo)
	newLogInfo += stringUtil.GetNewLineString()

	// 加上最后的分隔符
	newLogInfo += con_SEPERATOR
	newLogInfo += stringUtil.GetNewLineString()

	// 构造对象并添加到队列中
	writeLog(newLogObject(newLogInfo, level, ifIncludeHour))
}

// 常规的日志记录接口(ifIncludeHour=true)
// logInfo：需要记录的日志信息
// level：日志级别
func NormalLog(logInfo string, level logType) {
	Log(logInfo, level, true)
}

// 记录到文件并且答应到控制台
// logInfo：需要记录的日志信息
// level：日志级别
func LogAndPrint(logInfo string, level logType) {
	NormalLog(logInfo, level)
	log.Print(logInfo)
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
	for skip := con_MIN_SKIP; skip <= con_MAX_SKIP; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		logInfo += fmt.Sprintf("skip = %d, file = %s, line = %d", skip, file, line)
		logInfo += stringUtil.GetNewLineString()
	}

	// 加上最后的分隔符
	logInfo += con_SEPERATOR
	logInfo += stringUtil.GetNewLineString()

	// 构造对象并添加到队列中
	writeLog(newLogObject(logInfo, Error, true))
}
