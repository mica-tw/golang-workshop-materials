package main

import "net/http"

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("API Response"))
	})
	http.ListenAndServe(":8080", nil) // HL
}
