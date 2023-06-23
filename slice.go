package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var userInputSlice = make([]int, 0, 3)

	for {
		fmt.Printf("Please, enter an integer number to store or 'X' to exit:\n")
		// scan user input
		var userInput string
		fmt.Scan(&userInput)
		// if user input equals X then exit
		if userInput == "X" {
			break
		}
		// convert user input into integer
		userInputNumber, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Printf("Could not convert your \"input\" %s to integer. ", userInput)
			continue
		}
		// store int(X) in slice
		userInputSlice = append(userInputSlice, userInputNumber)
		// sort the slice
		sort.Slice(userInputSlice, func(i, j int) bool {
			return userInputSlice[i] < userInputSlice[j]
		})
		// print the current state of the slice
		fmt.Printf("Currently stored numbers: %v\n", userInputSlice)
	}
}
