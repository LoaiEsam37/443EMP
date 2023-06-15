package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndSplitFile(filename string, linesLimit int) ([][]string, error) {
	// Open the file with the specified filename
	file, err := os.Open(filename)
	if err != nil {
		// If there was an error opening the file, return nil and the error
		return nil, err
	}
	// Defer closing the file until the end of the function
	defer file.Close()

	// Create a variable to hold the lines of the file
	var lines [][]string
	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	// Create a counter variable to keep track of the number of lines in a sub-array
	i := 0
	// Create a variable to hold the lines of a sub-array
	var sublines []string
	fmt.Print("[ In Progress ] Spliting Domain Names")
	// Iterate over the lines in the file
	for scanner.Scan() {
		// Append the current line to the sub-array
		sublines = append(sublines, scanner.Text())
		// Increment the line counter
		i++
		// If the sub-array contains 1000 lines, append it to the main array and create a new sub-array
		if i == 1000 {
			lines = append(lines, sublines)
			sublines = nil
			i = 0
		}
	}
	// If the last sub-array contains less than 1000 lines, append it to the main array
	if len(sublines) > 0 {
		lines = append(lines, sublines)
	}
	fmt.Print("\r")
	fmt.Print("                                          ")
	fmt.Print("\r")
	fmt.Println("[ Done! ] Spliting Domain Names")

	// Return the main array of sub-arrays
	return lines, nil
}
