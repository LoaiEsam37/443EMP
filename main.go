package main

import (
	"github.com/LoaiEsam37/443EMP/util"
)

func main() {
	Timeout, InsecureSkipVerify, Input, Output, LinesPerSubarray := util.SetConfig()
	urls, err := util.ReadAndSplitFile(Input, LinesPerSubarray)
	if err != nil {
		panic(err)
	}
	util.MultiProcessingHandler(urls, Timeout, InsecureSkipVerify, Output)

}
