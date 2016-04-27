package intAndBytesUtil

import (
	"bytes"
	"encoding/binary"
)

// 整形转换成字节(无效，因为系统无法判断读取的字节数)
// n：int型数字
// order：大、小端的枚举
// 返回值：对应的字节数组
// func IntToBytes(n int, order binary.ByteOrder) []byte {
// 	bytesBuffer := bytes.NewBuffer([]byte{})
// 	binary.Write(bytesBuffer, order, n)

// 	return bytesBuffer.Bytes()
// }

// 整形转换成字节
// n：int16型数字
// order：大、小端的枚举
// 返回值：对应的字节数组
func Int16ToBytes(n int16, order binary.ByteOrder) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, n)

	return bytesBuffer.Bytes()
}

// 整形转换成字节
// n：int32型数字
// order：大、小端的枚举
// 返回值：对应的字节数组
func Int32ToBytes(n int32, order binary.ByteOrder) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, n)

	return bytesBuffer.Bytes()
}

// 整形转换成字节
// n：int64型数字
// order：大、小端的枚举
// 返回值：对应的字节数组
func Int64ToBytes(n int64, order binary.ByteOrder) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, n)

	return bytesBuffer.Bytes()
}
