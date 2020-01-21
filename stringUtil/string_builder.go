package stringUtil

import (
	"bytes"
)

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) AppendStrings(ss ...string) *StringBuilder {
	for i := range ss {
		builder.buffer.WriteString(ss[i])
	}
	return builder
}

func (builder *StringBuilder) Clear() *StringBuilder {
	var buffer bytes.Buffer
	builder.buffer = buffer
	return builder
}

func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}
