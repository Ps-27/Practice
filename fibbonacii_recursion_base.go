package main

import "fmt"

func fiboSeries(fNum int, SNum int, n int) {
	// fmt.Println(" inside fibo with ", fNum, " ", SNum, " ", n)
	temp := fNum + SNum
	if temp <= n {
		fmt.Printf(" %d ", temp)
		fiboSeries(SNum, temp, n)
	}

}
func main() {

	fNum := 0
	SNum := 1
	n := 24
	fmt.Print(fNum, " ", SNum)
	fiboSeries(fNum, SNum, n)
}
