package outx

import (
	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/logx"
)

var (
	// 返回码定义
	StatusSuccess int64 = 0
	StatusArgErr  int64 = 1
	StatusSysErr  int64 = 2
	StatusAuthErr int64 = 3

	// 默认错误文案定义
	DefaultArgErr  = errorx.New(`参数错误`)
	DefaultSysErr  = errorx.New(`服务异常`)
	DefaultAuthErr = errorx.New(`鉴权错误，请重新登录`)
)

type Response struct {
	Status int64 `json:"status"`
	Msg string `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

// OutSuccess 成功返回
func OutSuccess(data map[string]interface{}) (*Response, error) {
	return out(StatusSuccess, nil, data)
}

// OutArgErr 参数错误返回
func OutArgErr(err ...error) (*Response, error) {
	var e = DefaultArgErr
	if len(err) > 0 {
		e = err[0]
	}
	return out(StatusArgErr, e, nil)
}

// OutSysErr 内部异常返回
func OutSysErr(err ...error) (*Response, error) {
	var e = DefaultSysErr
	if len(err) > 0 {
		e = err[0]
	}
	return out(StatusSysErr, e, nil)
}

// OutAuthErr 鉴权错误返回
func OutAuthErr(err ...error) (*Response, error) {
	var e = DefaultAuthErr
	if len(err) > 0 {
		e = err[0]
	}
	return out(StatusAuthErr, e, nil)
}

func out(code int64, err error, data map[string]interface{}) (*Response, error) {
	var message string

	if err != nil {
		logx.Errorf("%+v", err)
		message = err.Error()
	}

	if data == nil {
		data = map[string]interface{}{}
	}

	return &Response{
		Status: code,
		Msg:    message,
		Data:   data,
	}, nil
}
