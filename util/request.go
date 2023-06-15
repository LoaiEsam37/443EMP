package util

import (
	"crypto/tls"
	"net/http"
	"time"
)

func Get(urls []string, timeout int, InsecureSkipVerify bool) (Domains []string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: InsecureSkipVerify,
		},
	}

	// make a seperate client to not interfere with other requests while multiprocessing
	client := http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}
	var vaildUrls []string
	for _, url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			continue
		} else if resp.StatusCode == 200 {
			vaildUrls = append(vaildUrls, url)
		}
		defer resp.Body.Close()
	}

	return vaildUrls
}
