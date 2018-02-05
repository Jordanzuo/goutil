package stringUtil

import (
	"testing"
)

// test 特殊字符
func TestEmoji1(t *testing.T) {
	tstVal := map[string]string{
		"中文":   "你好啊",
		"繁体中文": "這是什麼天氣",
		"泰文":   "สวัสดีครับ !",
		"英文":   "helloworld",
		"越南语":  "Đó là gì thời tiết.",
		"日语":   "これは何の天気ですか",
		"标点符号": "!@#$%^^&*())(__+{}[]|:<>",
	}

	for key, val := range tstVal {
		if IfHaveEmoji(val) {
			t.Errorf("语言处理错误：%s", key)
		}
	}

	emojiVal := "☀"
	if IfHaveEmoji(emojiVal) == false {
		t.Errorf("表情符号匹配错误:")
	}

	specialChar := "\\'\""
	if IfHaveEmoji(specialChar) {
		t.Errorf("特殊字符匹配错误:")
	}
}
