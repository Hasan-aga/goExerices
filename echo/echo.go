package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		name := strings.Join(args[1:], " ")
		fmt.Println(name)
	} else {
		fmt.Println("No arguments passed")
		os.Exit(1)
	}
}
