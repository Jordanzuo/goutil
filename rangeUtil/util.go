package rangeUtil

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	emptyRegex *regexp.Regexp
	validRegex *regexp.Regexp
)

func init() {
	emptyRegex = regexp.MustCompile(fmt.Sprintf("^[%s%s][%s%s]$", InclusiveOpenBracket, ExclusiveOpenBracket, InclusiveCloseBracket, ExclusiveCloseBracket))                          // [], [), (], ()
	validRegex = regexp.MustCompile(fmt.Sprintf("^[%s%s][[:digit:]]+,[[:digit:]]+[%s%s]$", InclusiveOpenBracket, ExclusiveOpenBracket, InclusiveCloseBracket, ExclusiveCloseBracket)) // [a,b], [a,b), (a,b], (a,b)
}

func removeBracket(input string) string {
	input = strings.TrimPrefix(input, InclusiveOpenBracket)
	input = strings.TrimPrefix(input, ExclusiveOpenBracket)

	input = strings.TrimSuffix(input, InclusiveCloseBracket)
	input = strings.TrimSuffix(input, ExclusiveCloseBracket)

	return input
}

func IsEmptyRange(input string) bool {
	if input == "" {
		return true
	}

	return emptyRegex.MatchString(input)
}

func IsValidFormat(input string) bool {
	if input == "" {
		return false
	}

	return validRegex.MatchString(input)
}
