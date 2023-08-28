package googlex

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fengde/gocommon/httpx"
)

// 翻译text，指定text原语言，以及需要翻译的语言
func Translate(text string, fromLang string, toLang string) (string, error) {
	resp, err := httpx.Get(fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", fromLang, toLang, url.QueryEscape(text)), nil, nil)
	if err != nil {
		return "", nil
	}

	s := resp.String()
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, "null,", "")
	s = strings.Trim(s, `"`)
	ps := strings.Split(s, `","`)
	return ps[0], nil
}

// JianTiZhongWen2En 简体中文翻译成英文
func JianTiZhongWen2En(text string) (string, error) {
	return Translate(text, "zh-CN", "en")
}

// FanTiZhongWen2En 繁体中文翻译成英文
func FanTiZhongWen2En(text string) (string, error) {
	return Translate(text, "zh-TW", "en")
}

// JianTi2FanTi 简体中文转繁体
func JianTi2FanTi(text string) (string, error) {
	return Translate(text, "zh-CN", "zh-TW")
}

// FanTi2JianTi 繁体转简体中文
func FanTi2JianTi(text string) (string, error) {
	return Translate(text, "zh-TW", "zh-CN")
}
