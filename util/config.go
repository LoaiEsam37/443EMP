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
}

func SetConfig() (int, bool) {
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

	fmt.Println("Timeout: ", Timeout)
	fmt.Println("InsecureSkipVerify: ", InsecureSkipVerify)

	return Timeout, InsecureSkipVerify
}
