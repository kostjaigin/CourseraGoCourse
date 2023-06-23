package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := ReadUserInput()
	fmt.Printf("user input (unsorted): %v\n", numbers)
	BubbleSort(numbers)
	fmt.Printf("user input (sorted): %v\n", numbers)
}

func BubbleSort(unsortedInput []int) {
	wrongs := len(unsortedInput)
	for wrongs != 0 {
		unsortedPart := unsortedInput[:wrongs]
		for i := 0; i < len(unsortedPart)-1; i++ {
			current := unsortedPart[i]
			next := unsortedPart[i+1]
			if current > next {
				Swap(unsortedPart, i)
			}
		}
		wrongs--
	}
}

// Swap swaps element at index with elements at index+1 in given slice //
func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func ReadUserInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter up to 10 integer numbers separated by spaces:\n")
	userInput, _ := reader.ReadString('\n')
	cleanedSplit := strings.Split(strings.TrimSpace(userInput), " ")
	numbers := make([]int, 0, 10)
	for _, numberString := range cleanedSplit {
		if len(numbers) >= cap(numbers) {
			fmt.Printf("%s and following numbers will be ignored, more than 10 numbers entered.\n", numberString)
			break
		}
		var number, err = strconv.Atoi(numberString)
		if err != nil {
			fmt.Printf("%s is not a valid number, ignoring it\n", numberString)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}
