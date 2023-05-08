package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d ", i)
		}
	}()

	go func() {
		for i := 'a'; i <= 'j'; i++ {
			fmt.Printf("%c ", i)
		}
	}()

	time.Sleep(2 * time.Second)
}
