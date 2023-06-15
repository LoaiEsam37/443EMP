package util

import (
	"fmt"
	"os"
	"sync"
)

func MultiProcessingHandler(urls [][]string, timeout int, InsecureSkipVerify bool) {

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	var mutex sync.Mutex

	// start the workers
	for j := 1; j < len(urls); j++ {
		wg.Add(1)
		go Worker(j, &wg, urls[j], timeout, InsecureSkipVerify, &mutex, file)
	}

	// use a separate goroutine to wait for the workers to finish
	go func() {
		wg.Wait()
	}()

	fmt.Println("all tasks finished!")
}
