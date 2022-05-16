package luvcheckerlib

import (
	"net/http"
	"net/url"
)

type Proxy struct {
	Address string
	InUse   bool
	Banned  bool
}

//HttpProxyType returns a http proxy client
func (p *Proxy) HttpProxyType() *http.Client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(p.Address)
	}
	return &http.Client{Transport: &http.Transport{Proxy: proxy}}
}

//CheckProxyHttp checks if a proxy is alive and working by pinging it and checking the response code of the response it gets back.
//If the response code is 200, the proxy is alive and working.
func CheckProxyHttp(proxy string) bool {
	resp, err := http.Get("http://" + proxy)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

//CheckProxySocks5 checks if a proxy is alive and working by pinging it and checking the response code of the response it gets back.
//If the response code is 200, the proxy is alive and working.
func CheckProxySocks5(proxy string) bool {
	resp, err := http.Get("socks5://" + proxy)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

//CheckProxySocks4 checks if a proxy is alive and working by pinging it and checking the response code of the response it gets back.
//If the response code is 200, the proxy is alive and working.
func CheckProxySocks4(proxy string) bool {
	resp, err := http.Get("socks4://" + proxy)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
