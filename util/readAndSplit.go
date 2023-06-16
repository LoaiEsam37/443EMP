package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndSplitFile(filename string, linesLimit int) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	var lines [][]string
	var sublines []string

	fmt.Println("Splitting domain names into their component parts...")
	for scanner.Scan() {
		sublines = append(sublines, scanner.Text())
		i++
		if i == linesLimit {
			lines = append(lines, sublines)
			sublines = nil
			i = 0
		}
	}
	if len(sublines) > 0 {
		lines = append(lines, sublines)
	}
	fmt.Println("Domain name splitting process has completed successfully.")
	return lines, nil
}
