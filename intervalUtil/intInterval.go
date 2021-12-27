package intervalUtil

import (
	"errors"
	"strconv"
	"strings"
)

// We use int64 as the underlying data type
type IntInterval struct {
	Lower        int64
	Upper        int64
	IntervalType IntervalType
}

func (this IntInterval) IsValid(value int64) bool {
	switch this.IntervalType {
	case InclusiveInclusive:
		return this.Lower <= value && value <= this.Upper
	case InclusiveExclusive:
		return this.Lower <= value && value < this.Upper
	case ExclusiveInclusive:
		return this.Lower < value && value <= this.Upper
	case ExclusiveExclusive:
		return this.Lower < value && value < this.Upper
	}

	return false
}

func NewIntInterval(input string) (value IntInterval, err error) {
	if !IsValidFormat(input) {
		err = errors.New("Invalid interval expression format")
		return
	}

	intervalType := parseIntervalType(input)
	input = removeBracket(input)
	itemList := strings.Split(input, Delimiter)
	if len(itemList) != 2 {
		err = errors.New("Invalid interval expression format")
		return
	}

	lower, err := strconv.ParseInt(itemList[0], 10, 64)
	if err != nil {
		err = errors.New("Parse interval expression to int64 failed.")
		return
	}

	upper, err := strconv.ParseInt(itemList[1], 10, 64)
	if err != nil {
		err = errors.New("Parse interval expression to int64 failed.")
		return
	}

	value = IntInterval{
		Lower:        lower,
		Upper:        upper,
		IntervalType: intervalType,
	}
	return
}
