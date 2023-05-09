package main

import "net/http"

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! from Go 🚀"))
}

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(greet)) // HL
}
