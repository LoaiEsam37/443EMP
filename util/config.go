package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Timeout         int `yaml:"timeout"`
	TLSClientConfig struct {
		InsecureSkipVerify bool `yaml:"InsecureSkipVerify"`
	} `yaml:"TLSClientConfig"`
	FileInfo struct {
		Filename         string `yaml:"filename"`
		LinesPerSubarray int    `yaml:"linesPerSubarray"`
	} `yaml:"fileInfo"`
}

func SetConfig() (int, bool, string, int) {
	file, err := os.Open("./config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	var Timeout = config.Timeout
	var InsecureSkipVerify = config.TLSClientConfig.InsecureSkipVerify
	var Filename = config.FileInfo.Filename
	var LinesPerSubarray = config.FileInfo.LinesPerSubarray

	fmt.Println("Timeout: ", Timeout)
	fmt.Println("InsecureSkipVerify: ", InsecureSkipVerify)

	return Timeout, InsecureSkipVerify, Filename, LinesPerSubarray
}
