package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Worker(
	id int,
	DomainNames []string,
	timeout int,
	InsecureSkipVerify bool,
	ch chan<- string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: InsecureSkipVerify,
		},
	}

	client := http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}
	for _, DomainName := range DomainNames {
		parsedURL, err := url.Parse(DomainName)
		if err != nil {
			panic(err)
		}

		if !strings.HasPrefix(parsedURL.Host, "www.") {
			DomainName = "http://www." + parsedURL.Host + parsedURL.Path
		} else if parsedURL.Scheme == "" {
			DomainName = "http://" + DomainName
		}
		resp, err := client.Get(DomainName)
		if err != nil {
			continue
		} else if resp.StatusCode == 200 {
			ch <- DomainName
		}
		defer resp.Body.Close()
	}
	close(ch)
	fmt.Println("Mission accomplished for Process", id+1, ". Initiating cool-down sequence...")
}
