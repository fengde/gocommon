package captchax

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
)

// 生成图像验证码
func NewCaptchaImage() (captchatId string, link string) {
	captchatId = captcha.New()

	return captchatId, fmt.Sprintf("/captcha/%s.png", captchatId)
}

// 生成音频验证码
func NewCaptchaAudio() (captchatId string, link string) {
	captchatId = captcha.New()

	return captchatId, fmt.Sprintf("/captcha/download/%s.wav", captchatId)
}

// 验证
func Verify(captchatId string, userSubmit string) bool {
	return captcha.VerifyString(captchatId, userSubmit)
}

// 验证码资源响应
func LinkHandle() http.Handler {
	return captcha.Server(captcha.StdWidth, captcha.StdHeight)
}
