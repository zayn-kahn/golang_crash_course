package main

import (
	"fmt"
)

func main() {
	// fruitArr:= [2]string{"apple","orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])

	fruitSlice:= []string{"apple","orange","banana","grape"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
