package main

import (
	"fmt"
)

func fibonacii(Fnum int, Snum int, n int) int {
	if n == 0 {
		return 0
	}
	ThirdNum := Fnum + Snum
	fmt.Print(ThirdNum, " ")
	Fnum = Snum
	Snum = ThirdNum
	return fibonacii(Fnum, Snum, n-1)
}

func main() {
	fmt.Println("Welcome!!")
	Fnum := 0
	Snum := 1
	n := 5
	fmt.Print(Fnum, " ", Snum, " ")
	fibonacii(Fnum, Snum, n-2)

}
