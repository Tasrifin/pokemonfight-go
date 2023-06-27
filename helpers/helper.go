package helpers

import (
	"crypto/tls"
	"net/http"
	"time"
)

func ReqHTTP(method string, url string) (result *http.Response, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(10) * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	result, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
