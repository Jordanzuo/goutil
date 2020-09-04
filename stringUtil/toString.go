package stringUtil

import (
	"bytes"
	"fmt"
)

func StringListToString(list []string, delimiter string) string {
	var buffer bytes.Buffer
	for i, v := range list {
		if i != len(list)-1 {
			buffer.WriteString(v)
			buffer.WriteString(delimiter)
		} else {
			buffer.WriteString(v)
		}
	}

	return buffer.String()
}

func IntListToString(list []int, delimiter string) string {
	var buffer bytes.Buffer
	for i, v := range list {
		if i != len(list)-1 {
			buffer.WriteString(fmt.Sprintf("%d", v))
			buffer.WriteString(delimiter)
		} else {
			buffer.WriteString(fmt.Sprintf("%d", v))
		}
	}

	return buffer.String()
}

func Int32ListToString(list []int32, delimiter string) string {
	var buffer bytes.Buffer
	for i, v := range list {
		if i != len(list)-1 {
			buffer.WriteString(fmt.Sprintf("%d", v))
			buffer.WriteString(delimiter)
		} else {
			buffer.WriteString(fmt.Sprintf("%d", v))
		}
	}

	return buffer.String()
}

func Int64ListToString(list []int64, delimiter string) string {
	var buffer bytes.Buffer
	for i, v := range list {
		if i != len(list)-1 {
			buffer.WriteString(fmt.Sprintf("%d", v))
			buffer.WriteString(delimiter)
		} else {
			buffer.WriteString(fmt.Sprintf("%d", v))
		}
	}

	return buffer.String()
}
