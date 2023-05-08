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

// Timeout for the whole conversation
func main() {
	ch := generator()
	timeOut := time.After(5 * time.Second) // HL

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeOut:
			fmt.Println("No more ids required. Leaving ...")
			return
		}
	}

}
