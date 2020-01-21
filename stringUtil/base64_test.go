package stringUtil

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	greeting := "Hello world"
	encoded := Base64Encode(greeting)
	decoded, err := Base64Decode(encoded)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if greeting != decoded {
		t.Errorf("Expected %s, but got %s", greeting, decoded)
		return
	}
}

func TestBase64Encode2(t *testing.T) {
	greeting := []byte("Hello world")
	encoded := Base64Encode2(greeting)
	decoded, err := Base64Decode2(encoded)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if isEqual(greeting, decoded) == false {
		t.Errorf("Expected %s, but got %s", greeting, decoded)
		return
	}
}

func isEqual(s1, s2 []byte) bool {
	if s1 == nil || s2 == nil {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
