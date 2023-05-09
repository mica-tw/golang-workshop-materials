package main

import "net/http"

type apiHandler struct{}

func (ah apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Response"))
}

type webHandler struct{}

func (wh webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Web Response"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{}) // HL
	mux.Handle("/web/", webHandler{}) // HL
	http.ListenAndServe(":8080", mux)
}
