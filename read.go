package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	var namesSlice = make([]Name, 0)

	fmt.Printf("Please, enter the filepath to the file to read the names from:\n")
	// scan user input //
	var userInput string
	fmt.Scan(&userInput)
	filepath := strings.TrimSpace(userInput)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("An error encountered trying to open the file %s, try again with another file. Error: %s\n", filepath, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// split by whitespace //
		data := strings.Split(line, " ")
		if len(data) != 2 {
			fmt.Printf("Given file %s contains a line %s which does not correspond to expected format\n", filepath, line)
			os.Exit(1)
		}
		// create struct //
		var obj Name
		copy(obj.fname[:], data[0])
		copy(obj.lname[:], data[1])
		// attach to slice //
		namesSlice = append(namesSlice, obj)
	}

	fmt.Printf("File %s read. It contains following names:\n", filepath)
	for _, name := range namesSlice {
		fmt.Printf("First Name: %s | Last Name: %s\n", name.fname, name.lname)
	}

}
