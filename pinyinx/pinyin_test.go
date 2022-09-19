package pinyinx

import (
	"testing"
)

func TestConver(t *testing.T) {
	t.Log(Convert("我爱你"))
	t.Log(Convert("我是中国人"))
}

func TestConvertWithTone(t *testing.T) {
	t.Log(ConvertWithTone("我爱你"))
	t.Log(ConvertWithTone("我是中国人"))
}

func TestConvertWithToneNumber(t *testing.T) {
	t.Log(ConvertWithToneNumber("我爱你"))
	t.Log(ConvertWithToneNumber("我是中国人"))
}
