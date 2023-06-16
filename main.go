package main

import (
	"github.com/LoaiEsam37/443EMP/util"
)

func main() {

	println("     _     _   _          ____              ")
	println("    | |__ | |_| |_ _ __  / ___|_   _ _ __  ")
	println("    | '_ \\| __| __| '_ \\| |  _| | | | '_ \\ ")
	println("    | | | | |_| |_| |_) | |_| | |_| | | | |")
	println("    |_| |_|\\__|\\__| .__/ \\____|\\__,_|_| |_|")
	println("                  |_|                      ")
	println("")

	Timeout, InsecureSkipVerify, Input, Output, LinesPerSubarray := util.SetConfig()
	println("Initiating domain name parsing process...")
	urls, err := util.ReadAndSplitFile(Input, LinesPerSubarray)
	if err != nil {
		panic(err)
	}
	println("[!] Attention user! This tool will require approximately", (len(urls) * 50 / 1024), "MB of RAM resources to operate at optimal performance levels Please ensure that your system has sufficient resources available before launching the tool.")
	util.MultiProcessingHandler(urls, Timeout, InsecureSkipVerify, Output)

}
