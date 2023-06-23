package main

import (
	"fmt"
	"strings"
)

// this is how it was supposed to work, but Scan reads input separating it by spaces (userInput only gets filled with values until the first space is faced)

func main() {

	var userInput string

	fmt.Printf("Please, enter text: ")

	fmt.Scan(&userInput)

	lowercaseUserInput := strings.ToLower(userInput)

	var startsWithI bool = strings.HasPrefix(lowercaseUserInput, "i")

	var endsWithN bool = strings.HasSuffix(lowercaseUserInput, "n")

	var containsA bool = strings.ContainsRune(lowercaseUserInput, 'a')

	if startsWithI && endsWithN && containsA {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
