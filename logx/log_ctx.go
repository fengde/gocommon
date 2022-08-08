package logx

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/fengde/gocommon/toolx"
)

var LOGID = "logid"

// NewCtx 生成带LogID的ctx
func NewCtx(logID ...string) context.Context {
	if len(logID) > 0 && logID[0] != "" {
		return context.WithValue(context.Background(), LOGID, logID[0])
	}
	return context.WithValue(context.Background(), LOGID, fmt.Sprintf("%v%v", time.Now().UnixNano(), toolx.NewNumberCode(5)))
}

// GetLogID 获取logid
func GetLogID(ctx context.Context) string {
	v := ctx.Value(LOGID)
	if v != nil {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func getCtxLogger(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return logger.WithField(LOGID, "ctx is nil")
	}
	return logger.WithField(LOGID, GetLogID(ctx))
}

func DebugWithCtx(ctx context.Context, args ...interface{}) {
	getCtxLogger(ctx).Debug(format(args...))
}

func DebugfWithCtx(ctx context.Context, format string, args ...interface{}) {
	getCtxLogger(ctx).Debugf(format, args...)
}

func InfoWithCtx(ctx context.Context, args ...interface{}) {
	getCtxLogger(ctx).Info(format(args...))
}

func InfofWithCtx(ctx context.Context, format string, args ...interface{}) {
	getCtxLogger(ctx).Infof(format, args...)
}

func WarnWithCtx(ctx context.Context, args ...interface{}) {
	getCtxLogger(ctx).Warn(format(args...))
}

func WarnfWithCtx(ctx context.Context, format string, args ...interface{}) {
	getCtxLogger(ctx).Warnf(format, args...)
}

func ErrorWithCtx(ctx context.Context, args ...interface{}) {
	getCtxLogger(ctx).Error(format(args...))
}

func ErrorfWithCtx(ctx context.Context, format string, args ...interface{}) {
	getCtxLogger(ctx).Errorf(format, args...)
}

func ErrorStackWithCtx(ctx context.Context, err error) {
	getCtxLogger(ctx).Errorf("%+v", err)
}
