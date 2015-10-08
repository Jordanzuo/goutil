package stringUtil

import (
	"fmt"
	"testing"
)

func TestSubstr(t *testing.T) {
	str := "Hello, Jordan.左贤清"
	substr := Substring(str, 0, 5)
	expectedstr := "Hello"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substring(str, 0, 10)
	expectedstr = "Hello, Jor"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substring(str, 0, 15)
	expectedstr = "Hello, Jordan.左"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substring(str, 0, 20)
	expectedstr = "Hello, Jordan.左贤清"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	guid1 := GetNewGUID()
	guid2 := GetNewGUID()
	fmt.Printf("guid1:%s, guid2:%s", guid1, guid2)
	fmt.Println("length of %s is %d", guid1, len(guid1))
	if guid1 == guid2 {
		t.Errorf("%s should not be equal with %s", guid1, guid2)
	}
}
