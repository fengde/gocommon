package emailx

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"

	"github.com/go-gomail/gomail"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Mailer struct {
	mailer *gomail.Dialer
}

func NewMailer(host string, port int, user, password string) *Mailer {
	mailer := gomail.NewDialer(host, port, user, password)
	mailer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	return &Mailer{
		mailer: mailer,
	}
}

// 发送邮件
func (p *Mailer) SendEmail(from string, to []string, topic, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", topic)
	msg.SetBody("text/html; charset=UTF-8", content)
	return p.mailer.DialAndSend(msg)
}

// 检查smtp服务器是否可通
func (p *Mailer) CheckConf() error {
	closer, err := p.mailer.Dial()
	if err == nil {
		closer.Close()
	}
	return err
}

// 发送带附件的邮件
func (p *Mailer) SendAttach(from string, to []string, topic, content string, fileName string, bytes []byte) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", topic)
	msg.SetBody("text/html; charset=UTF-8", content)
	msg.Attach(Utf8ToGbk(fileName), gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(bytes)
		return err
	}))
	return p.mailer.DialAndSend(msg)
}

func Utf8ToGbk(str string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return str
	}
	return string(d)
}
