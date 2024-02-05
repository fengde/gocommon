package httpx

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type Response struct {
	*resty.Response
}

func newResponse(resp *resty.Response) *Response {
	return &Response{
		resp,
	}
}

type GetInput struct {
	Url     string
	Headers map[string]string
	Params  map[string]string
	Timeout time.Duration
}

// Get 通用的http/get请求封装
func Get(input *GetInput) (*Response, error) {
	var r *resty.Request
	if input.Timeout > 0 {
		r = resty.New().SetTimeout(input.Timeout).R()
	} else {
		r = resty.New().R()
	}

	if len(input.Headers) > 0 {
		r.SetHeaders(input.Headers)
	}
	if len(input.Params) > 0 {
		r.SetQueryParams(input.Params)
	}

	resp, err := r.EnableTrace().Get(input.Url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return newResponse(resp), nil
}

type PostFormInput struct {
	Url     string
	Headers map[string]string
	Body    map[string]interface{}
	Timeout time.Duration
}

// PostForm 通用的http/post application/x-www-form-urlencoded 请求封装
func PostForm(input *PostFormInput) (*Response, error) {
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}

	_body := map[string]string{}
	for k, v := range input.Body {
		_body[k] = fmt.Sprintf("%v", v)
	}

	input.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"

	if input.Timeout > 0 {
		return do("post", input.Url, input.Headers, _body, input.Timeout)
	}

	return do("post", input.Url, input.Headers, _body)
}

type PostXMLInput struct {
	Url     string
	Headers map[string]string
	Body    string
	Timeout time.Duration
}

// PostXML 通用的http/post text/xml 请求封装
func PostXML(input *PostXMLInput) (*Response, error) {
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}

	input.Headers["Content-Type"] = "text/xml; charset=utf-8"

	if input.Timeout > 0 {
		return do("post", input.Url, input.Headers, input.Body, input.Timeout)
	}

	return do("post", input.Url, input.Headers, input.Body)
}

type PostJSONInput struct {
	Url     string
	Headers map[string]string
	Body    interface{}
	Timeout time.Duration
}

// PostJSON 通用的http/post application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PostJSON(input *PostJSONInput) (*Response, error) {
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}

	input.Headers["Content-Type"] = "application/json; charset=utf-8"

	if input.Timeout > 0 {
		return do("post", input.Url, input.Headers, input.Body, input.Timeout)
	}

	return do("post", input.Url, input.Headers, input.Body)
}

type PutJSONInput struct {
	Url     string
	Headers map[string]string
	Body    interface{}
	Timeout time.Duration
}

// PutJSON 通用的Put方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func PutJSON(input *PutJSONInput) (*Response, error) {
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}

	input.Headers["Content-Type"] = "application/json; charset=utf-8"

	if input.Timeout > 0 {
		return do("put", input.Url, input.Headers, input.Body, input.Timeout)
	}

	return do("put", input.Url, input.Headers, input.Body)
}

type DeleteJSONInput struct {
	Url     string
	Headers map[string]string
	Body    interface{}
	Timeout time.Duration
}

// DeleteJSON 通用的Delete方法 application/json 请求封装;
// body参数支持传：string，[]byte，struct，map
func DeleteJSON(input *DeleteJSONInput) (*Response, error) {
	if input.Headers == nil {
		input.Headers = map[string]string{}
	}

	input.Headers["Content-Type"] = "application/json; charset=utf-8"
	if input.Timeout > 0 {
		return do("delete", input.Url, input.Headers, input.Body, input.Timeout)
	}

	return do("delete", input.Url, input.Headers, input.Body)
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
		return nil, errors.Errorf("暂未支持的method：%v", method)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return newResponse(resp), nil
}
