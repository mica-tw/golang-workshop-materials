package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func idGenerator() <-chan string { // HL
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

func main() {
	ch := idGenerator() // HL
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println("Received id:", value)
	}
	fmt.Println("Leaving..")
}
