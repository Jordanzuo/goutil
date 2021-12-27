package intervalUtil

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
	validRegex = regexp.MustCompile(fmt.Sprintf("^[%s%s][-]?[[:digit:]]+,[-]?[[:digit:]]+[%s%s]$", InclusiveOpenBracket, ExclusiveOpenBracket, InclusiveCloseBracket, ExclusiveCloseBracket)) // [a,b], [a,b), (a,b], (a,b)
}

func removeBracket(input string) string {
	input = strings.TrimPrefix(input, InclusiveOpenBracket)
	input = strings.TrimPrefix(input, ExclusiveOpenBracket)

	input = strings.TrimSuffix(input, InclusiveCloseBracket)
	input = strings.TrimSuffix(input, ExclusiveCloseBracket)

	return input
}

func IsValidFormat(input string) bool {
	if input == "" {
		return false
	}

	return validRegex.MatchString(input)
}
