package rangeUtil

import (
	"fmt"
	"regexp"
)

type RangeType uint8

const (
	DefaultRangeType   RangeType = iota
	InclusiveInclusive           // [a, b]
	InclusiveExclusive           // [a, b)
	ExclusiveInclusive           // (a, b]
	ExclusiveExclusive           // (a, b)
)

const (
	InclusiveOpenBracket  = "["
	InclusiveCloseBracket = "]"
	ExclusiveOpenBracket  = "("
	ExclusiveCloseBracket = ")"
	Delimiter             = ","
)

var (
	inclusiveInclusiveRegex *regexp.Regexp
	inclusiveExclusiveRegex *regexp.Regexp
	exclusiveInclusiveRegex *regexp.Regexp
	exclusiveExclusiveRegex *regexp.Regexp
)

func init() {
	inclusiveInclusiveRegex = regexp.MustCompile(fmt.Sprintf("^[%s][[:digit:]]+,[[:digit:]]+[%s]$", InclusiveOpenBracket, InclusiveCloseBracket))
	inclusiveExclusiveRegex = regexp.MustCompile(fmt.Sprintf("^[%s][[:digit:]]+,[[:digit:]]+[%s]$", InclusiveOpenBracket, ExclusiveCloseBracket))
	exclusiveInclusiveRegex = regexp.MustCompile(fmt.Sprintf("^[%s][[:digit:]]+,[[:digit:]]+[%s]$", ExclusiveOpenBracket, InclusiveCloseBracket))
	exclusiveExclusiveRegex = regexp.MustCompile(fmt.Sprintf("^[%s][[:digit:]]+,[[:digit:]]+[%s]$", ExclusiveOpenBracket, ExclusiveCloseBracket))
}

func parseRangeType(input string) RangeType {
	switch {
	case inclusiveInclusiveRegex.MatchString(input):
		return InclusiveInclusive
	case inclusiveExclusiveRegex.MatchString(input):
		return InclusiveExclusive
	case exclusiveInclusiveRegex.MatchString(input):
		return ExclusiveInclusive
	case exclusiveExclusiveRegex.MatchString(input):
		return ExclusiveExclusive
	}

	return DefaultRangeType
}
