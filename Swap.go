package main

import "fmt"

func SwapContent(obj []string) {
	for i, j := 0, len(obj)-1; i < j; i, j = i+1, j-1 {
		obj[i], obj[j] = obj[j], obj[i]
	}
}

func main() {

	obj := []string{"prity", "sinha", "pooja", "sinha"}
	SwapContent(obj)
	fmt.Println(obj)
}
