package main

import (
	"fmt"
	"strconv"
)

// Define Person Struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, myself " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	p1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	// fmt.Println(p1)
	// p1.age++
	// fmt.Println(p1)
	fmt.Println(p1.greet())
	p1.hasBirthday()
	p1.getMarried("Williams")
	fmt.Println(p1.greet())
}
