package apierrx

import "errors"

// 通用Code, 实际工程下，定义自己服务的Code
// code编码规范：
//
//	    高1位   固定标识
//		高2-4位 服务编码
//	    高5-7位 逻辑编码
var (
	CommonAPICodeErrArg     = NewAPICode(1000001, "参数错误")
	CommonAPICodeErrData    = NewAPICode(1000002, "数据错误")
	CommonAPICodeErrAuth    = NewAPICode(1000003, "认证失败")
	CommonAPICodeErrIll     = NewAPICode(1000004, "违规操作")
	CommonAPICodeErrService = NewAPICode(1000005, "服务异常")
	CommonAPICodeErrCfg     = NewAPICode(1000006, "配置异常")
)

type IAPICode interface {
	Code() int
	Message() string
	Reference() string
	ServiceCode() int
	BusinessCode() int
}

func NewAPICode(code int, message string, reference ...string) IAPICode {
	ref := ""
	if len(reference) > 0 {
		ref = reference[0]
	}

	return &apiCode{
		code: code,
		msg:  message,
		ref:  ref,
	}
}

type apiCode struct {
	code int
	msg  string
	ref  string
}

func (a *apiCode) Code() int {
	return a.code
}

func (a *apiCode) Message() string {
	return a.msg
}

func (a *apiCode) Reference() string {
	return a.ref
}

// ServiceCode 服务编码
func (a *apiCode) ServiceCode() int {
	return a.Code() % 1000000 / 1000
}

// BusinessCode 逻辑编码
func (a *apiCode) BusinessCode() int {
	return a.Code() % 1000
}

// ParseAPICode 解析err 对应的APICode
func ParseAPICode(err error) IAPICode {
	for {
		if e, ok := err.(interface {
			Code() IAPICode
		}); ok {
			return e.Code()
		}
		if errors.Unwrap(err) == nil {
			return CommonAPICodeErrService
		}
		err = errors.Unwrap(err)
	}
}

// IsAPICode 判断err是不是某个APICode
func IsAPICode(err error, code IAPICode) bool {
	if err == nil {
		return false
	}

	for {
		if e, ok := err.(interface {
			Code() IAPICode
		}); ok {
			if e.Code().Code() == code.Code() {
				return true
			}
		}

		if errors.Unwrap(err) == nil {
			return false
		}

		err = errors.Unwrap(err)
	}
}
