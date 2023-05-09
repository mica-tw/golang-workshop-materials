package main

import (
	"fmt"
	"log"
	"net/http"
)

type greeter string // HL

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Custom-Header", "Hello")
	w.Header().Set("Another-Custom-Header", "World")
	fmt.Fprintln(w, "Hello, World! from", g)
}

func main() {
	log.Println("Starting server... http://localhost:8081")
	http.ListenAndServe(":8081", greeter("Go 🚀")) // HL
}
