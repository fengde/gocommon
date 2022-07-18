package ginx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/fengde/gocommon/jsonx"
	"github.com/fengde/gocommon/logx"
	"github.com/fengde/gocommon/safex"
	"github.com/fengde/gocommon/timex"
	"github.com/fengde/gocommon/toolx"
	"github.com/gin-gonic/gin"
)

const RequestIdName = "request_id"

type Context struct {
	*gin.Context
}

func GetReqeustId(ginc *gin.Context) string {
	requestId := ginc.GetString(RequestIdName)
	if requestId == "" {
		requestId = fmt.Sprintf("%v%s", timex.NowUnixNano(), toolx.NewNumberCode(4))
		ginc.Set(RequestIdName, requestId)
	}
	return requestId
}

func Handler(f func(c *Context)) gin.HandlerFunc {
	return func(ginc *gin.Context) {
		defer safex.Recover(func() {
			ginc.JSON(http.StatusOK, gin.H{
				"status":      "fail",
				"message":     "内部异常",
				"data":        map[string]interface{}{},
				RequestIdName: GetReqeustId(ginc),
			})
		})

		f(&Context{
			Context: ginc,
		})
	}
}

// OutSuccess 成功返回
func (c *Context) OutSuccess(data interface{}) {
	c.Out("success", "", data)
}

// OutFail 失败返回
func (c *Context) OutFail(errmsg string) {
	c.Out("fail", errmsg, map[string]interface{}{})
}

// OutRelogin 提示重新登录返回
func (c *Context) OutRelogin() {
	c.Out("login", "need login", map[string]interface{}{})
}

// Out 通用返回
func (c *Context) Out(status string, message string, data interface{}) {
	ginH := gin.H{
		"status":      status,
		"message":     message,
		"data":        data,
		RequestIdName: c.RequestId(),
	}
	c.Set("out", jsonx.MarshalToStringNoErr(ginH))
	c.JSON(http.StatusOK, ginH)
}

// GetJsonData 解析json数据，按照govalidator做数据校验
func (c *Context) GetJsonData(r interface{}) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}
	if _, err := govalidator.ValidateStruct(r); err != nil {
		return err
	}
	return nil
}

// RequestId 返回http请求id
func (c *Context) RequestId() string {
	requestId, _ := c.Get(RequestIdName)
	return fmt.Sprintf("%v", requestId)
}

// LogCtx 返回日志ctx
func (c *Context) LogCtx() context.Context {
	return logx.NewCtx(c.RequestId())
}
