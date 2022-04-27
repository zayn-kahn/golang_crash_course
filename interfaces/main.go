package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	Circle := circle{5}
	rect := rectangle{10, 5}

	fmt.Printf("Circle Area: %f\n", getArea(Circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rect))
}
