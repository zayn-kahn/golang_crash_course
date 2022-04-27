package main

import "fmt"

func main() {
	// var name = "Gopher"
	var age = 37
	name := "Gopher" // short hand Can only be used in func scope
	fmt.Println(name, age)
	fmt.Printf("%T\n %T\n", name, age)
}
