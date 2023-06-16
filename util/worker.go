package util

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func Worker(
	id int,
	wg *sync.WaitGroup,
	DomainNames []string,
	timeout int,
	InsecureSkipVerify bool,
	mutex *sync.Mutex,
	file *os.File) {
	// defer to mark the task as done when the function finishes
	defer wg.Done()
	// perform some task
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: InsecureSkipVerify,
		},
	}

	// Acquire the lock before writing to the file
	mutex.Lock()
	defer mutex.Unlock()
	// make a seperate client to not interfere with other requests while multiprocessing
	client := http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}
	fmt.Println("worker: ", id, "started")
	for _, DomainName := range DomainNames {
		resp, err := client.Get(DomainName)
		if err != nil {
			continue
		} else if resp.StatusCode == 200 {
			_, err := file.Write([]byte(DomainName))
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
		}
		defer resp.Body.Close()
	}
	fmt.Println("worker", id, "finished")
}
