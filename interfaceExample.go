package main

import (
	"fmt"
)


type shape interface {
	getArea() float32
}
type circle struct {
	redius float32
}
type rectangle struct {
	length  int
	breadth int
}

func (s circle) getArea() float32 {
	return 3.14 * s.redius * s.redius
}

func (s rectangle) getArea() float32 {
	return float32(s.length) * float32(s.breadth)
}
func printArea(s shape) {
	fmt.Println("Area :", s.getArea())
}
func main() {
	fmt.Println("Welcome for Interface Example!")
	c := circle{7.0}
	printArea(c)
	r := rectangle{4, 6}
	printArea(r)
}


