package main

import "fmt"

func main() {
	i := 42
	j := &i // j is a pointer/memory address (means where this value is stored in memory) to i
	fmt.Println(i)
	fmt.Println(j)
	// use * to read the value of the pointer
	fmt.Printf("%T\n%T\n%T\n", i, *j, j)

	// change the value of the pointer
	*j = 21
	fmt.Println(i)
}
