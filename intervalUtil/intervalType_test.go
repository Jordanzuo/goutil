package intervalUtil

import (
	"testing"
)

func TestParseRangeType(t *testing.T) {
	input := ""
	expect := DefaultIntervalType
	get := parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[1,100]"
	expect = InclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[1,100)"
	expect = InclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(1,100]"
	expect = ExclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(1,100)"
	expect = ExclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[-1,100]"
	expect = InclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[-1,100)"
	expect = InclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(-1,100]"
	expect = ExclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(-1,100)"
	expect = ExclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[-10,-1]"
	expect = InclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[-10,-1)"
	expect = InclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(-10,-1]"
	expect = ExclusiveInclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(-10,-1)"
	expect = ExclusiveExclusive
	get = parseIntervalType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}
}
