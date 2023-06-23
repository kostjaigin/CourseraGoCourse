package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// using bufio reader as alternative to fmt.Scan, because Scan stores input separated by spaces //
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter your name: ")
	userInput, _ := reader.ReadString('\n')
	var username = strings.TrimSpace(userInput)
	fmt.Printf("Great, now please enter your address in one row: ")
	userInput, _ = reader.ReadString('\n')
	var address = strings.TrimSpace(userInput)

	userMap := map[string]string{
		"name":    username,
		"address": address,
	}

	userMapJson, _ := json.Marshal(userMap)
	fmt.Printf(string(userMapJson))
}
