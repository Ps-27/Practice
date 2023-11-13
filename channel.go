package main

import "fmt"

func square(ch chan int) {
	n := 7
	s := n * n
	fmt.Println(<-ch)
	fmt.Println(s)
	// ch <- 8

}

func sub(ch1, ch2 chan int) {
	fmt.Println("inside sub")
	n1 := <-ch1
	n2 := <-ch2
	fmt.Println("value", n1, n2)
	ch1 <- n1 - n2
}
func main() {

	fmt.Println("Enter into main ")
	ch := make(chan int)
	go square(ch)
	ch <- 5
	ch1 := make(chan int)
	ch2 := make(chan int)
	go add(ch1, ch2)
	ch1 <- 4
	ch2 <- 5

	fmt.Println("add :", <-ch1)

	ch3 := make(chan int)
	ch4 := make(chan int)
	go sub(ch3, ch4)
	ch3 <- 7
	ch4 <- 8
	res := <-ch3
	fmt.Println(res)

}

func add(ch1, ch2 chan int) {
	s := <-ch1 + <-ch2
	ch1 <- s
}
