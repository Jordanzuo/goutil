package intAndBytesUtil

import (
	"encoding/binary"
	"testing"
)

func TestInt16ToBytes(t *testing.T) {
	var expectedBigEndian []byte = []byte{1, 0}
	var expectedLittleEndian []byte = []byte{0, 1}
	var givenInt int16 = 256

	result := Int16ToBytes(givenInt, binary.BigEndian)
	if equal(result, expectedBigEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedBigEndian)
	}

	result = Int16ToBytes(givenInt, binary.LittleEndian)
	if equal(result, expectedLittleEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedLittleEndian)
	}
}

func TestInt32ToBytes(t *testing.T) {
	var expectedBigEndian []byte = []byte{0, 0, 1, 0}
	var expectedLittleEndian []byte = []byte{0, 1, 0, 0}
	var givenInt int32 = 256

	result := Int32ToBytes(givenInt, binary.BigEndian)
	if equal(result, expectedBigEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedBigEndian)
	}

	result = Int32ToBytes(givenInt, binary.LittleEndian)
	if equal(result, expectedLittleEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedLittleEndian)
	}
}

func TestInt64ToBytes(t *testing.T) {
	var expectedBigEndian []byte = []byte{0, 0, 0, 0, 0, 0, 1, 0}
	var expectedLittleEndian []byte = []byte{0, 1, 0, 0, 0, 0, 0, 0}
	var givenInt int64 = 256

	result := Int64ToBytes(givenInt, binary.BigEndian)
	if equal(result, expectedBigEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedBigEndian)
	}

	result = Int64ToBytes(givenInt, binary.LittleEndian)
	if equal(result, expectedLittleEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedLittleEndian)
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
