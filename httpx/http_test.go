package httpx

import (
	"testing"
)

func TestPostXML(t *testing.T) {
	r, err := PostXML(&PostXMLInput{
		Url:  "http://httpbin.org/anything",
		Body: "<abc></abc>",
	})
	t.Log(r, err)
}

func TestPostForm(t *testing.T) {
	r, err := PostForm(&PostFormInput{
		Url: "http://httpbin.org/anything",
		Body: map[string]interface{}{
			"abc": 1,
			"t":   "def",
		},
	})
	t.Log(r, err)
}
