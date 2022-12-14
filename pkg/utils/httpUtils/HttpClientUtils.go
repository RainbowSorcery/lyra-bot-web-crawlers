package httpUtils

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

var httpClient *resty.Client
var lock sync.Mutex

func GetHttpClientInstance() *resty.Client {
	lock.Lock()
	defer lock.Unlock()
	if httpClient == nil {
		httpClient = resty.New()
		httpClient.SetProxy("http://127.0.0.1:7890")
		httpClient.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	}

	return httpClient
}
