package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) { // HL
	fmt.Printf("WORKER: %d: starting\n", id)
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("WORKER: %d: finishing\n", id)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg) // NOTE: WaitGroup should not be passed by value
	}
	wg.Wait()
}
