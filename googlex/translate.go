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
