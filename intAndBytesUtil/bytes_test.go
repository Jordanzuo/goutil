package intAndBytesUtil

import (
	"encoding/binary"
	"testing"
)

func TestBytesToInt(t *testing.T) {
	var givenBigEndian []byte = []byte{0, 0, 1, 0}
	var givenLittleEndian []byte = []byte{0, 1, 0, 0}
	var expectedInt int32 = 256

	result := BytesToInt32(givenBigEndian, binary.BigEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenBigEndian, result, expectedInt)
	}

	result = BytesToInt32(givenLittleEndian, binary.LittleEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenLittleEndian, result, expectedInt)
	}
}

func TestBytesToInt16(t *testing.T) {
	var givenBigEndian []byte = []byte{1, 0}
	var givenLittleEndian []byte = []byte{0, 1}
	var expectedInt int16 = 256

	result := BytesToInt16(givenBigEndian, binary.BigEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenBigEndian, result, expectedInt)
	}

	result = BytesToInt16(givenLittleEndian, binary.LittleEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenLittleEndian, result, expectedInt)
	}
}

func TestBytesToInt32(t *testing.T) {
	var givenBigEndian []byte = []byte{0, 0, 1, 0}
	var givenLittleEndian []byte = []byte{0, 1, 0, 0}
	var expectedInt int32 = 256

	result := BytesToInt32(givenBigEndian, binary.BigEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenBigEndian, result, expectedInt)
	}

	result = BytesToInt32(givenLittleEndian, binary.LittleEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenLittleEndian, result, expectedInt)
	}
}

func TestBytesToInt64(t *testing.T) {
	var givenBigEndian []byte = []byte{0, 0, 0, 0, 0, 0, 1, 0}
	var givenLittleEndian []byte = []byte{0, 1, 0, 0, 0, 0, 0, 0}
	var expectedInt int64 = 256

	result := BytesToInt64(givenBigEndian, binary.BigEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenBigEndian, result, expectedInt)
	}

	result = BytesToInt64(givenLittleEndian, binary.LittleEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenLittleEndian, result, expectedInt)
	}
}
