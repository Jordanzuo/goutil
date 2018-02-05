package logUtil

import (
	"fmt"
)

// 信息日志记录
// format:日志格式
// args:参数列表
func InfoLog(format string, args ...interface{}) {
	if len(args) <= 0 {
		NormalLog(format, Info)
	} else {
		NormalLog(fmt.Sprintf(format, args...), Info)
	}
}

// 错误日志记录
// format:日志格式
// args:参数列表
func WarnLog(format string, args ...interface{}) {
	if len(args) <= 0 {
		NormalLog(format, Warn)
	} else {
		NormalLog(fmt.Sprintf(format, args...), Warn)
	}
}

// 调试日志记录
// format:日志格式
// args:参数列表
func DebugLog(format string, args ...interface{}) {
	if len(args) <= 0 {
		NormalLog(format, Debug)
	} else {
		NormalLog(fmt.Sprintf(format, args...), Debug)
	}
}

// 警告日志记录
// format:日志格式
// args:参数列表
func ErrorLog(format string, args ...interface{}) {
	if len(args) <= 0 {
		NormalLog(format, Error)
	} else {
		NormalLog(fmt.Sprintf(format, args...), Error)
	}
}

// 致命错误日志记录
// format:日志格式
// args:参数列表
func FatalLog(format string, args ...interface{}) {
	if len(args) <= 0 {
		NormalLog(format, Fatal)
	} else {
		NormalLog(fmt.Sprintf(format, args...), Fatal)
	}
}
