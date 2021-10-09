package rangeUtil

import (
	"testing"
)

func TestNewIntRange(t *testing.T) {
	input := ""
	_, err := NewIntRange(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "{1,5}"
	_, err = NewIntRange(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "{1,5]"
	_, err = NewIntRange(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "[1,100]"
	intRange, err := NewIntRange(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType := InclusiveInclusive
	getRangeType := intRange.RangeType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower := int64(1)
	getLower := intRange.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper := int64(100)
	getUpper := intRange.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid := false
	getValid := intRange.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "[1,100)"
	intRange, err = NewIntRange(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = InclusiveExclusive
	getRangeType = intRange.RangeType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intRange.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intRange.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "(1,100]"
	intRange, err = NewIntRange(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = ExclusiveInclusive
	getRangeType = intRange.RangeType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intRange.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intRange.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "(1,100)"
	intRange, err = NewIntRange(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = ExclusiveExclusive
	getRangeType = intRange.RangeType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intRange.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intRange.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get RangeType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intRange.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intRange.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}
}
