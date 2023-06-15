package main

import (
	"fmt"

	"github.com/LoaiEsam37/443EMP/util"
)

func main() {
	url := "www.example.com"
	resp, err := util.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}
