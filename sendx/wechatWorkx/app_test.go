package wechatWorkx

import "testing"

func getClient() *APPClient {
	return NewAPPClient("corpID", "corpSecrect")
}

func TestAPPClient_SendText(t *testing.T) {
	err := getClient().Send("{\"touser\": \"fengde\", \"agentid\": *, \"msgtype\":\"text\", \"text\":{\"content\": \"hi... test\"}}")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestAPPClient_NewChatGroup(t *testing.T) {
	chatID, err := getClient().NewChatGroup("fengde test", "fengde", []string{"fengde", ""})
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Log(chatID)
}

func TestAPPClient_SendTextToGroup(t *testing.T) {
	err := getClient().SendToGroup("{}")
	if err != nil {
		t.Errorf("%+v", err)
	}
}
