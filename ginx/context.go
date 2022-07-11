package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func Handler(c *gin.Context, f func(c *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(&Context{
			Context: c,
		})
	}
}

func (c *Context) OutSuccess(data interface{}) {
	c.Out("success", "", data)
}

func (c *Context) OutFail(errmsg string) {
	c.Out("fail", errmsg, map[string]interface{}{})
}

func (c *Context) OutRelogin() {
	c.Out("login", "need login", map[string]interface{}{})
}

func (c *Context) Out(status string, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
