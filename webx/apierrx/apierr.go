package apierrx

import (
	"encoding/json"
	"fmt"
	"io"
)

var WrapC = NewAPIError

// NewAPIError handler 层调用，返回封装之后的error 给到middle层统一返回
func NewAPIError(code IAPICode, cause ...error) error {
	var c error
	if len(cause) > 0 {
		c = cause[0]
	}
	return &apiError{
		code:  code,
		cause: c,
		stack: callers(),
	}
}

type apiError struct {
	code  IAPICode
	cause error
	*stack
}

// Error implement interface error
func (a *apiError) Error() string {
	return fmt.Sprintf("[%d] - %s", a.code.Code(), a.code.Message())
}

func (a *apiError) Code() IAPICode {
	return a.code
}

func (a *apiError) Unwrap() error {
	return a.cause
}

func (a *apiError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			str := a.Error()
			if a.Unwrap() != nil {
				str += " " + a.Unwrap().Error()
			}
			_, _ = io.WriteString(s, str)
			a.stack.Format(s, verb)
			return
		}
		if s.Flag('#') {
			cause := ""
			if a.Unwrap() != nil {
				cause = a.Unwrap().Error()
			}
			data, _ := json.Marshal(errorMessage{
				Code:      a.code.Code(),
				Message:   a.code.Message(),
				Reference: a.code.Reference(),
				Cause:     cause,
				Stack:     fmt.Sprintf("%+v", a.stack),
			})
			_, _ = io.WriteString(s, string(data))
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, a.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", a.Error())
	}
}

func (a *apiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&errorMessage{
		Code:      a.code.Code(),
		Message:   a.code.Message(),
		Reference: a.code.Reference(),
	})
}

func (a *apiError) UnmarshalJSON(data []byte) error {
	e := &errorMessage{}
	if err := json.Unmarshal(data, e); err != nil {
		return err
	}
	a.code = NewAPICode(e.Code, e.Message, e.Reference)
	return nil
}

type errorMessage struct {
	Code      int    `json:"code"`
	Message   string `json:"message,omitempty"`
	Reference string `json:"reference,omitempty"`
	Cause     string `json:"cause,omitempty"`
	Stack     string `json:"stack,omitempty"`
}
