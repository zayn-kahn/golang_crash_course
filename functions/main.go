package main

import (
	"fmt"
)

func greeting (name string) string {
	return "Hello, " + name + "!"
}

func sum (a, b int) int {
	fmt.Println( a,"+", b, "=", a+b)
	return a + b
}

func main() {

	fmt.Println(greeting("josh"))
	sum(5, 5)
	
}
