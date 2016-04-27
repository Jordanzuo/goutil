package intAndBytesUtil

import (
	"bytes"
	"encoding/binary"
)

// 字节数组转换成整形
// b：字节数组
// order：大、小端的枚举
// 返回值：对应的int值
func BytesToInt(b []byte, order binary.ByteOrder) int {
	bytesBuffer := bytes.NewBuffer(b)

	var result int
	binary.Read(bytesBuffer, order, &result)

	return result
}

// 字节数组转换成整形
// b：字节数组
// order：大、小端的枚举
// 返回值：对应的int16值
func BytesToInt16(b []byte, order binary.ByteOrder) int16 {
	bytesBuffer := bytes.NewBuffer(b)

	var result int16
	binary.Read(bytesBuffer, order, &result)

	return result
}

// 字节数组转换成整形
// b：字节数组
// order：大、小端的枚举
// 返回值：对应的int32值
func BytesToInt32(b []byte, order binary.ByteOrder) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var result int32
	binary.Read(bytesBuffer, order, &result)

	return result
}

// 字节数组转换成整形
// b：字节数组
// order：大、小端的枚举
// 返回值：对应的int64值
func BytesToInt64(b []byte, order binary.ByteOrder) int64 {
	bytesBuffer := bytes.NewBuffer(b)

	var result int64
	binary.Read(bytesBuffer, order, &result)

	return result
}
