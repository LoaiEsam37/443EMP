package util

import (
	"crypto/tls"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	return client.Get(url)
}
