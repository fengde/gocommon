package httpx

import (
	"time"

	"github.com/fengde/gocommon/errorx"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	*resty.Response
}

func newResponse(resp *resty.Response) *Response {
	return &Response{
		resp,
	}
}

// Get 通用的http/get请求封装
func Get(url string, headers map[string]string, params map[string]string, timeout ...time.Duration) (*Response, error) {
	var r *resty.Request
	if len(timeout) > 0 {
		r = resty.New().SetTimeout(timeout[0]).R()
	} else {
		r = resty.New().R()
	}

	if len(headers) > 0 {
		r.SetHeaders(headers)
	}
	if len(params) > 0 {
		r.SetQueryParams(params)
	}

	resp, err := r.EnableTrace().Get(url)
	if err != nil {
		return nil, errorx.WithStack(err)
	}

	return newResponse(resp), nil
}

// PostJSON 通用的http/post application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PostJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	return json("post", url, headers, body, timeout...)
}

// PutJSON 通用的Put方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PutJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	return json("put", url, headers, body, timeout...)
}

// DeleteJSON 通用的Delete方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func DeleteJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	return json("delete", url, headers, body, timeout...)
}

func json(method string, url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	var r *resty.Request
	if len(timeout) > 0 {
		r = resty.New().SetTimeout(timeout[0]).R().SetHeader("Content-Type", "application/json")
	} else {
		r = resty.New().R().SetHeader("Content-Type", "application/json")
	}

	if len(headers) > 0 {
		r.SetHeaders(headers)
	}
	if body != nil {
		r.SetBody(body)
	}

	var resp *resty.Response
	var err error

	switch method {
	case "post":
		resp, err = r.EnableTrace().Post(url)
	case "put":
		resp, err = r.EnableTrace().Put(url)
	case "delete":
		resp, err = r.EnableTrace().Delete(url)
	default:
		return nil, errorx.Errorf("暂未支持的method：%v", method)
	}

	if err != nil {
		return nil, errorx.WithStack(err)
	}

	return newResponse(resp), nil
}
