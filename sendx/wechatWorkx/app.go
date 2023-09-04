package wechatWorkx

import (
	"time"

	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/httpx"
	"github.com/fengde/gocommon/jsonx"
	"github.com/fengde/gocommon/logx"
)

// APPClient 企业微信应用客户端，用于应用给用户，给群发消息
type APPClient struct {
	CorpID     string
	CorpSecret string
}

// NewAPPClient 传入企业ID, 企业秘钥 返回 企业微信应用推送客户端
func NewAPPClient(corpID, corpSecret string) *APPClient {
	wwc := APPClient{
		CorpID:     corpID,
		CorpSecret: corpSecret,
	}
	return &wwc
}

type tokenSer struct {
	Token     string
	ExpiredAt time.Time
}

var token tokenSer

// getAccessToken 获取access_token
func (p *APPClient) getAccessToken() (string, error) {
	if !token.ExpiredAt.IsZero() && time.Now().Before(token.ExpiredAt) {
		return token.Token, nil
	}
	resp, err := httpx.Get(&httpx.GetInput{
		Url: `https://qyapi.weixin.qq.com/cgi-bin/gettoken`,
		Params: map[string]string{
			"corpid":     p.CorpID,
			"corpsecret": p.CorpSecret,
		},
	})
	if err != nil {
		return "", err
	}

	result := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}{}
	if err := jsonx.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	if result.AccessToken != "" {
		token.Token = result.AccessToken
		token.ExpiredAt = time.Now().Add(time.Second * time.Duration(result.ExpiresIn-100))
		return result.AccessToken, nil
	}
	return "", errorx.New("无法获取到企业微信access_token")
}

type SendResult struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// Send 发送消息
// data 参照 https://open.work.weixin.qq.com/api/doc/90000/90135/90236
func (p *APPClient) Send(data string) error {
	if len(data) < 1 {
		return errorx.New("参数错误")
	}

	accessToken, err := p.getAccessToken()
	if err != nil {
		return err
	}

	logx.Info(accessToken)

	resp, err := httpx.PostJSON(&httpx.PostJSONInput{
		Url:  `https://qyapi.weixin.qq.com/cgi-bin/message/send?debug=1&access_token=` + accessToken,
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

// NewChatGroup 新建群聊组，返回群聊ID
func (p *APPClient) NewChatGroup(title string, owner string, users []string) (string, error) {
	if len(title) < 1 || len(owner) < 1 || len(users) < 1 {
		return "", errorx.New("参数错误")
	}

	accessToken, err := p.getAccessToken()
	if err != nil {
		return "", err
	}

	resp, err := httpx.PostJSON(&httpx.PostJSONInput{
		Url: `https://qyapi.weixin.qq.com/cgi-bin/appchat/create?debug=1&access_token=` + accessToken,
		Body: map[string]interface{}{
			"name":     title,
			"owner":    owner,
			"userlist": users,
		},
	})
	if err != nil {
		return "", err
	}

	result := struct {
		*SendResult
		ChatID string `json:"chatid"`
	}{}
	if err := jsonx.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	if result.Errcode != 0 {
		return "", errorx.Errorf("errmsg: %s", result.Errmsg)
	}

	return result.ChatID, nil
}

// SendToGroup 发送消息到群
// data 参照 https://open.work.weixin.qq.com/api/doc/90000/90135/90248
func (p *APPClient) SendToGroup(data string) error {
	if len(data) < 1 {
		return errorx.New("参数错误")
	}

	accessToken, err := p.getAccessToken()
	if err != nil {
		return err
	}

	resp, err := httpx.PostJSON(&httpx.PostJSONInput{
		Url:  `https://qyapi.weixin.qq.com/cgi-bin/appchat/send?debug=1&access_token=` + accessToken,
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
