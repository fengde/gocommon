package proxyx

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 直接将请求代理到address
func HandlerFunc(address string) http.HandlerFunc {
	u, _ := url.Parse(address)
	proxy := httputil.NewSingleHostReverseProxy(u)

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
