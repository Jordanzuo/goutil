package fileUtil

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 大文件对象，可用于连续写入内容而不关闭文件，直到达到指定的大小
type BigFile struct {
	// 文件夹名称
	path string

	// 当前文件名称
	fileName string

	// 当前文件大小
	fileSize int

	// 最大的文件大小
	maxFileSize int

	// 文件对象
	file *os.File

	// 获得新文件名称的方法
	newFileNameFunc func() string
}

// 获取文件的完整路径
func (this *BigFile) getFullPath() string {
	return filepath.Join(this.path, this.fileName)
}

// 初始化文件对象
func (this *BigFile) initFile() error {
	// 初始化文件名称
	this.fileName = this.newFileNameFunc()

	// 初始化文件大小
	this.fileSize = 0

	// 打开文件
	file, err := os.OpenFile(this.getFullPath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return fmt.Errorf("打开文件%s错误，错误信息为：%s", this.getFullPath(), err)
	} else {
		this.file = file
	}

	return nil
}

// 保存消息
// message：消息内容
// 返回值：无
func (this *BigFile) SaveMessage(message string) error {
	if this.file == nil {
		return fmt.Errorf("文件对象为空，path:%s", this.getFullPath())
	}

	// 增加文件大小
	this.fileSize += len([]byte(message))

	// 写入消息
	message = fmt.Sprintf("%s\n", message)
	if _, err := this.file.WriteString(message); err != nil {
		return fmt.Errorf("向文件%s写入信息错误，错误信息为：%s", this.getFullPath(), err)
	}

	// 如果达到了文件的上限，则关闭文件并重新打开一个新文件
	if this.fileSize >= this.maxFileSize {
		this.Close()
		this.initFile()
	}

	return nil
}

// 关闭对象
// 返回值：无
func (this *BigFile) Close() {
	if this.file != nil {
		this.file.Close()
		this.file = nil
	}
}

// 创建新的大文件对象
// _path:文件夹路径
// _maxFileSize:单个文件大小的最大值（单位：Byte）
// 返回值：
// 大文件对象
// 错误对象
func NewBigFile(_path string, _maxFileSize int) (*BigFile, error) {
	return NewBigFileWithNewFileNameFunc(_path, _maxFileSize, newFileName)
}

// 创建新的大文件对象
// _path:文件夹路径
// _maxFileSize:单个文件大小的最大值（单位：Byte）
// _newFileNameFunc:创建新文件名称的方法
// 返回值：
// 大文件对象
// 错误对象
func NewBigFileWithNewFileNameFunc(_path string, _maxFileSize int, _newFileNameFunc func() string) (*BigFile, error) {
	// 判断文件夹是否存在，如果不存在则创建
	if !IsDirExists(_path) {
		os.MkdirAll(_path, os.ModePerm|os.ModeTemporary)
	}

	// 初始化对象
	obj := &BigFile{
		path:            _path,
		maxFileSize:     _maxFileSize,
		newFileNameFunc: _newFileNameFunc,
	}

	// 初始化文件对象
	if err := obj.initFile(); err != nil {
		obj.Close()
		return nil, err
	}

	return obj, nil
}

// 创建新的文件名称
// 返回值：
// 新的文件名称
func newFileName() string {
	convertTime := func(num int) string {
		if num < 10 {
			return fmt.Sprintf("0%d", num)
		} else {
			return fmt.Sprintf("%d", num)
		}
	}

	now := time.Now()
	year := now.Year()
	month := convertTime(int(now.Month()))
	day := convertTime(now.Day())
	hour := convertTime(now.Hour())
	minute := convertTime(now.Minute())
	second := convertTime(now.Second())

	return fmt.Sprintf("%d%s%s%s%s%s.data", year, month, day, hour, minute, second)
}
