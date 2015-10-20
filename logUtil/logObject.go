package logUtil

type LogObject struct {
	// 日志信息
	LogInfo string

	// 日志等级
	Level LogType

	// 日志文件名称是否包含小时
	IfIncludeHour bool
}

func NewLogObject(logInfo string, level LogType, ifIncludeHour bool) *LogObject {
	return &LogObject{
		LogInfo:       logInfo,
		Level:         level,
		IfIncludeHour: ifIncludeHour,
	}
}
