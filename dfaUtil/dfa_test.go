package dfaUtil

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	sensitiveList := []string{"中国", "中国人"}
	input := "我来自中国cd"

	util := newDFAUtil(sensitiveList)
	if util.IsMatch(input) == false {
		t.Errorf("Expected true, but got false")
	}
}

func TestHandleWord(t *testing.T) {
	sensitiveList := []string{"中国", "中国人", "来自", "习近平", "会议"}
	input := "我来自中国cd，习近平出席了会议"

	util := newDFAUtil(sensitiveList)
	newInput := util.HandleWord(input, '*')
	expected := "我****cd，***出席了**"
	if newInput != expected {
		t.Errorf("Expected %s, but got %s", expected, newInput)
	}
}
