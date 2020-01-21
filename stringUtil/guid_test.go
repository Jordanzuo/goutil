package stringUtil

import (
	"testing"
)

func TestGetNewGUID(t *testing.T) {
	guidMap := make(map[string]bool, 1024)
	count := 10
	for i := 0; i < count; i++ {
		guid := GetNewGUID()
		guidMap[guid] = true
	}

	if len(guidMap) != count {
		t.Errorf("there should be %d 条不重复的数据，但是现在只有%d条", count, len(guidMap))
	}
}

func TestGetEmptyGUID(t *testing.T) {
	guid := GetEmptyGUID()
	expected := "00000000-0000-0000-0000-000000000000"
	if guid != expected {
		t.Errorf("guid should be %s, but now is %s", expected, guid)
	}
}

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

func TestGetNewUUID(t *testing.T) {
	guidMap := make(map[string]bool, 1024)
	count := 10
	for i := 0; i < count; i++ {
		guid := GetNewUUID()
		guidMap[guid] = true
	}

	if len(guidMap) != count {
		t.Errorf("there should be %d 条不重复的数据，但是现在只有%d条", count, len(guidMap))
	}
}
