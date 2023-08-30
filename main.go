package main

import (
	"fmt"
	"github.com/go-helios/eos/cmd/eos"
)

func main() {
	err := eos.Execute()
	if err != nil {
		fmt.Println("执行错误:", err.Error())
	}
}
