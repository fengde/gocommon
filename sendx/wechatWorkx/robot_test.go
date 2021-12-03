package wechatWorkx

import "testing"

func TestRobotClient_Send(t *testing.T) {
	client := NewRobotClient()
	err := client.Send("key", `{"msgtype":"text", "text": {"content":"机器人自动消息群发测试"} }`)
	if err != nil {
		t.Errorf("%+v", err)
	}
}
