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
	partitions := SplitSlice(numbers)

	// channel for results
	sortedParts := make(chan []int, 4)
	defer close(sortedParts)

	// process the partitions concurrently
	for i, partition := range partitions {
		go func(routineID int, unsorted []int, results chan []int) {
			fmt.Printf("go routine %d sorts: %v\n", routineID, unsorted)
			BubbleSort(unsorted)
			results <- unsorted
		}(i+1, partition, sortedParts)
	}

	// collect the results
	results := make([]int, 0, len(numbers))
	for range partitions {
		sorted := <-sortedParts
		results = append(results, sorted...)
	}
	// sort again (e.g., 1 9 6 5 4 3 8 7) --> ((1 9)(5 6)(3 4)(7 8))
	BubbleSort(results)
	fmt.Println(results)
}

func ReadUserInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter integer numbers (min. 4) separated by spaces:\n")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	cleanedSplit := strings.Split(strings.TrimSpace(userInput), " ")

	if len(cleanedSplit) < 4 {
		fmt.Println("Please, enter at least 4 numbers.")
		os.Exit(1)
	}

	numbers := make([]int, 0)
	for _, numberString := range cleanedSplit {
		var number, err = strconv.Atoi(numberString)
		if err != nil {
			fmt.Printf("%s is not a valid number, ignoring it.\n", numberString)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func SplitSlice(userInput []int) [][]int {
	quarter := len(userInput) / 4
	return [][]int{
		userInput[0:quarter],
		userInput[quarter : 2*quarter],
		userInput[2*quarter : 3*quarter],
		userInput[3*quarter:],
	}
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
