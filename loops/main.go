package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 5; i++ {
	fmt.Printf("Number: %d\n", i)
	}

	// fizzbuzz
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Printf("fizzbuzz: %d\n", i)
		} else if i%5 == 0 {
			fmt.Printf("buzz: %d\n", i)
		} else if i%3 == 0 {
			fmt.Printf("fizz: %d\n", i)
		}
	}
}
