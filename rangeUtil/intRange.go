package rangeUtil

import (
	"errors"
	"strconv"
	"strings"
)

// We use int64 as the underlying data type
type IntRange struct {
	Lower     int64
	Upper     int64
	RangeType RangeType
}

func (this IntRange) IsValid(value int64) bool {
	switch this.RangeType {
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

func NewIntRange(input string) (value IntRange, err error) {
	if IsEmptyRange(input) {
		err = errors.New("Empty range expression")
		return
	}

	if !IsValidFormat(input) {
		err = errors.New("Invalid range expression format")
		return
	}

	rangeType := parseRangeType(input)

	input = removeBracket(input)
	itemList := strings.Split(input, Delimiter)
	if len(itemList) != 2 {
		err = errors.New("Invalid range expression format")
		return
	}

	lower, err := strconv.ParseInt(itemList[0], 10, 64)
	if err != nil {
		err = errors.New("Parse range expression to int64 failed.")
		return
	}

	upper, err := strconv.ParseInt(itemList[1], 10, 64)
	if err != nil {
		err = errors.New("Parse range expression to int64 failed.")
		return
	}

	value = IntRange{
		Lower:     lower,
		Upper:     upper,
		RangeType: rangeType,
	}
	return
}
