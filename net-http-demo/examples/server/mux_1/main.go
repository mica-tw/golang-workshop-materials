package main

import "net/http"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Response"))
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Web Response"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", apiHandler) // HL
	mux.HandleFunc("/web/", webHandler) // HL
	http.ListenAndServe(":8080", mux)
}
