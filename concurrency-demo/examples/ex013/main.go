package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func id() string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(hash[:])
}

func generator(quit chan bool) <-chan string { // HL
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case ch <- id(): // Do nothing
			case <-quit:
				log.Println("Generator stopping")
				close(ch)
				return
			}
		}
	}()
	return ch
}

// Quit channel (example)
func main() {
	quit := make(chan bool) // HL
	ch := generator(quit)

	go func() {
		time.Sleep(5 * time.Second)
		quit <- true
	}()

	for v := range ch { // Receive values from channel until it is closed // HL
		log.Println("Received:", v)
	}
	_, ok := <-ch // Test if channel is closed // HL

	if !ok {
		log.Println("Generator channel closed.")
	}

	log.Println("Leaving ...")

}
