package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    messages, error := greetings.HelloAll([]string{"Hasan", "Ali"})
	if error != nil{
		log.Fatal(error)
	}
	fmt.Println(messages)
	
}