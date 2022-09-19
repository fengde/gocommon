package proxyx

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// // 直接将请求代理到address
// func Handler(address string) http.Handler {
// 	u, _ := url.Parse(address)
// 	proxy := httputil.NewSingleHostReverseProxy(u)

// 	return &ProxyHandler{proxy: proxy}
// }

// type ProxyHandler struct {
// 	proxy *httputil.ReverseProxy
// }

// func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.proxy.ServeHTTP(w, r)
// }

// 直接将请求代理到address
func HandlerFunc(address string) http.HandlerFunc {
	u, _ := url.Parse(address)
	proxy := httputil.NewSingleHostReverseProxy(u)

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
