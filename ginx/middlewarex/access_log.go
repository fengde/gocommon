package middlewarex

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/fengde/gocommon/ginx"
	"github.com/fengde/gocommon/logx"
	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		ctx := logx.NewCtx(ginx.GetReqeustId(c))

		// start
		{
			body, _ := c.GetRawData()
			// 将原body塞回去
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			headers := []string{}

			for k, v := range c.Request.Header {
				headers = append(headers, k+":"+v[0])
			}

			// 请求前
			logx.InfofWithCtx(ctx, `=============收到请求=============
%v | %v/%v | %v
From: %v
Header:
%v
Body:
%v
Body size: %v bytes
`, c.Request.Method, c.Request.Host, c.Request.URL, c.Request.Proto,
				c.RemoteIP(),
				strings.Join(headers, "\n"),
				string(body),
				len(body))
		}

		c.Next()

		// end
		{
			headers := []string{}

			for k, v := range c.Writer.Header() {
				headers = append(headers, k+":"+v[0])
			}

			out := c.GetString("out")

			logx.InfofWithCtx(ctx, `=============请求结束=============
http status: %v
Header:
%v
Body:
%v
Body size: %v bytes
请求耗时: %v
`, c.Writer.Status(), strings.Join(headers, "\n"), out, len(out), time.Since(start))

		}
	}
}
