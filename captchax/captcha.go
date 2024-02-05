package captchax

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
)

// NewCaptchaImage 生成图像验证码，二维码过期时间 10min
func NewCaptchaImage(length ...int) (captchatId string, link string) {
	if len(length) > 0 {
		captchatId = captcha.NewLen(length[0])
	} else {
		captchatId = captcha.New()
	}

	return captchatId, fmt.Sprintf("/captcha/%s.png", captchatId)
}

// NewCaptchaAudio 生成音频验证码，二维码过期时间 10min
func NewCaptchaAudio(length ...int) (captchatId string, link string) {
	if len(length) > 0 {
		captchatId = captcha.NewLen(length[0])
	} else {
		captchatId = captcha.New()
	}

	return captchatId, fmt.Sprintf("/captcha/download/%s.wav", captchatId)
}

// Verify 验证
func Verify(captchatId string, userSubmit string) bool {
	return captcha.VerifyString(captchatId, userSubmit)
}

// LinkHandle 验证码资源响应
func LinkHandle() http.Handler {
	return captcha.Server(captcha.StdWidth, captcha.StdHeight)
}
