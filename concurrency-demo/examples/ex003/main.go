package main

import "fmt"

func squareNumbers(ch chan int) { // Sender
	for i := 0; i < 10; i++ {
		ch <- i * i
	}
}

func main() {
	ch := make(chan int)
	go squareNumbers(ch)

	for i := 0; i < 10; i++ { // Receiver
		v := <-ch
		fmt.Printf("%d ", v)
	}
}
