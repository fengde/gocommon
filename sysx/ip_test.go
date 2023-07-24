package sysx

import (
	"testing"
)

func TestGetInternetIP(t *testing.T) {
	t.Log(GetInternetIP())
}

func TestGetLocalIP(t *testing.T) {
	t.Log(GetLocalIP())
}
