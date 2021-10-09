package rangeUtil

import (
	"testing"
)

func TestParseRangeType(t *testing.T) {
	input := ""
	expect := DefaultRangeType
	get := parseRangeType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[1,100]"
	expect = InclusiveInclusive
	get = parseRangeType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "[1,100)"
	expect = InclusiveExclusive
	get = parseRangeType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(1,100]"
	expect = ExclusiveInclusive
	get = parseRangeType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}

	input = "(1,100)"
	expect = ExclusiveExclusive
	get = parseRangeType(input)
	if get != expect {
		t.Errorf("Expect to get %d, but get %d instead.", expect, get)
		return
	}
}
