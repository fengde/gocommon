package logx

import log "github.com/sirupsen/logrus"

// Formatter 输出的格式
type Formatter log.Formatter

var (
	TextFormatter = &log.TextFormatter{}
	JSONFormatter = &log.JSONFormatter{}
)

// Level 日志级别
type Level log.Level

// 对外只开放【debug, info, warn, error】几种日志级别
var (
	DebugLevel = Level(log.DebugLevel)
	InfoLevel = Level(log.InfoLevel)
	WarnLevel = Level(log.WarnLevel)
	ErrorLevel = Level(log.ErrorLevel)
	PanicLevel = Level(log.PanicLevel)
	FatalLevel = Level(log.FatalLevel)
	TraceLevel = Level(log.TraceLevel)
)

func Level2LogrusLevel(level Level) log.Level {
	switch level {
	case DebugLevel:
		return log.DebugLevel
	case InfoLevel:
		return log.InfoLevel
	case WarnLevel:
		return log.WarnLevel
	case ErrorLevel:
		return log.ErrorLevel
	case PanicLevel:
		return log.PanicLevel
	case FatalLevel:
		return log.FatalLevel
	case TraceLevel:
		return log.TraceLevel
	default:
		panic("unknow level")
	}
}

// Hook 拓展，可支持到ES, Kafka等写入
type Hook log.Hook