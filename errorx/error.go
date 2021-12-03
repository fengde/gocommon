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

// WithStack 给err附带上Stack信息
func WithStack(err error) error {
	return errors.WithStack(err)
}

// Wrap 给err附带上Stack信息和message信息
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf 给err附带上Stack信息和message信息
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// Stack 返回错误堆栈信息
func GetStack(err error) string {
	return fmt.Sprintf("%+v", err)
}
