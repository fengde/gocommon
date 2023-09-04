package wechatWorkx

import (
	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/httpx"
	"github.com/fengde/gocommon/jsonx"
	"github.com/fengde/gocommon/structx/setx"
	"github.com/tidwall/gjson"
)

type RobotClient struct{}

// NewRobotClient 新建群机器人客户端
func NewRobotClient() *RobotClient {
	rc := RobotClient{}
	return &rc
}

var msgtypes = setx.Set{}

func init() {
	msgtypes.Store("text", "markdown", "image", "news", "file", "template_card")
}

// Send 发送内容到群
// data参数参照 https://work.weixin.qq.com/api/doc/90000/90136/91770
func (p *RobotClient) Send(key, data string) error {
	msgtype := gjson.Get(data, "msgtype").String()
	if !msgtypes.Has(msgtype) {
		return errorx.Errorf(`msgtype格式不支持: %s`, msgtype)
	}
	if !gjson.Get(data, msgtype).Exists() {
		return errorx.Errorf(`内容不存在: %s`, msgtype)
	}

	resp, err := httpx.PostJSON(&httpx.PostJSONInput{
		Url:  `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=` + key,
		Body: data,
	})
	if err != nil {
		return err
	}

	result := SendResult{}
	if err := jsonx.Unmarshal(resp.Body(), &result); err != nil {
		return err
	}

	if result.Errcode != 0 {
		return errorx.Errorf("errmsg: %s", result.Errmsg)
	}

	return nil
}
