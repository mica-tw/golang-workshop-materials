package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateFP() string {
	rand.Seed(time.Now().UnixNano())

	hash := md5.Sum([]byte("some str"))
	hashStr := hex.EncodeToString(hash[:])

	b := make([]byte, 16)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}

	return hashStr[:8] + string(b) + hashStr[24:]
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 16)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}

	return string(b)
}

type Algorithm func() string

// START OMIT
type Message struct {
	body string
	wait chan bool // HL
}

// STOP OMIT

func service(name string, fn Algorithm) <-chan Message { // Returns receive-only channel of Message. // HL
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.
	go func() {                  // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s [%d]: %s", name, i, fn()), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // Return the channel to the caller. // HL
}

func merge(inputs ...<-chan Message) <-chan Message { // AKA FanIn // HL
	c := make(chan Message)
	for _, input := range inputs {
		ch := input
		go func() {
			for {
				c <- <-ch
			}
		}()
	}
	return c
}
func main() {
	c := merge(service("ID", generateID), service("FINGERPRINT", generateFP)) // HL
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.body)
		msg2 := <-c
		fmt.Println(msg2.body)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("Leaving....")
}
