package fileUtil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Batch file is used to write data to hard disk sequentially.
// The time to flush data to hard disk is determined either by size or time.
type BatchFile struct {
	DirectoryName   string
	FileName        string
	FileSuffix      string
	timeIntervalSec int64
	maxTimeSec      int64
	currFileSize    int64
	maxFileSize     int64
	file            *os.File   // The underlying file object
	mutex           sync.Mutex // Protects
	closingChannel  chan bool
	closedChannel   chan bool
}

func NewBatchFile(directoryName, fileName, fileSuffix string, size, intervalSec int64) (*BatchFile, error) {
	if !IsDirExists(directoryName) {
		os.MkdirAll(directoryName, os.ModePerm|os.ModeTemporary)
	}

	batchFilePtr := &BatchFile{
		DirectoryName:   directoryName,
		FileName:        fileName,
		FileSuffix:      fileSuffix,
		timeIntervalSec: intervalSec,
		maxFileSize:     size,
		closingChannel:  make(chan bool, 1),
		closedChannel:   make(chan bool, 1),
	}

	err := batchFilePtr.initFile()
	if err != nil {
		batchFilePtr.Close()
		return nil, err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("checkNewFile failed. Error: %v", r)
			}
		}()

		tick := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-tick.C:
				func() {
					batchFilePtr.mutex.Lock()
					defer batchFilePtr.mutex.Unlock()
					batchFilePtr.checkNewFile()
					log.Println("checkNewFile periodically")
				}()
			case <-batchFilePtr.closingChannel:
				log.Println("quit periodical check goroutine")
				batchFilePtr.closedChannel <- true
				return
			}
		}
	}()

	return batchFilePtr, nil
}

func (this *BatchFile) getFilePrefix() string {
	return fmt.Sprintf("%s_", this.FileName)
}

func (this *BatchFile) getFilePath() string {
	return filepath.Join(this.DirectoryName, fmt.Sprintf("%s.%s", this.FileName, this.FileSuffix))
}

func (this *BatchFile) getBakFilePath() string {
	return filepath.Join(this.DirectoryName, fmt.Sprintf("%s_%d.%s", this.FileName, time.Now().UnixNano(), this.FileSuffix))
}

func (this *BatchFile) GetBakFilePathList() ([]string, error) {
	return GetFileList2(this.DirectoryName, this.getFilePrefix(), this.FileSuffix)
}

func (this *BatchFile) initFile() error {
	// Reset fields
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
	if time.Now().Unix() >= this.maxTimeSec || this.currFileSize >= this.maxFileSize {
		// Sync file
		this.file.Sync()
		this.file.Close()

		// Rename file
		oldPath := this.getFilePath()
		newPath := this.getBakFilePath()
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("checkNewFile.os.Rename failed. Error: %v", err)
		}
		log.Printf("Create a bak file: %s\n", newPath)

		// Open a new file
		this.initFile()
	}
}

func (this *BatchFile) WriteString(message string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// Update statistics information
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
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// Update statistics information
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

	this.closingChannel <- true
	<-this.closedChannel
	log.Printf("BatchFile.Close")
}
