package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func HelloAll(names []string)(map[string]string,error){
	messages := make(map[string]string)
	for _,name := range names{
		message, error := Hello(name)
		if error != nil{
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}
func Hello(name string) (string, error){
	    // If no name was given, return an error with a message.
		if name == "" {
			return "", errors.New("empty name")
		}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string{
	formats :=[]string{
		"Hello, %v",
		"Merhaba, %v",
		"Selam, %v",
	}
	return formats[rand.Intn(len(formats))]
}