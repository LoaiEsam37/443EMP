package util

import (
	"fmt"
	"os"
	"sync"
)

func MultiProcessingHandler(urls [][]string, timeout int, InsecureSkipVerify bool, Output string) {

	file, err := os.OpenFile(Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	numWorkers := len(urls)

	// create a slice of channels
	channels := make([]chan string, numWorkers)

	// create a channel for each worker
	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan string)
	}
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	// start the workers
	for i, ch := range channels {
		go Worker(i, urls[i], timeout, InsecureSkipVerify, ch)
		fmt.Println("Process", i+1, "is scanning", len(urls[i]), "Domain Names")
	}
	for {
		allClosed := true
		for _, ch := range channels {
			select {
			case vaildUrl, open := <-ch:
				if open {
					allClosed = false
					dataWithNewline := vaildUrl + "\n"
					_, err = file.WriteString(dataWithNewline)
					if err != nil {
						panic(err)
					}
				}
			default:
				allClosed = false
			}
		}
		if allClosed {
			break // all channels are closed, exit the loop
		}
	}
}
