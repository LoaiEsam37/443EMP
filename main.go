package main

import (
	"github.com/LoaiEsam37/443EMP/util"
)

func main() {
	Timeout, InsecureSkipVerify, Filename, LinesPerSubarray := util.SetConfig()
	urls, err := util.ReadAndSplitFile(Filename, LinesPerSubarray)
	if err != nil {
		panic(err)
	}
	Domains := util.Get(urls, Timeout, InsecureSkipVerify)
	if err != nil {
		panic(err)
	}

}
