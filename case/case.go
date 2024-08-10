package main

import (
	"fmt"
)

func main() {
	name := "hasan"
	switch {
	case name == "Hasan":
		fmt.Println("Hasan")

	case name == "hasan":
		fmt.Println("hasan")
	}
}
