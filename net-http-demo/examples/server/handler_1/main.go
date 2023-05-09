package main

import (
	"fmt"
	"net/http"
)

type greeter func(w http.ResponseWriter, r *http.Request) // HL

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g(w, r) // HL
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! 🚀")
}

func main() {
	http.ListenAndServe(":8081", greeter(handleRequest)) // HL
}
