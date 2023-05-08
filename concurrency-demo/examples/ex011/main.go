package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func generator() <-chan string { // HL
	ch := make(chan string)
	go func() {
		for {
			hash := sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
			ch <- hex.EncodeToString(hash[:])
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

// Timeout using select (example)
func main() {
	ch := generator() // HL
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(500 * time.Millisecond): // Timeout for each message // HL
			fmt.Println("Response took more than 500 Millisecond. Timing out...")
			return
		}
	}
}
