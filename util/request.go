package util

import (
	"crypto/tls"
	"net/http"
	"time"
)

func Get(url string, timeout int, InsecureSkipVerify bool) (*http.Response, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: InsecureSkipVerify,
		},
	}

	client := &http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}

	return client.Get(url)
}
