package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
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
	fmt.Println("Scanning process for Process", id+1, "has begun")
	for _, DomainName := range DomainNames {
		parsedURL, err := url.Parse(DomainName)
		if err != nil {
			panic(err)
		}

		if parsedURL.Scheme == "" {
			DomainName = "http://" + DomainName
		}
		resp, err := client.Get(DomainName)
		if err != nil {
			continue
		} else if resp.StatusCode == 200 {
			ch <- DomainName
		} else if resp.StatusCode == 429 {
			fmt.Println("The server that was contacted has returned a 429 error code.")
		}
		defer resp.Body.Close()
	}
	close(ch)
	fmt.Println("Mission accomplished for Process", id+1, ". Initiating cool-down sequence...")
}
