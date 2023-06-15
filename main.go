package main

import (
	"fmt"

	"github.com/LoaiEsam37/443EMP/util"
)

func main() {
	Timeout, InsecureSkipVerify := util.SetConfig()
	url := "https://www.google.com"
	resp, err := util.Get(url, Timeout, InsecureSkipVerify)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}
