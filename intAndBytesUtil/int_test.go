package intAndBytesUtil

import (
	"encoding/binary"
	"testing"
)

var (
	ExpectedBigEndian    []byte
	ExpectedLittleEndian []byte
	GivenInt             int = 256
)

func init() {
	ExpectedBigEndian = []byte{0, 0, 1, 0}
	ExpectedLittleEndian = []byte{0, 1, 0, 0}
}

func TestIntToBytes(t *testing.T) {
	result := IntToBytes(GivenInt, binary.BigEndian)
	if equal(result, ExpectedBigEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", GivenInt, result, ExpectedBigEndian)
	}

	result = IntToBytes(GivenInt, binary.LittleEndian)
	if equal(result, ExpectedLittleEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", GivenInt, result, ExpectedLittleEndian)
	}
}

func equal(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}
