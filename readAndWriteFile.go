package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// READ FILE //

	filepath := "/Users/konstantin.igin/desktop/test_users.csv"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("An error occured while opening the file: %s, error: %s\n", filepath, err)
		os.Exit(1)
	}
	defer file.Close()

	var values []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("An error occured while scanning line: %s, err: %s\n", line, err)
			os.Exit(1)
		}
		values = append(values, value)
	}

	// for i, v := range values {
	// 	fmt.Printf("%d: %d\n", i, v)
	// }

	// WRITE INTO CSV FILE //
	outputFilepath := "/Users/konstantin.igin/desktop/users.csv"
	outputFile, err := os.Create(outputFilepath)
	if err != nil {
		fmt.Printf("Error creating file: %s, err: %s\n", outputFilepath, err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	var records [][]string
	for _, val := range values {
		records = append(records, []string{strconv.Itoa(val)})
	}

	writer.WriteAll(records)
	writer.Flush()
}
