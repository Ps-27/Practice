package main

import (
	"fmt"
)

func main() {
	a := []int{2, 5, 9, 7, 1, 0}
	fmt.Println(a)

	// s, ss := 0, 0
	// if a[0] > a[1] {
	// 	s, ss = a[1], a[0]
	// } else {
	// 	s, ss = a[0], a[1]
	// }
	// for i, _ := range a {
	// 	if a[i] < s {
	// 		ss = s
	// 		s = a[i]
	// 	} else if a[i] < ss {
	// 		ss = a[i]
	// 	}
	// }

	s, ss := a[0], a[0]
	for i, _ := range a {
		if a[i] < s {
			ss = s
			s = a[i]
		} else if a[i] < ss {
			ss = a[i]
		}
	}
	fmt.Println("smallest:", s)
	fmt.Println("Second Smallest:", ss)
}
