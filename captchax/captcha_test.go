package captchax

import (
	"testing"

	"github.com/fengde/gocommon/logx"
)

func TestNewCaptchaImage(t *testing.T) {
	id, link := NewCaptchaImage()
	logx.Info(id, link)
}

func TestNewCaptchaAudio(t *testing.T) {
	id, link := NewCaptchaAudio()
	logx.Info(id, link)
}
