package main

import (
	"fmt"
	"math"
)

func main() {

	var numberToTruncate float64

	fmt.Printf("Please, enter a floating point number to truncate: ")

	fmt.Scan(&numberToTruncate)

	var truncatedNumber = int(math.Trunc(numberToTruncate))
	// alternative:
	// truncatedNumber := int(math.Trunc(numberToTruncate))
	// truncatedNumber := int(numberToTruncate)

	fmt.Printf("Truncated integer version of %f is %d\n", numberToTruncate, truncatedNumber)
}
