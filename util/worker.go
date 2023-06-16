package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func Worker(
	id int,
	DomainNames []string,
	timeout int,
	InsecureSkipVerify bool,
	ch chan<- string) {
	// defer to mark the task as done when the function finishes
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: InsecureSkipVerify,
		},
	}

	// make a seperate client to not interfere with other requests while multiprocessing
	client := http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}
	for _, DomainName := range DomainNames {
		resp, err := client.Get(DomainName)
		if err != nil {
			continue
		} else if resp.StatusCode == 200 {
			ch <- DomainName
		}
		defer resp.Body.Close()
	}
	close(ch)
	fmt.Println("Process", id+1, "is Done")
}
