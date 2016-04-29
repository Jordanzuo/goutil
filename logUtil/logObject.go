package logUtil

type logObject struct {
	// 日志信息
	logInfo string

	// 日志等级
	level LogType

	// 日志文件名称是否包含小时
	ifIncludeHour bool
}

func NewLogObject(_logInfo string, _level LogType, _ifIncludeHour bool) *logObject {
	return &logObject{
		logInfo:       _logInfo,
		level:         _level,
		ifIncludeHour: _ifIncludeHour,
	}
}
