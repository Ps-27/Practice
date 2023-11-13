package main

import (
	"fmt"
	"sort"
)

// interface
type AgeFactor []Human

type Human struct {
	name string
	age  int
}

func (p AgeFactor) Len() int {
	return len(p)
}
func (p AgeFactor) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p AgeFactor) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}

// interface
type MarksFactor []stud

// student struct
type stud struct {
	name  string
	marks int
}

func (m MarksFactor) Len() int {
	return len(m)
}
func (m MarksFactor) Less(i, j int) bool {
	return m[i].marks > m[j].marks
}
func (m MarksFactor) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func main() {
	person := []Human{
		{"Sachin", 30},
		{"Sanni", 20},
		{"silpa", 23},
	}
	sort.Sort(AgeFactor(person))
	fmt.Println(person)

	students := []stud{
		{"prity", 80},
		{"pooja", 90},
	}
	sort.Sort(MarksFactor(students))
	fmt.Println(students)
}
