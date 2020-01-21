package logUtil

import (
	"fmt"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/Jordanzuo/goutil/fileUtil"
	"github.com/Jordanzuo/goutil/timeUtil"
)

var (
	// 上一次日志压缩的日期 时间戳
	preCompressDate int64

	// 压缩锁对象
	compressLock sync.Mutex
)

func compress() {
	// 检查是否需要进行数据压缩
	nowDate := timeUtil.GetDate(time.Now()).Unix()
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
	go doCompress()
}

// 日志压缩
func doCompress() {
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
