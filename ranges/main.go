package main

import "fmt"

func main() {
	ids := []int{11, 42, 23, 74, 125}
	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i, id)
	}
	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	sum := 0
	for _, id := range ids {
		sum += id
		fmt.Println("Sum is:", sum)
	}
	fmt.Println("Sum is:", sum)
	email := map[string]string{
		"golang": "golang@gmail.com",
		"python": "python@gmail.com",
		"rust":   "rust@gmail.com",
	}
	for k, v := range email {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k:= range email {
		fmt.Println("Name: "+ k)
	}
}
