package base64x

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	url := "www.5lmh.com"
	fmt.Println(Encode([]byte(url)))
	fmt.Println(UrlEncode([]byte(url)))
}
