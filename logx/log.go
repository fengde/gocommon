package logx

import (
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
}

// SetLogFile 设置输出到文件
func SetLogFile(logpath string, logSaveDays int) {
	os.MkdirAll(filepath.Dir(logpath), os.ModePerm)
	writer, _ := rotatelogs.New(
		logpath+".%Y-%m-%d",
		rotatelogs.WithLinkName(logpath),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(uint(logSaveDays)),
	)

	logger.SetOutput(io.MultiWriter(writer, os.Stdout))
}

// SetFormatter 设置日志输出格式
func SetFormatter(formatter Formatter) {
	logger.SetFormatter(formatter)
}

// SetLevel 设置日志打印级别
func SetLevel(level Level) {
	logger.SetLevel(log.Level(level))
}

// AddHook 添加hook
func AddHook(hook Hook) {
	logger.AddHook(hook)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func ErrorStack(err error) {
	logger.Errorf("%+v", err)
}
