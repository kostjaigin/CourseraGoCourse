package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// using bufio reader as alternative to fmt.Scan, because Scan stores input separated by white-spaces
	// and would not work on example "I d skd a efju N"
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please, enter text: ")

	userInput, _ := reader.ReadString('\n')

	lowercaseTrimmedUserInput := strings.TrimSpace(strings.ToLower(userInput))

	var startsWithI = strings.HasPrefix(lowercaseTrimmedUserInput, "i")

	var endsWithN = strings.HasSuffix(lowercaseTrimmedUserInput, "n")

	var containsA = strings.ContainsRune(lowercaseTrimmedUserInput, 'a')

	if startsWithI && endsWithN && containsA {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
