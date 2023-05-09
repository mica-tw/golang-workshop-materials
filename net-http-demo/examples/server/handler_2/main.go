package main

import (
	"fmt"
	"net/http"
)

type context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
type greeter func(c *context) // HL

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g(&context{Writer: w, Request: r})
}

func main() {
	http.ListenAndServe(":8081", greeter(func(c *context) { // HL
		fmt.Fprintln(c.Writer, "Hello World! 🚀") // HL
	})) // HL
}
