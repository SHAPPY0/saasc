package main

import (
	"fmt"
	"github.com/shappy0/saasc/cmd"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("%v.\n", err)
		}
	}()
	cmd.Run()
}