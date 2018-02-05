package fileUtil

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// 对一组文件进行tar打包
// sourceList:源文件完整路径列表
// target:目标文件名称
// 返回值
// 错误对象
func Tar(sourceList []string, target string) error {
	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tarball := tar.NewWriter(tarFile)
	defer tarball.Close()

	// 对源文件遍历处理
	for _, item := range sourceList {
		info, err := os.Stat(item)
		if err != nil || info.IsDir() {
			continue
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		header.Name = filepath.Base(item)

		if err := tarball.WriteHeader(header); err != nil {
			return err
		}

		file, err := os.Open(item)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = io.Copy(tarball, file); err != nil {
			return err
		}
	}

	return nil
}

// 对一组文件进行tar解包
// source:源文件完整路径
// target:目标文件名称
// 返回值
// 错误对象
func Untar(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	// 给目标文件夹赋值，如果传空，则默认为当前文件夹
	if target == "" {
		target = filepath.Dir(source)
	}

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		targetFilePath := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(targetFilePath, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}

	return nil
}
