package stringUtil

import (
	"testing"
)

// test IsGUIDEmpty
func TestIsGUIDEmpty(t *testing.T) {
	isOk := IsGUIDEmpty("")
	if isOk == false {
		t.Error("test is Not pass:")
		return
	}

	isOk = IsGUIDEmpty("00000000-0000-0000-0000-000000000000")
	if isOk == false {
		t.Error("test is Not pass:00000000-0000-0000-0000-000000000000")
		return
	}

	isOk = IsGUIDEmpty("00000000-0000-0000-0000-000000000001")
	if isOk == true {
		t.Error("test is Not pass:00000000-0000-0000-0000-000000000001")
		return
	}
}
