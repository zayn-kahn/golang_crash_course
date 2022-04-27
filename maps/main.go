package main

import (
	"fmt"
)

func main() {
	// maps
	// email := make(map[string]string)
	// // assign key value pairs
	// email["golang"] = "golang@gamil.com"
	// email["python"] = "python@gmail.com"
	// email["rust"] = "rust@gmail.com"

	// declare map and add keyvalue
	email := map[string]string{
		"golang": "golang@gmail.com",
		"python": "python@gmail.com",
		"rust":   "rust@gmail.com"}

	fmt.Println(email)
	fmt.Println(email["golang"])

	delete(email, "python")
	fmt.Println(email)
}
