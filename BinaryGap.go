package main

import "fmt"

func BinaryGap(n int) int {
	fmt.Println(n)
	t := []rune(n)
	fmt.Println(t)
	return 0
}
func main() {

	n := 10100100
	res := BinaryGap(n)
	fmt.Println(res)
}
