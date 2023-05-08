package main

import (
	"log"
	"time"
)

// Signal channels
func worker1(ch chan struct{}) {
	log.Println("[worker 1]", "starting")
	// Do some work...
	time.Sleep(500 * time.Millisecond)
	log.Println("[worker 1]", "finishing")
	ch <- struct{}{} // Signal completion // HL
}

func worker2(ch chan struct{}) {
	<-ch // Wait for signal
	// Do some work...
	log.Println("[worker 2]", "starting")
	time.Sleep(500 * time.Millisecond)
	log.Println("[worker 2]", "finishing")
}

func main() {
	ch := make(chan struct{})

	go worker1(ch)
	go worker2(ch)

	// Wait for both workers to finish
	time.Sleep(2 * time.Second)
}
