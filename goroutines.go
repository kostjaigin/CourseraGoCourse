package main

import (
	"fmt"
	"time"
)

var counter int // communication layer for two goroutines, assigned to 0 initially

func main() {

	// We launch two goroutines, each incrementing shared global variable counter

	go func() { // I had to read out how to do this, cause we haven't seen this in the course yet...
		counter++
	}()

	go func() {
		counter++
	}()

	// they both might read the same initial value of 0 and increment it to 1 simultaniously

	time.Sleep(1 * time.Second) // allow them to complete

	// in some (rare) cases we will see 1 instead of 2
	fmt.Println(counter)

}
