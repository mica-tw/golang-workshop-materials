package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printNumbers() // HL
	go printLetters() // HL
	time.Sleep(2 * time.Second)
}
