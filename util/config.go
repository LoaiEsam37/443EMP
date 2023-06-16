package util

import (
	"flag"
)

type Config struct {
	Timeout         int
	TLSClientConfig struct {
		InsecureSkipVerify bool
	}
	FileInfo struct {
		Input            string
		Output           string
		LinesPerSubarray int
	}
}

func SetConfig() (int, bool, string, string, int) {
	var config Config

	flag.IntVar(&config.Timeout, "timeout", 4, "timeout in seconds")
	flag.BoolVar(&config.TLSClientConfig.InsecureSkipVerify, "insecure", true, "Skip TLS certificate verification")
	flag.StringVar(&config.FileInfo.Input, "input", "", "Path to Input File")
	flag.StringVar(&config.FileInfo.Output, "output", "", "Path to Output File")
	flag.IntVar(&config.FileInfo.LinesPerSubarray, "lines", 1000, "Number of URLs per Process")

	flag.IntVar(&config.Timeout, "t", 3, "shortcut for -timeout")
	flag.BoolVar(&config.TLSClientConfig.InsecureSkipVerify, "v", true, "shortcut for -insecure")
	flag.StringVar(&config.FileInfo.Input, "i", "", "shortcut for -input")
	flag.StringVar(&config.FileInfo.Output, "o", "", "shortcut for -output")
	flag.IntVar(&config.FileInfo.LinesPerSubarray, "n", 1000, "shortcut for -lines")

	flag.Parse()

	if config.FileInfo.Input == "" {
		panic("Error: input file path is not specified")
	} else if config.FileInfo.Output == "" {
		panic("Error: output file path is not specified")
	}

	return config.Timeout, config.TLSClientConfig.InsecureSkipVerify, config.FileInfo.Input, config.FileInfo.Output, config.FileInfo.LinesPerSubarray
}
