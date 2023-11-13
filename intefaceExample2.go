package main

import (
	"fmt"
)

// interface
type tank interface {
	Volume() float64
	Tarea() float64
}

//struct
type cirTank struct {
	radius float64
	height float64
}

func (c cirTank) Tarea() float64 {
	return 2*c.radius*c.height +
		2*3.14*c.radius*c.radius
}

func (c cirTank) Volume() float64 {
	return 3.14 * c.height * c.radius * c.radius

}
func Print(t tank) {
	fmt.Println("volume :", t.Volume())
	fmt.Println("TArea: ", t.Tarea())

}
func main() {
	fmt.Println("Inside main!!")
	t := cirTank{7, 1}
	Print(t)

}
