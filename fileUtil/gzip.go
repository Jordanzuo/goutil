package fileUtil

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 对文件进行gzip压缩
// source:源文件完整路径
// target:目标文件文件夹（如果传空字符串，则为当前文件夹）
// 返回值
// 错误对象
func Gzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 给目标文件夹赋值，如果传空，则默认为当前文件夹
	if target == "" {
		target = filepath.Dir(source)
	}
	fileName := filepath.Base(source)

	targetFilePath := filepath.Join(target, fmt.Sprintf("%s.gz", fileName))
	writer, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = fileName
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)

	return err
}

// 对文件进行gzip解压缩
// source:源文件完整路径
// target:目标文件文件夹（解压缩文件的名字是内部自动赋值）
// 返回值
// 错误对象
func UnGzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	// 给目标文件夹赋值，如果传空，则默认为当前文件夹
	if target == "" {
		target = filepath.Dir(source)
	}

	targetFilePath := filepath.Join(target, archive.Name)
	writer, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)

	return err
}
