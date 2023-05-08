package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func word() string {
	const charset = "abcdefghijklmnopqrstuvwxyz"

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func upper(in chan string, out chan string) {
	for value := range in {
		out <- strings.ToUpper(value)
	}
	close(out)
}
func prefix(in chan string, out chan string) {
	for value := range in {
		out <- "id_" + value
	}
	close(out)
}

func dump(in chan string) {
	for value := range in {
		fmt.Println(value)
	}
}

// Daisy-chain example
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go upper(ch1, ch2)
	go prefix(ch2, ch3)
	go dump(ch3)

	for i := 0; i < 10; i++ {
		ch1 <- word()
		time.Sleep(1 * time.Second)
	}
	close(ch1)
}
