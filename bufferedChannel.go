package main

import (
	"fmt"
	"sync"
)

func listenTochan(ch chan int, wg sync.WaitGroup) {
	wg.Done()

	for {
		i := <-ch
		fmt.Println("Got", i, "from channel inside 1st routine")

		// time.Sleep(1 * time.Second)
	}
}

func listenTochan1(ch chan int, wg sync.WaitGroup) {
	wg.Done()
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel inside 2nd routine")
	}
}
func main() {

	// ch := make(chan int) //unbuffered
	ch := make(chan int, 4)

	var wg sync.WaitGroup

	wg.Add(2)
	go listenTochan(ch, wg)

	go listenTochan1(ch, wg)
	for i := 0; i <= 10; i++ {
		fmt.Println("sending", i, " to channel...")
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println("Done!")
	// close(ch)
}
