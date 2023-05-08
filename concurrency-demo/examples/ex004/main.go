package main

import (
	"fmt"
	"time"
)

func squareNumbers(ch chan int) { // Sender
	for i := 0; i < 10; i++ {
		ch <- i * i
	}
	time.Sleep(2 * time.Second)
	close(ch)
}

func main() {
	ch := make(chan int)
	go squareNumbers(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
