package googlex

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	text, err := Translate("我爱你", "zh-CN", "zh-TW")
	t.Log(text, err)

	text, err = Translate("我爱你", "zh-CN", "en")
	t.Log(text, err)
}
