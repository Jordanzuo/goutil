package logUtil

type logType int

// 日志等级
const (
	// 信息
	Info logType = iota

	// 警告
	Warn

	// 调试
	Debug

	// 错误
	Error

	// 致命
	Fatal
)

var levels = [...]string{
	"Info",
	"Warn",
	"Debug",
	"Error",
	"Fatal",
}

func (t logType) String() string {
	return levels[t]
}
