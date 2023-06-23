package main

import (
	"fmt"
	"os"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
	return fn
}

func main() {
	a, v0, s0 := readInitialValues()
	findDisplacement := GenDisplaceFn(a, v0, s0)

	var userInput string
	fmt.Printf("Please, enter the time t value:\n")
	fmt.Scan(&userInput)
	t := readFloat(userInput)

	result := findDisplacement(t)
	fmt.Printf("Displacement after the entered time value: %f\n", result)
}

func readInitialValues() (float64, float64, float64) {
	var userInput string
	fmt.Printf("Please, enter the acceleration value:\n")
	fmt.Scan(&userInput)
	a := readFloat(userInput)
	fmt.Printf("Please, enter the initial speed value:\n")
	fmt.Scan(&userInput)
	v0 := readFloat(userInput)
	fmt.Printf("Please, enter the initial dispacement value:\n")
	fmt.Scan(&userInput)
	s0 := readFloat(userInput)
	return a, v0, s0
}

func readFloat(userInput string) float64 {
	result, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Printf("Could not convert your input %s into float64, try again.", userInput)
		os.Exit(1)
	}
	return result
}
