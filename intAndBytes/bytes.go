package intAndBytes

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

	var x int32
	binary.Read(bytesBuffer, order, &x)

	return int(x)
}
