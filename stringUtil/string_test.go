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
	fmt.Printf("guid1:%s, guid2:%s,", guid1, guid2)
	fmt.Printf("length of %s is %d\n", guid1, len(guid1))
	if guid1 == guid2 {
		t.Errorf("%s should not be equal with %s", guid1, guid2)
	}
}

func TestSplit(t *testing.T) {
	s := "1,2;3|4||5,6;7|8||9,"
	// seps := []string{",", ";", "|", "||"}
	retList := Split(s, nil)
	if retList[0] != "1" || retList[1] != "2" || retList[2] != "3" || retList[3] != "4" || retList[4] != "5" || retList[5] != "6" || retList[6] != "7" || retList[7] != "8" || retList[8] != "9" {
		t.Errorf("ecptected:123456789, but now got:%v", retList)
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

func TestSplitToIntRegion(t *testing.T) {
	idRegionStr := ""
	outerSep := ","
	innerSep := "-"
	var err error

	if _, err = SplitToIntRegion(idRegionStr, outerSep, innerSep); err == nil {
		t.Errorf("PraseIdRegion should failed, but now not.err:%s", err)
	}

	idRegionStr = ","
	if _, err = SplitToIntRegion(idRegionStr, outerSep, innerSep); err == nil {
		t.Errorf("PraseIdRegion should failed, but now not.err:%s", err)
	}

	idRegionStr = "1-100,101,200"
	if _, err = SplitToIntRegion(idRegionStr, outerSep, innerSep); err == nil {
		t.Errorf("PraseIdRegion should failed, but now not.err:%s", err)
	}

	idRegionStr = "1-100,101-20"
	if _, err = SplitToIntRegion(idRegionStr, outerSep, innerSep); err == nil {
		t.Errorf("PraseIdRegion should failed, but now not.err:%s", err)
	}

	idRegionStr = "1-100,101-200"
	if idRegionList, err := SplitToIntRegion(idRegionStr, outerSep, innerSep); err != nil {
		t.Errorf("PraseIdRegion should succeed, but now failed.err:%s", err)
	} else {
		if idRegionList[0].Lower != 1 || idRegionList[0].Upper != 100 ||
			idRegionList[1].Lower != 101 || idRegionList[1].Upper != 200 {
			t.Errorf("SplitToIntRegion should succeed, but now failed. idRegionStr:%s, idRegionList:%v", idRegionStr, idRegionList)
		}
	}
}

// test IsEmpty
func TestIsEmpty(t *testing.T) {
	isOk := IsEmpty("")
	if isOk == false {
		t.Error("\"\" test is Not pass")
		return
	}

	isOk = IsEmpty(" ")
	if isOk == false {
		t.Error("\" \" test is Not pass")
		return
	}

	isOk = IsEmpty(" \t\n")
	if isOk == false {
		t.Error("\" \\t\\n\" test is Not pass")
		return
	}
}

// test 特殊字符
func TestIfHaveSpecialChar(t *testing.T) {
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
		if IfHaveSpecialChar(val) {
			t.Errorf("语言处理错误：%s", key)
		}
	}

	specialChar := "\\'\""
	if IfHaveSpecialChar(specialChar) == false {
		t.Errorf("特殊字符匹配错误:")
	}
}

func TestIsDistinct_string(t *testing.T) {
	list := make([]string, 0, 8)
	result := IsDistinct_string(list)
	fmt.Printf("list:%v,result:%v-------1\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------1")
	}

	list = append(list, "Hello")
	result = IsDistinct_string(list)
	fmt.Printf("list:%v,result:%v-------2\n", list, result)
	if result == false {
		t.Errorf("it's should be true, but now false-------2")
	}

	list = append(list, "Hello")
	result = IsDistinct_string(list)
	fmt.Printf("list:%v,result:%v-------3\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------3")
	}

	list = append(list, "")
	result = IsDistinct_string(list)
	fmt.Printf("list:%v,result:%v-------4\n", list, result)
	if result {
		t.Errorf("it's should be false, but now true-------4")
	}
}
