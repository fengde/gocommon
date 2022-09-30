package tlsx

import (
	"crypto/tls"
	"net/http"

	"github.com/fengde/gocommon/timex"
)

// 获取网站的证书过期时间
func TLSExpireTime(seedUrl string) (string, error) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	resp, err := client.Get(seedUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if len(resp.TLS.PeerCertificates) > 0 {
		return timex.Time2String(resp.TLS.PeerCertificates[0].NotAfter), nil
	}

	return "", nil
}
