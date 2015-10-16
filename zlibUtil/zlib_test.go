package zlibUtil

import (
	"compress/zlib"
	"testing"
)

var (
	InitString    = `{"Code":4,"Message":"IPNotAuthorized","Data":null}`
	CompressBytes []byte
)

func TestCompress(t *testing.T) {
	data := ([]byte)(InitString)
	result, _ := Compress(data, zlib.DefaultCompression)
	// if isEqual(result, CompressBytes) == false {
	// 	t.Errorf("压缩失败，期待%v，实际%v", InitBytes, result)
	// }

	CompressBytes = result
}

func TestDecompress(t *testing.T) {
	data, _ := Decompress(CompressBytes)
	result := string(data)
	if result != InitString {
		t.Errorf("解压缩失败，期待%s，实际%s", InitString, result)
	}
}

func isEqual(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
