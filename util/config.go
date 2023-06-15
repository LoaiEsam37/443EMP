package util

import (
	"fmt"
	"io/ioutil"

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
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
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
