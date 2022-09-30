package tlsx

import "testing"

func TestTLSExpireTime(t *testing.T) {
	result, err := TLSExpireTime("https://www.geesunn.com")
	t.Log(result, err)
}
