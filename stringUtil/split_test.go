package stringUtil

import (
	"fmt"
	"testing"
)

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

func TestSplitToFloat64(t *testing.T) {
	result, err := SplitToFloat64Slice("1.11,2.22", ",")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v\n", result)
}
