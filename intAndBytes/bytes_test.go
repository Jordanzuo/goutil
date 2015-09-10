package intAndBytes

import (
	"encoding/binary"
	"testing"
)

var (
	GivenBigEndian    []byte
	GivenLittleEndian []byte
	ExpectedInt       int = 256
)

func init() {
	GivenBigEndian = []byte{0, 0, 1, 0}
	GivenLittleEndian = []byte{0, 1, 0, 0}
}

func TestBytesToInt(t *testing.T) {
	result := BytesToInt(GivenBigEndian, binary.BigEndian)
	if result != ExpectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", GivenBigEndian, result, ExpectedInt)
	}

	result = BytesToInt(GivenLittleEndian, binary.LittleEndian)
	if result != ExpectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", GivenLittleEndian, result, ExpectedInt)
	}
}
