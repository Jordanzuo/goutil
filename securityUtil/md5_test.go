package securityUtil

import (
	"testing"
)

var (
	ExpectedUpperString = "5EB63BBBE01EEED093CB22BB8F5ACDC3"
	ExpectedLowerString = "5eb63bbbe01eeed093cb22bb8f5acdc3"
)

func TestMd5String(t *testing.T) {
	s := "hello world"
	result := Md5String(s, true)
	if result != ExpectedUpperString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedUpperString)
	}

	result = Md5String(s, false)
	if result != ExpectedLowerString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedLowerString)
	}
}

func TestMd5Bytes(t *testing.T) {
	s := "hello world"
	b := []byte(s)
	result := Md5Bytes(b, true)
	if result != ExpectedUpperString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedUpperString)
	}

	result = Md5Bytes(b, false)
	if result != ExpectedLowerString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedLowerString)
	}
}
