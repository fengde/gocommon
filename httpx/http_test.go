package httpx

import (
	"testing"
)

func TestPostXML(t *testing.T) {
	r, err := PostXML("http://httpbin.org/anything", nil, "<abc></abc>")
	t.Log(r, err)
}

func TestPostForm(t *testing.T) {
	r, err := PostForm("http://httpbin.org/anything", nil, map[string]interface{}{
		"abc": 1,
		"t":   "def",
	})
	t.Log(r, err)
}
