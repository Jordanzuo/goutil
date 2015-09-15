package stringUtil

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	str := "Hello, Jordan.左贤清"
	substr := Substr(str, 0, 5)
	expectedstr := "Hello"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substr(str, 0, 10)
	expectedstr = "Hello, Jor"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substr(str, 0, 15)
	expectedstr = "Hello, Jordan.左"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}

	substr = Substr(str, 0, 20)
	expectedstr = "Hello, Jordan.左贤清"

	if substr != expectedstr {
		t.Errorf("Failed. Expected:%s, Got %s\n", expectedstr, substr)
	}
}
