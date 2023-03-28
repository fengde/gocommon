package httpx

import (
	"fmt"
	"strings"
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

// PostForm 通用的http/post application/x-www-form-urlencoded 请求封装
func PostForm(url string, headers map[string]string, body map[string]interface{}, timeout ...time.Duration) (*Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	_body := map[string]string{}
	for k, v := range body {
		_body[k] = fmt.Sprintf("%v", v)
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"

	return do("post", url, headers, _body, timeout...)
}

// PostXML 通用的http/post text/xml 请求封装
func PostXML(url string, headers map[string]string, body string, timeout ...time.Duration) (*Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "text/xml; charset=utf-8"

	return do("post", url, headers, body, timeout...)
}

// PostJSON 通用的http/post application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PostJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json; charset=utf-8"

	return do("post", url, headers, body, timeout...)
}

// PutJSON 通用的Put方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PutJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json; charset=utf-8"

	return do("put", url, headers, body, timeout...)
}

// DeleteJSON 通用的Delete方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func DeleteJSON(url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json; charset=utf-8"

	return do("delete", url, headers, body, timeout...)
}

func do(method string, url string, headers map[string]string, body interface{}, timeout ...time.Duration) (*Response, error) {
	var r *resty.Request
	if len(timeout) > 0 {
		r = resty.New().SetTimeout(timeout[0]).R()
	} else {
		r = resty.New().R()
	}

	if len(headers) > 0 {
		r.SetHeaders(headers)
	}

	if body != nil {
		if ct, ok := headers["Content-Type"]; ok {
			switch {
			case strings.HasPrefix(ct, "application/x-www-form-urlencoded"):
				r.SetFormData(body.(map[string]string))
			case strings.HasPrefix(ct, "application/json"):
				r.SetBody(body)
			default:
				r.SetBody(body)
			}
		}
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
