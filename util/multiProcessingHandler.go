package util

import (
	"fmt"
	"os"
)

func MultiProcessingHandler(urls [][]string, timeout int, InsecureSkipVerify bool, Output string) {

	file, err := os.OpenFile(Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	numWorkers := len(urls)

	channels := make([]chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan string)
	}

	for i, ch := range channels {
		go Worker(i, urls[i], timeout, InsecureSkipVerify, ch)
		fmt.Println("Initiating scan for Process", i+1, "...")
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
			break
		}
	}
	println("Mission accomplished! All targets have been scanned.")
}
