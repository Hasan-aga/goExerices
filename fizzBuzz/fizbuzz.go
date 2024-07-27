package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(proceduralFizzBuzz(100))
}

func proceduralFizzBuzz(max int) string {
	var result []string
	for i := 1; i <= max; i++ {
		word := ""
		if i%3 == 0 {
			word += "Fizz"
		}
		if i%5 == 0 {
			word += "Buzz"
		}
		if len(word) == 0 {
			word += fmt.Sprint(i)
		}
		result = append(result, word)

	}
	return strings.Join(result, "\n")
}
