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
		Input            string `yaml:"input"`
		Output           string `yaml:"output"`
		LinesPerSubarray int    `yaml:"linesPerSubarray"`
	} `yaml:"fileInfo"`
}

func SetConfig() (int, bool, string, string, int) {
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
	var Input = config.FileInfo.Input
	var Output = config.FileInfo.Output
	var LinesPerSubarray = config.FileInfo.LinesPerSubarray

	fmt.Println("InsecureSkipVerify: ", InsecureSkipVerify)
	fmt.Println("LinesPerSubarray: ", LinesPerSubarray)
	fmt.Println("Timeout: ", Timeout)
	fmt.Println("Input: ", Input)
	fmt.Println("Output: ", Output)

	return Timeout, InsecureSkipVerify, Input, Output, LinesPerSubarray
}
