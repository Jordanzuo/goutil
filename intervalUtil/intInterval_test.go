package intervalUtil

import (
	"testing"
)

func TestNewIntInterval(t *testing.T) {
	input := ""
	_, err := NewIntInterval(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "{1,5}"
	_, err = NewIntInterval(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "{1,5]"
	_, err = NewIntInterval(input)
	if err == nil {
		t.Errorf("Expect to get an error, but now there isn't.")
		return
	}

	input = "[1,100]"
	intInterval, err := NewIntInterval(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType := InclusiveInclusive
	getRangeType := intInterval.IntervalType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower := int64(1)
	getLower := intInterval.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper := int64(100)
	getUpper := intInterval.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid := false
	getValid := intInterval.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "[1,100)"
	intInterval, err = NewIntInterval(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = InclusiveExclusive
	getRangeType = intInterval.IntervalType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intInterval.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intInterval.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "(1,100]"
	intInterval, err = NewIntInterval(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = ExclusiveInclusive
	getRangeType = intInterval.IntervalType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intInterval.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intInterval.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	input = "(1,100)"
	intInterval, err = NewIntInterval(input)
	if err != nil {
		t.Errorf("Expect to get no error, but now there is one.")
		return
	}

	expectRangeType = ExclusiveExclusive
	getRangeType = intInterval.IntervalType
	if getRangeType != expectRangeType {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectRangeType, getRangeType)
		return
	}

	expectLower = int64(1)
	getLower = intInterval.Lower
	if getLower != expectLower {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectLower, getLower)
		return
	}

	expectUpper = int64(100)
	getUpper = intInterval.Upper
	if getUpper != expectUpper {
		t.Errorf("Expect to get IntervalType: %d, but get %d instead.", expectUpper, getUpper)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(0)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(1)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = true
	getValid = intInterval.IsValid(50)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(100)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}

	expectValid = false
	getValid = intInterval.IsValid(101)
	if getValid != expectValid {
		t.Errorf("Expect to get valid: %v, but get %v instead.", expectValid, getValid)
		return
	}
}
