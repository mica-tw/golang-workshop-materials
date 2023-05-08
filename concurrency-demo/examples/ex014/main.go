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

func generator(quit chan string) <-chan string { // HL
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case ch <- id(): // Do nothing
			case v := <-quit:
				log.Println("[Generator]", "Signal from main:", v, "... Generator stopping...")
				quit <- "generator quitting"
				return
			}
		}
	}()
	return ch
}

// Receive on quit channel
func main() {
	quit := make(chan string)
	ch := generator(quit)

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Received:", v)
	}

	quit <- "quit" // Signal generator, to quit // HL

	// Receive on the quit channel to wait for generator to quit
	log.Println("[Main]", "Signal from generator:", <-quit) // HL
}
