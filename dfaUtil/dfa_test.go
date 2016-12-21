package dfaUtil

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	sensitiveList := []string{"中国", "中国人"}
	input := "我来自中国cd"

	util := NewDFAUtil(sensitiveList)
	if util.IsMatch(input) == false {
		t.Errorf("Expected true, but got false")
	}
}

func TestHandleWord(t *testing.T) {
	sensitiveList := []string{"中国", "中国人", "习近平的", "会议", "近平"}
	input := "我来自中国cd，习近平出席了会议,习近平"

	util := NewDFAUtil(sensitiveList)
	newInput := util.HandleWord(input, '*')
	expected := "我来自**cd，习**出席了**,习**"
	if newInput != expected {
		t.Errorf("Expected %s, but got %s", expected, newInput)
	}
}
