package debugUtil

import (
	"fmt"
)

var (
	mIsDebug = false
)

func SetDebug(isDebug bool) {
	mIsDebug = isDebug
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) {
	if mIsDebug {
		fmt.Print(a...)
	}
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) {
	if mIsDebug {
		fmt.Printf(format, a...)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	if mIsDebug {
		fmt.Println(a...)
	}
}
