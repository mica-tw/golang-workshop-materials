package main

import (
	"fmt"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func worker(id int, sem Semaphore) {
	sem.Acquire()
	defer sem.Release()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Semaphore
func main() {
	sem := NewSemaphore(2)

	for i := 1; i <= 10; i++ {
		go worker(i, sem)
	}
	time.Sleep(3 * time.Second)
}
