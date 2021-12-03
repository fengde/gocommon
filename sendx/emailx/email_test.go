package emailx

import (
	"testing"
)

func getEmailClient() Client {
	c := NewEmailClient("smtp.qq.com", 465, "940952619@qq.com", "*")
	return c
}

func TestEmailClient_SendHTML(t *testing.T) {
	c := getEmailClient()
	if err := c.SendHTML("hi", "<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"utf-8\">\n<title>菜鸟教程(runoob.com)</title>\n</head>\n<body>\n    <h1>我的第一个标题</h1>\n    <p>我的第一个段落。</p>\n</body>\n</html>",
		[]string{""}, []string{""}, nil, []string{"./email.go"}); err != nil {
		t.Error(err)
	}
}

func TestEmailClient_SendText(t *testing.T) {
	c := getEmailClient()
	if err := c.SendText("hi", "hello world", []string{""}, nil, nil, nil); err != nil {
		t.Error(err)
	}
}
