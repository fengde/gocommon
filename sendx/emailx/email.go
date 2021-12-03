package emailx

import (
	"gopkg.in/gomail.v2"
)

type Client interface {
	SendText(subject string, content string, to []string, bcc []string, cc []string, file []string) error
	SendHTML(subject string, content string, to []string, bcc []string, cc []string, file []string) error
}

type EmailClient struct {
	dialer *gomail.Dialer
	user   string
	retry  int64
}

// NewEmailClient 新建邮件客户端
func NewEmailClient(host string, port int, user, password string, retry ...int64) Client {
	var retryCount int64
	if len(retry) > 0 {
		retryCount = retry[0]
	}
	ec := EmailClient{
		dialer: gomail.NewDialer(host, port, user, password),
		user:   user,
		retry:  retryCount,
	}
	return &ec
}

/* SendText 推送文本
参数出说明：
	subject 主题
	content 内容
	from 发送人
	to   接收人
	bcc  暗抄送
	cc   抄送
	files 本地附件路径
*/
func (p *EmailClient) SendText(subject string, content string, to []string, bcc []string, cc []string, file []string) error {
	return p.send(subject, "text/plain", content, to, bcc, cc, file)
}

/* SendHTML 推送html
参数出说明：
	subject 主题
	content 内容
	from 发送人
	to   接收人
	bcc  暗抄送
	cc   抄送
	files 附件路径
*/
func (p *EmailClient) SendHTML(subject string, content string, to []string, bcc []string, cc []string, file []string) error {
	return p.send(subject, "text/html", content, to, bcc, cc, file)
}

func (p *EmailClient) send(subject string, contentType string, content string, to []string, bcc []string, cc []string, file []string) error {
	var err error
	for i := 0; i < int(p.retry+1); i++ {
		m := gomail.NewMessage()
		m.SetHeader("From", p.user)
		m.SetHeader("To", to...)
		m.SetHeader("Bcc", bcc...)
		m.SetHeader("Cc", cc...)
		m.SetHeader("Subject", subject)
		m.SetBody(contentType, content)
		for _, f := range file {
			m.Attach(f)
		}
		if err = p.dialer.DialAndSend(m); err != nil {
			continue
		}

		break
	}

	return err
}
