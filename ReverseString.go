package main

import "fmt"

func reverse(s string) {
	rev := []byte(s)
	l := len(rev)
	for i := 0; i < l/2; i++ {
		rev[i], rev[l-i-1] = rev[l-i-1], rev[i]
	}
	fmt.Println("Reverse of string is :", string(rev))
}
func reverse1(str string) {
	byteStr := []byte(str)
	l := len(byteStr)
	for i := 0; i < l/2; i++ {
		byteStr[i], byteStr[l-i-1] = byteStr[l-i-1], byteStr[i]
	}
	fmt.Printf("Reverse of %s is %s", str, string(byteStr))
}

func re(str string) {
	bs := []byte(str)
	l := len(str)
	for i := 0; i < l/2; i++ {
		bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
	}
	fmt.Println("  inside new")
	fmt.Println(string(bs))
}
func main() {
	var str string

	fmt.Println("Enter string :")
	fmt.Scan(&str)

	fmt.Println("Original string is :", str)

	reverse(str)
	reverse1(str)
	re(str)
}
