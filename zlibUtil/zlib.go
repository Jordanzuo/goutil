package zlibUtil

import (
	"bytes"
	"compress/zlib"
	"io"
)

// 压缩
// in：待压缩数据
// level：压缩等级
// 返回值：
// 压缩后数据
// 对应的错误
func Compress(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	compressor, err := zlib.NewWriterLevelDict(&buffer, level, nil)
	if err != nil {
		return nil, err
	}

	compressor.Write(data)
	compressor.Close()

	return buffer.Bytes(), nil
}

// 解压缩
// in:待解压缩数据
// 返回值：
// 解压缩后数据
// 对应的错误
func Decompress(data []byte) ([]byte, error) {
	dataReader := bytes.NewReader(data)
	zlibReader, err := zlib.NewReader(dataReader)
	if err != nil {
		return nil, err
	}
	defer zlibReader.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, zlibReader)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
