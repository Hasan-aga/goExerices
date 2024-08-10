package main

import "fmt"

func main() {
	i := 10
	p := &i         // get address of
	fmt.Println(p)  // print address
	fmt.Println(*p) // print value (pointer to address)
}
