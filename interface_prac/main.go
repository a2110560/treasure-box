package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rect struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * math.Pi * c.radius
}

func main() {
	c1 := circle{3.0}
	r1 := rect{5, 7}
	shapes := []shape{c1, r1}
	for _, item := range shapes {
		fmt.Println(item.area())
	}
}
