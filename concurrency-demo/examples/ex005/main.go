package main

import (
	"fmt"
	"time"
)

type SendFn func(int) int

func sender(ch chan int, fn SendFn) {
	for i := 0; i < 10; i++ {
		ch <- fn(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go sender(ch1, func(i int) int { return 2*i + 1 })
	go sender(ch2, func(i int) int { return 2 * i })
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("Received %v from channel 1\n", v1)
		case v2 := <-ch2:
			fmt.Printf("Received %v from channel 2\n", v2)
		default:
			fmt.Printf("No communication...\n")
			time.Sleep(1 * time.Second)
		}
	}
}
