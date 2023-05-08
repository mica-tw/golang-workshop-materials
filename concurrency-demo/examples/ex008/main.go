package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func id(prefix string) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 16)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}

	return prefix + string(b)
}
func fingerprint(input string) string {
	rand.Seed(time.Now().UnixNano())

	hash := md5.Sum([]byte(input))
	hashStr := hex.EncodeToString(hash[:])

	b := make([]byte, 16)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}

	return hashStr[:8] + string(b) + hashStr[24:]
}

type Algorithm func(string) string

func service(fn Algorithm, param string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fn(param)
		}
	}()
	return ch
}

func main() {
	idSvc := service(id, "ch_")
	fpSvc := service(fingerprint, "001")
	for i := 0; i < 5; i++ {
		fmt.Println(<-idSvc)
		fmt.Println(<-fpSvc) // Waits for id service even if fingerprint is ready
	}
}
