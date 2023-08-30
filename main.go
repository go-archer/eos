package main

import (
	"eos/cmd/eos"
	"fmt"
)

func main() {
	err := eos.Execute()
	if err != nil {
		fmt.Println("执行错误:", err.Error())
	}
}
