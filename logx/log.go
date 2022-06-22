package logx

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fengde/gocommon/errorx"

	"github.com/evalphobia/logrus_sentry"
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
// 参数：
//	logpath 日志文件路径
//  logSaveDays 日志保留天数
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

// AddSentryHook 添加sentry hook。指定levels类型日志推送到sentry
func AddSentryHook(dsn string, levels []Level) error {
	var tmp []log.Level
	for _, level := range levels {
		tmp = append(tmp, Level2LogrusLevel(level))
	}
	hook, err := logrus_sentry.NewSentryHook(dsn, tmp)
	if err != nil {
		return errorx.WithStack(err)
	}

	AddHook(hook)

	return nil
}

func Debug(args ...interface{}) {
	logger.Debug(format(args...))
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(format(args...))
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(format(args...))
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(format(args...))
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func ErrorStack(err error) {
	logger.Errorf("%+v", err)
}

func format(args ...interface{}) string {
	var replace []string

	for i := 0; i < len(args); i++ {
		replace = append(replace, "%v")
	}

	return fmt.Sprintf(strings.Join(replace, " "), args...)
}
