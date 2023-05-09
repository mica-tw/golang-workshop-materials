package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("https://httpbingo.org/ip") // HL
	// handle error

	fmt.Println(res.Status)
	defer res.Body.Close()          // HL
	body, _ := io.ReadAll(res.Body) // HL
	// handle error
	fmt.Println(string(body))
}
