package main

import (
	"fmt"
	"os"
)

func factorial(n int) int {
	if n < 0 {
		fmt.Println("Invalid input Enter valid number:")
		os.Exit(1)
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Print("Enter number:")

	var n int
	fmt.Scan(&n)

	fmt.Printf("\nFactorial of %d is %d", n, factorial(n))
}
