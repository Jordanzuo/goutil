package debugUtil

import (
	"fmt"
)

var (
	isDebug = false
)

// 设置DEBUG状态
// _isDebug：是否是DEBUG
func SetDebug(_isDebug bool) {
	isDebug = _isDebug
}

// 是否处于调试状态
func IsDebug() bool {
	return isDebug
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) {
	if isDebug {
		fmt.Print(a...)
	}
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) {
	if isDebug {
		fmt.Printf(format, a...)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	if isDebug {
		fmt.Println(a...)
	}
}
