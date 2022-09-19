package pinyinx

import (
	"github.com/mozillazg/go-pinyin"
)

// 汉字转拼音（中国人-> zhong guo ren)
func Convert(chinese string) []string {
	return pinyin.LazyConvert(chinese, nil)
}

// 汉字转拼音（带声调，中国人->zhōng guó rén）
func ConvertWithTone(chinese string) []string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone

	var back []string
	for _, item := range pinyin.Pinyin(chinese, a) {
		if len(item) > 0 {
			back = append(back, item[0])
		}
	}
	return back
}

// 汉字转拼音（带声调，中国人->zho1ng guo2 re2n）
func ConvertWithToneNumber(chinese string) []string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone2
	var back []string
	for _, item := range pinyin.Pinyin(chinese, a) {
		if len(item) > 0 {
			back = append(back, item[0])
		}
	}
	return back
}
