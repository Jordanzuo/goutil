package intAndBytesUtil

import (
	"bytes"
	"encoding/binary"
)

// 整形转换成字节
// n：int型数字
// order：大、小端的枚举
// 返回值：对应的字节数组
func IntToBytes(n int, order binary.ByteOrder) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, x)

	return bytesBuffer.Bytes()
}
