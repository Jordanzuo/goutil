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

func TestSplitToIntSlice(t *testing.T) {
	s := "1, 2, 3, 4, 5, a"
	if _, err := SplitToIntSlice(s, ","); err == nil {
		t.Errorf("Expected got err, but got nil")
	}

	s = "1, 5, 39,"
	if intSlice, err := SplitToIntSlice(s, ","); err != nil {
		t.Errorf("Expected got nil, but got error:%s", err)
	} else {
		// fmt.Printf("intSlice:%v\n", intSlice)
		if intSlice[0] != 1 || intSlice[1] != 5 || intSlice[2] != 39 {
			t.Errorf("Expected got %s, but got %v", s, intSlice)
		}
	}
}
