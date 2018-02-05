package debugUtil

import (
	"fmt"
)

var (
	isDebug                         = false
	code            Code            = Code_Hightlight
	foregroundColor ForegroundColor = Foreground_Purple
	backgroundColor BackgroundColor = BackgroundColor_Black
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

// 设置显示信息
func SetDisplayInfo(_code Code, _foregroundColor ForegroundColor, _backgroundColor BackgroundColor) {
	code = _code
	foregroundColor = _foregroundColor
	backgroundColor = _backgroundColor
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) {
	if !isDebug {
		return
	}

	for _, v := range a {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, code, backgroundColor, foregroundColor, v, 0x1B)
	}
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) {
	if !isDebug {
		return
	}

	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, code, backgroundColor, foregroundColor, fmt.Sprintf(format, a...), 0x1B)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	if !isDebug {
		return
	}

	for _, v := range a {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, code, backgroundColor, foregroundColor, v, 0x1B)
	}
	fmt.Println()
}
