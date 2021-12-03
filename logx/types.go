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
)

// Hook 拓展，可支持到ES, Kafka等写入
type Hook log.Hook