package fileUtil

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Batch file is used to write data to hard disk sequentially.
// The time to flush data to hard disk is determined either by size or time.
type BatchFile struct {
	DirectoryName   string
	FileName        string
	timeIntervalSec int64
	currTimeSec     int64
	maxTimeSec      int64
	currFileSize    int64
	maxFileSize     int64
	file            *os.File // The underlying file object
}

func NewBatchFile(directoryName, fileName string, size, intervalSec int64) (*BatchFile, error) {
	if !IsDirExists(directoryName) {
		os.MkdirAll(directoryName, os.ModePerm|os.ModeTemporary)
	}

	batchFilePtr := &BatchFile{
		DirectoryName:   directoryName,
		FileName:        fileName,
		timeIntervalSec: intervalSec,
		maxFileSize:     size,
	}

	err := batchFilePtr.initFile()
	if err != nil {
		batchFilePtr.Close()
		return nil, err
	}

	return batchFilePtr, nil
}

func (this *BatchFile) getFilePath() string {
	return filepath.Join(this.DirectoryName, this.FileName)
}

func (this *BatchFile) getBakFilePath() string {
	return filepath.Join(this.DirectoryName, fmt.Sprintf("%s_%d", this.FileName, time.Now().UnixNano()))
}

func (this *BatchFile) initFile() error {
	// Reset fields
	this.currTimeSec = time.Now().Unix()
	this.maxTimeSec = time.Now().Unix() + this.timeIntervalSec
	this.currFileSize = 0

	// Open file
	file, err := os.OpenFile(this.getFilePath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("Open file: %s failed. Error: %v", this.getFilePath(), err)
	}

	this.file = file
	return nil
}

func (this *BatchFile) checkNewFile() {
	// Close current file and open a new file when either condition is met
	if this.currTimeSec >= this.maxTimeSec || this.currFileSize >= this.maxFileSize {
		// Sync file
		this.file.Sync()
		this.file.Close()

		// Rename file
		oldPath := this.getFilePath()
		newPath := this.getBakFilePath()
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println(err)
		}
		//os.Rename(this.getFilePath(), this.getBakFilePath())

		// Open a new file
		this.initFile()
	}
}

func (this *BatchFile) WriteString(message string) error {
	// Update statistics information
	this.currTimeSec = time.Now().Unix()
	this.currFileSize += int64(len([]byte(message)))

	// Write message, add \n at the end
	message = fmt.Sprintf("%s\n", message)
	_, err := this.file.WriteString(message)
	if err != nil {
		return fmt.Errorf("file.WriteString to %s failed. Error: %v", this.getFilePath(), err)
	}

	// Check whether close current file and open a new file
	this.checkNewFile()

	return nil
}

func (this *BatchFile) WriteBytes(message []byte) error {
	// Update statistics information
	this.currTimeSec = time.Now().Unix()
	this.currFileSize += int64(len(message))

	_, err := this.file.Write(message)
	if err != nil {
		return fmt.Errorf("file.WriteString to %s failed. Error: %v", this.getFilePath(), err)
	}

	// Check whether close current file and open a new file
	this.checkNewFile()

	return nil
}

func (this *BatchFile) Close() {
	if this.file != nil {
		this.file.Close()
		this.file = nil
	}
}
