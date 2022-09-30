package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

// New 新建error
func New(message string) error {
	return errors.New(message)
}

// Errorf 新建error，支持传参字符串
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// WithStack 给err附带上stack信息，适用于报错”路径分析“
func WithStack(err error) error {
	return errors.WithStack(err)
}

// WithMessage 给err附带上message, 一般用于中间服务层逻辑
func WithMessage(err error, message string) error {
	return errors.WithMessage(err, message)
}

// WithMessagef 给err附带上message, 一般用于中间服务层逻辑
func WithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args...)
}

// Wrap 给err附带上stack信息和message信息，一般用于底层结构
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf 给err附带上stack信息和message信息，一般用于底层结构
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// GetStack 返回错误堆栈信息，方便打印
func GetStack(err error) string {
	return fmt.Sprintf("%+v", err)
}
